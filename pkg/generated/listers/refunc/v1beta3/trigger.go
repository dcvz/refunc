/*
Copyright 2021 The refunc Authors

TODO: choose a opensource licence.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1beta3

import (
	v1beta3 "github.com/refunc/refunc/pkg/apis/refunc/v1beta3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TriggerLister helps list Triggers.
// All objects returned here must be treated as read-only.
type TriggerLister interface {
	// List lists all Triggers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta3.Trigger, err error)
	// Triggers returns an object that can list and get Triggers.
	Triggers(namespace string) TriggerNamespaceLister
	TriggerListerExpansion
}

// triggerLister implements the TriggerLister interface.
type triggerLister struct {
	indexer cache.Indexer
}

// NewTriggerLister returns a new TriggerLister.
func NewTriggerLister(indexer cache.Indexer) TriggerLister {
	return &triggerLister{indexer: indexer}
}

// List lists all Triggers in the indexer.
func (s *triggerLister) List(selector labels.Selector) (ret []*v1beta3.Trigger, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta3.Trigger))
	})
	return ret, err
}

// Triggers returns an object that can list and get Triggers.
func (s *triggerLister) Triggers(namespace string) TriggerNamespaceLister {
	return triggerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TriggerNamespaceLister helps list and get Triggers.
// All objects returned here must be treated as read-only.
type TriggerNamespaceLister interface {
	// List lists all Triggers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta3.Trigger, err error)
	// Get retrieves the Trigger from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta3.Trigger, error)
	TriggerNamespaceListerExpansion
}

// triggerNamespaceLister implements the TriggerNamespaceLister
// interface.
type triggerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Triggers in the indexer for a given namespace.
func (s triggerNamespaceLister) List(selector labels.Selector) (ret []*v1beta3.Trigger, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta3.Trigger))
	})
	return ret, err
}

// Get retrieves the Trigger from the indexer for a given namespace and name.
func (s triggerNamespaceLister) Get(name string) (*v1beta3.Trigger, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta3.Resource("trigger"), name)
	}
	return obj.(*v1beta3.Trigger), nil
}