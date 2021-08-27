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
// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ResourceTrackerLister helps list ResourceTrackers.
// All objects returned here must be treated as read-only.
type ResourceTrackerLister interface {
	// List lists all ResourceTrackers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.ResourceTracker, err error)
	// ResourceTrackers returns an object that can list and get ResourceTrackers.
	ResourceTrackers(namespace string) ResourceTrackerNamespaceLister
	ResourceTrackerListerExpansion
}

// resourceTrackerLister implements the ResourceTrackerLister interface.
type resourceTrackerLister struct {
	indexer cache.Indexer
}

// NewResourceTrackerLister returns a new ResourceTrackerLister.
func NewResourceTrackerLister(indexer cache.Indexer) ResourceTrackerLister {
	return &resourceTrackerLister{indexer: indexer}
}

// List lists all ResourceTrackers in the indexer.
func (s *resourceTrackerLister) List(selector labels.Selector) (ret []*v1beta1.ResourceTracker, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ResourceTracker))
	})
	return ret, err
}

// ResourceTrackers returns an object that can list and get ResourceTrackers.
func (s *resourceTrackerLister) ResourceTrackers(namespace string) ResourceTrackerNamespaceLister {
	return resourceTrackerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResourceTrackerNamespaceLister helps list and get ResourceTrackers.
// All objects returned here must be treated as read-only.
type ResourceTrackerNamespaceLister interface {
	// List lists all ResourceTrackers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.ResourceTracker, err error)
	// Get retrieves the ResourceTracker from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.ResourceTracker, error)
	ResourceTrackerNamespaceListerExpansion
}

// resourceTrackerNamespaceLister implements the ResourceTrackerNamespaceLister
// interface.
type resourceTrackerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResourceTrackers in the indexer for a given namespace.
func (s resourceTrackerNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.ResourceTracker, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ResourceTracker))
	})
	return ret, err
}

// Get retrieves the ResourceTracker from the indexer for a given namespace and name.
func (s resourceTrackerNamespaceLister) Get(name string) (*v1beta1.ResourceTracker, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("resourcetracker"), name)
	}
	return obj.(*v1beta1.ResourceTracker), nil
}
