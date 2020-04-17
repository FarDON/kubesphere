/*
Copyright 2019 The KubeSphere authors.

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

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha2 "kubesphere.io/kubesphere/pkg/apis/iam/v1alpha2"
)

// RoleBindingLister helps list RoleBindings.
type RoleBindingLister interface {
	// List lists all RoleBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.RoleBinding, err error)
	// Get retrieves the RoleBinding from the index for a given name.
	Get(name string) (*v1alpha2.RoleBinding, error)
	RoleBindingListerExpansion
}

// roleBindingLister implements the RoleBindingLister interface.
type roleBindingLister struct {
	indexer cache.Indexer
}

// NewRoleBindingLister returns a new RoleBindingLister.
func NewRoleBindingLister(indexer cache.Indexer) RoleBindingLister {
	return &roleBindingLister{indexer: indexer}
}

// List lists all RoleBindings in the indexer.
func (s *roleBindingLister) List(selector labels.Selector) (ret []*v1alpha2.RoleBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.RoleBinding))
	})
	return ret, err
}

// Get retrieves the RoleBinding from the index for a given name.
func (s *roleBindingLister) Get(name string) (*v1alpha2.RoleBinding, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("rolebinding"), name)
	}
	return obj.(*v1alpha2.RoleBinding), nil
}