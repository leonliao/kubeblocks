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
	clientset "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned"
	appsv1 "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/typed/apps/v1"
	fakeappsv1 "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/typed/apps/v1/fake"
	appsv1alpha1 "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/typed/apps/v1alpha1"
	fakeappsv1alpha1 "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/typed/apps/v1alpha1/fake"
	appsv1beta1 "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/typed/apps/v1beta1"
	fakeappsv1beta1 "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/typed/apps/v1beta1/fake"
	dataprotectionv1alpha1 "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/typed/dataprotection/v1alpha1"
	fakedataprotectionv1alpha1 "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/typed/dataprotection/v1alpha1/fake"
	extensionsv1alpha1 "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/typed/extensions/v1alpha1"
	fakeextensionsv1alpha1 "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/typed/extensions/v1alpha1/fake"
	storagev1alpha1 "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/typed/storage/v1alpha1"
	fakestoragev1alpha1 "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/typed/storage/v1alpha1/fake"
	workloadsv1alpha1 "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/typed/workloads/v1alpha1"
	fakeworkloadsv1alpha1 "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/typed/workloads/v1alpha1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// AppsV1alpha1 retrieves the AppsV1alpha1Client
func (c *Clientset) AppsV1alpha1() appsv1alpha1.AppsV1alpha1Interface {
	return &fakeappsv1alpha1.FakeAppsV1alpha1{Fake: &c.Fake}
}

// AppsV1beta1 retrieves the AppsV1beta1Client
func (c *Clientset) AppsV1beta1() appsv1beta1.AppsV1beta1Interface {
	return &fakeappsv1beta1.FakeAppsV1beta1{Fake: &c.Fake}
}

// AppsV1 retrieves the AppsV1Client
func (c *Clientset) AppsV1() appsv1.AppsV1Interface {
	return &fakeappsv1.FakeAppsV1{Fake: &c.Fake}
}

// DataprotectionV1alpha1 retrieves the DataprotectionV1alpha1Client
func (c *Clientset) DataprotectionV1alpha1() dataprotectionv1alpha1.DataprotectionV1alpha1Interface {
	return &fakedataprotectionv1alpha1.FakeDataprotectionV1alpha1{Fake: &c.Fake}
}

// ExtensionsV1alpha1 retrieves the ExtensionsV1alpha1Client
func (c *Clientset) ExtensionsV1alpha1() extensionsv1alpha1.ExtensionsV1alpha1Interface {
	return &fakeextensionsv1alpha1.FakeExtensionsV1alpha1{Fake: &c.Fake}
}

// StorageV1alpha1 retrieves the StorageV1alpha1Client
func (c *Clientset) StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface {
	return &fakestoragev1alpha1.FakeStorageV1alpha1{Fake: &c.Fake}
}

// WorkloadsV1alpha1 retrieves the WorkloadsV1alpha1Client
func (c *Clientset) WorkloadsV1alpha1() workloadsv1alpha1.WorkloadsV1alpha1Interface {
	return &fakeworkloadsv1alpha1.FakeWorkloadsV1alpha1{Fake: &c.Fake}
}
