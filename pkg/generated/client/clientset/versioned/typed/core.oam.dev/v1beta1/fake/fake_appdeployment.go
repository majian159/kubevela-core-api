/*
Copyright 2021 The KubeVela Authors.

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

	v1beta1 "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAppDeployments implements AppDeploymentInterface
type FakeAppDeployments struct {
	Fake *FakeCoreV1beta1
	ns   string
}

var appdeploymentsResource = schema.GroupVersionResource{Group: "core.oam.dev", Version: "v1beta1", Resource: "appdeployments"}

var appdeploymentsKind = schema.GroupVersionKind{Group: "core.oam.dev", Version: "v1beta1", Kind: "AppDeployment"}

// Get takes name of the appDeployment, and returns the corresponding appDeployment object, and an error if there is any.
func (c *FakeAppDeployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.AppDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appdeploymentsResource, c.ns, name), &v1beta1.AppDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AppDeployment), err
}

// List takes label and field selectors, and returns the list of AppDeployments that match those selectors.
func (c *FakeAppDeployments) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.AppDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appdeploymentsResource, appdeploymentsKind, c.ns, opts), &v1beta1.AppDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.AppDeploymentList{ListMeta: obj.(*v1beta1.AppDeploymentList).ListMeta}
	for _, item := range obj.(*v1beta1.AppDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appDeployments.
func (c *FakeAppDeployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appdeploymentsResource, c.ns, opts))

}

// Create takes the representation of a appDeployment and creates it.  Returns the server's representation of the appDeployment, and an error, if there is any.
func (c *FakeAppDeployments) Create(ctx context.Context, appDeployment *v1beta1.AppDeployment, opts v1.CreateOptions) (result *v1beta1.AppDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appdeploymentsResource, c.ns, appDeployment), &v1beta1.AppDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AppDeployment), err
}

// Update takes the representation of a appDeployment and updates it. Returns the server's representation of the appDeployment, and an error, if there is any.
func (c *FakeAppDeployments) Update(ctx context.Context, appDeployment *v1beta1.AppDeployment, opts v1.UpdateOptions) (result *v1beta1.AppDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appdeploymentsResource, c.ns, appDeployment), &v1beta1.AppDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AppDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppDeployments) UpdateStatus(ctx context.Context, appDeployment *v1beta1.AppDeployment, opts v1.UpdateOptions) (*v1beta1.AppDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(appdeploymentsResource, "status", c.ns, appDeployment), &v1beta1.AppDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AppDeployment), err
}

// Delete takes name of the appDeployment and deletes it. Returns an error if one occurs.
func (c *FakeAppDeployments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(appdeploymentsResource, c.ns, name), &v1beta1.AppDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppDeployments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appdeploymentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.AppDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched appDeployment.
func (c *FakeAppDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.AppDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appdeploymentsResource, c.ns, name, pt, data, subresources...), &v1beta1.AppDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AppDeployment), err
}
