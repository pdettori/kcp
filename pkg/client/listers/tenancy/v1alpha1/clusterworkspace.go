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

	tenancyv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/tenancy/v1alpha1"
)

// ClusterWorkspaceClusterLister can list ClusterWorkspaces across all workspaces, or scope down to a ClusterWorkspaceLister for one workspace.
type ClusterWorkspaceClusterLister interface {
	List(selector labels.Selector) (ret []*tenancyv1alpha1.ClusterWorkspace, err error)
	Cluster(cluster logicalcluster.Name) ClusterWorkspaceLister
}

type clusterWorkspaceClusterLister struct {
	indexer cache.Indexer
}

// NewClusterWorkspaceClusterLister returns a new ClusterWorkspaceClusterLister.
func NewClusterWorkspaceClusterLister(indexer cache.Indexer) *clusterWorkspaceClusterLister {
	return &clusterWorkspaceClusterLister{indexer: indexer}
}

// List lists all ClusterWorkspaces in the indexer across all workspaces.
func (s *clusterWorkspaceClusterLister) List(selector labels.Selector) (ret []*tenancyv1alpha1.ClusterWorkspace, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*tenancyv1alpha1.ClusterWorkspace))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get ClusterWorkspaces.
func (s *clusterWorkspaceClusterLister) Cluster(cluster logicalcluster.Name) ClusterWorkspaceLister {
	return &clusterWorkspaceLister{indexer: s.indexer, cluster: cluster}
}

type ClusterWorkspaceLister interface {
	List(selector labels.Selector) (ret []*tenancyv1alpha1.ClusterWorkspace, err error)
	Get(name string) (*tenancyv1alpha1.ClusterWorkspace, error)
}

// clusterWorkspaceLister can list all ClusterWorkspaces inside a workspace.
type clusterWorkspaceLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all ClusterWorkspaces in the indexer for a workspace.
func (s *clusterWorkspaceLister) List(selector labels.Selector) (ret []*tenancyv1alpha1.ClusterWorkspace, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*tenancyv1alpha1.ClusterWorkspace)
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

// Get retrieves the ClusterWorkspace from the indexer for a given workspace and name.
func (s *clusterWorkspaceLister) Get(name string) (*tenancyv1alpha1.ClusterWorkspace, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(tenancyv1alpha1.Resource("ClusterWorkspace"), name)
	}
	return obj.(*tenancyv1alpha1.ClusterWorkspace), nil
}
