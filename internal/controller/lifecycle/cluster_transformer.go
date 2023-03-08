/*
Copyright ApeCloud, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package lifecycle

import (
	"encoding/json"

	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	"github.com/apecloud/kubeblocks/internal/constant"
	"github.com/apecloud/kubeblocks/internal/controller/builder"
	"github.com/apecloud/kubeblocks/internal/controller/component"
	"github.com/apecloud/kubeblocks/internal/controller/graph"
	"github.com/apecloud/kubeblocks/internal/controller/plan"
	intctrltypes "github.com/apecloud/kubeblocks/internal/controller/types"
	intctrlutil "github.com/apecloud/kubeblocks/internal/controllerutil"
)

// clusterTransformer transforms a Cluster to a K8s objects DAG
// TODO: remove cli and ctx, we should read all objects needed, and then do pure objects computation
type clusterTransformer struct {
	cc  compoundCluster
	cli client.Client
	ctx intctrlutil.RequestCtx
}

func (c *clusterTransformer) Transform(dag *graph.DAG) error {
	// put the cluster object first, it will be root vertex of DAG
	patch := client.MergeFrom(c.cc.cluster.DeepCopy())
	rootVertex := &lifecycleVertex{obj: c.cc.cluster, patch: patch}
	dag.AddVertex(rootVertex)

	// we copy the K8s objects prepare stage directly first
	// TODO: refactor plan.PrepareComponentResources
	resourcesQueue := make([]client.Object, 0, 3)
	task := intctrltypes.ReconcileTask{
		Cluster:           c.cc.cluster,
		ClusterDefinition: &c.cc.cd,
		ClusterVersion:    &c.cc.cv,
		Resources:         &resourcesQueue,
	}

	clusterBackupResourceMap, err := getClusterBackupSourceMap(c.cc.cluster)
	if err != nil {
		return err
	}

	clusterCompSpecMap := c.cc.cluster.GetDefNameMappingComponents()
	clusterCompVerMap := c.cc.cv.GetDefNameMappingComponents()
	process1stComp := true

	prepareComp := func(synthesizedComp *component.SynthesizedComponent) error {
		iParams := task
		iParams.Component = synthesizedComp
		if process1stComp && len(synthesizedComp.Services) > 0 {
			if err := prepareConnCredential(&iParams); err != nil {
				return err
			}
			process1stComp = false
		}

		// build info that needs to be restored from backup
		backupSourceName := clusterBackupResourceMap[synthesizedComp.Name]
		if len(backupSourceName) > 0 {
			if err := component.BuildRestoredInfo(c.ctx, c.cli, c.cc.cluster.Namespace, synthesizedComp, backupSourceName); err != nil {
				return err
			}
		}
		return plan.PrepareComponentResources(c.ctx, c.cli, &iParams)
	}

	for _, compDef := range c.cc.cd.Spec.ComponentDefs {
		compDefName := compDef.Name
		compVer := clusterCompVerMap[compDefName]
		compSpecs := clusterCompSpecMap[compDefName]
		for _, compSpec := range compSpecs {
			if err := prepareComp(component.BuildComponent(c.ctx, *c.cc.cluster, c.cc.cd, compDef, compSpec, compVer)); err != nil {
				return err
			}
		}
	}

	// now task.Resources to DAG vertices
	for _, object := range *task.Resources {
		vertex := &lifecycleVertex{obj: object}
		dag.AddVertex(vertex)
		dag.Connect(rootVertex, vertex)
	}
	return nil
}

func prepareConnCredential(task *intctrltypes.ReconcileTask) error {
	secret, err := builder.BuildConnCredential(task.GetBuilderParams())
	if err != nil {
		return err
	}
	// must make sure secret resources are created before others
	task.InsertResource(secret)
	return nil
}

// getClusterBackupSourceMap gets the backup source map from cluster.annotations
func getClusterBackupSourceMap(cluster *appsv1alpha1.Cluster) (map[string]string, error) {
	compBackupMapString := cluster.Annotations[constant.RestoreFromBackUpAnnotationKey]
	if len(compBackupMapString) == 0 {
		return nil, nil
	}
	compBackupMap := map[string]string{}
	err := json.Unmarshal([]byte(compBackupMapString), &compBackupMap)
	return compBackupMap, err
}