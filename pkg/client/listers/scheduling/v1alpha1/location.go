//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	schedulingv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/scheduling/v1alpha1"
)

// LocationClusterLister can list Locations across all workspaces, or scope down to a LocationLister for one workspace.
type LocationClusterLister interface {
	List(selector labels.Selector) (ret []*schedulingv1alpha1.Location, err error)
	Cluster(cluster logicalcluster.Name) LocationLister
}

type locationClusterLister struct {
	indexer cache.Indexer
}

// NewLocationClusterLister returns a new LocationClusterLister.
func NewLocationClusterLister(indexer cache.Indexer) *locationClusterLister {
	return &locationClusterLister{indexer: indexer}
}

// List lists all Locations in the indexer across all workspaces.
func (s *locationClusterLister) List(selector labels.Selector) (ret []*schedulingv1alpha1.Location, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*schedulingv1alpha1.Location))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get Locations.
func (s *locationClusterLister) Cluster(cluster logicalcluster.Name) LocationLister {
	return &locationLister{indexer: s.indexer, cluster: cluster}
}

type LocationLister interface {
	List(selector labels.Selector) (ret []*schedulingv1alpha1.Location, err error)
	Get(name string) (*schedulingv1alpha1.Location, error)
}

// locationLister can list all Locations inside a workspace.
type locationLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all Locations in the indexer for a workspace.
func (s *locationLister) List(selector labels.Selector) (ret []*schedulingv1alpha1.Location, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*schedulingv1alpha1.Location)
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}

	return ret, err
}

// Get retrieves the Location from the indexer for a given workspace and name.
func (s *locationLister) Get(name string) (*schedulingv1alpha1.Location, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schedulingv1alpha1.Resource("Location"), name)
	}
	return obj.(*schedulingv1alpha1.Location), nil
}
