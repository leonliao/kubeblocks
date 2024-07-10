/*
Copyright (C) 2022-2024 ApeCloud Co., Ltd

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/apecloud/kubeblocks/apis/apps/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeParametersDefinitions implements ParametersDefinitionInterface
type FakeParametersDefinitions struct {
	Fake *FakeAppsV1beta1
}

var parametersdefinitionsResource = v1beta1.SchemeGroupVersion.WithResource("parametersdefinitions")

var parametersdefinitionsKind = v1beta1.SchemeGroupVersion.WithKind("ParametersDefinition")

// Get takes name of the parametersDefinition, and returns the corresponding parametersDefinition object, and an error if there is any.
func (c *FakeParametersDefinitions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ParametersDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(parametersdefinitionsResource, name), &v1beta1.ParametersDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ParametersDefinition), err
}

// List takes label and field selectors, and returns the list of ParametersDefinitions that match those selectors.
func (c *FakeParametersDefinitions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ParametersDefinitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(parametersdefinitionsResource, parametersdefinitionsKind, opts), &v1beta1.ParametersDefinitionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ParametersDefinitionList{ListMeta: obj.(*v1beta1.ParametersDefinitionList).ListMeta}
	for _, item := range obj.(*v1beta1.ParametersDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested parametersDefinitions.
func (c *FakeParametersDefinitions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(parametersdefinitionsResource, opts))
}

// Create takes the representation of a parametersDefinition and creates it.  Returns the server's representation of the parametersDefinition, and an error, if there is any.
func (c *FakeParametersDefinitions) Create(ctx context.Context, parametersDefinition *v1beta1.ParametersDefinition, opts v1.CreateOptions) (result *v1beta1.ParametersDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(parametersdefinitionsResource, parametersDefinition), &v1beta1.ParametersDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ParametersDefinition), err
}

// Update takes the representation of a parametersDefinition and updates it. Returns the server's representation of the parametersDefinition, and an error, if there is any.
func (c *FakeParametersDefinitions) Update(ctx context.Context, parametersDefinition *v1beta1.ParametersDefinition, opts v1.UpdateOptions) (result *v1beta1.ParametersDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(parametersdefinitionsResource, parametersDefinition), &v1beta1.ParametersDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ParametersDefinition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeParametersDefinitions) UpdateStatus(ctx context.Context, parametersDefinition *v1beta1.ParametersDefinition, opts v1.UpdateOptions) (*v1beta1.ParametersDefinition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(parametersdefinitionsResource, "status", parametersDefinition), &v1beta1.ParametersDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ParametersDefinition), err
}

// Delete takes name of the parametersDefinition and deletes it. Returns an error if one occurs.
func (c *FakeParametersDefinitions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(parametersdefinitionsResource, name, opts), &v1beta1.ParametersDefinition{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeParametersDefinitions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(parametersdefinitionsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ParametersDefinitionList{})
	return err
}

// Patch applies the patch and returns the patched parametersDefinition.
func (c *FakeParametersDefinitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ParametersDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(parametersdefinitionsResource, name, pt, data, subresources...), &v1beta1.ParametersDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ParametersDefinition), err
}
