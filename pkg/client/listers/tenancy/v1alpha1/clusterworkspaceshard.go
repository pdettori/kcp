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

// ClusterWorkspaceShardClusterLister can list ClusterWorkspaceShards across all workspaces, or scope down to a ClusterWorkspaceShardLister for one workspace.
type ClusterWorkspaceShardClusterLister interface {
	List(selector labels.Selector) (ret []*tenancyv1alpha1.ClusterWorkspaceShard, err error)
	Cluster(cluster logicalcluster.Name) ClusterWorkspaceShardLister
}

type clusterWorkspaceShardClusterLister struct {
	indexer cache.Indexer
}

// NewClusterWorkspaceShardClusterLister returns a new ClusterWorkspaceShardClusterLister.
func NewClusterWorkspaceShardClusterLister(indexer cache.Indexer) *clusterWorkspaceShardClusterLister {
	return &clusterWorkspaceShardClusterLister{indexer: indexer}
}

// List lists all ClusterWorkspaceShards in the indexer across all workspaces.
func (s *clusterWorkspaceShardClusterLister) List(selector labels.Selector) (ret []*tenancyv1alpha1.ClusterWorkspaceShard, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*tenancyv1alpha1.ClusterWorkspaceShard))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get ClusterWorkspaceShards.
func (s *clusterWorkspaceShardClusterLister) Cluster(cluster logicalcluster.Name) ClusterWorkspaceShardLister {
	return &clusterWorkspaceShardLister{indexer: s.indexer, cluster: cluster}
}

type ClusterWorkspaceShardLister interface {
	List(selector labels.Selector) (ret []*tenancyv1alpha1.ClusterWorkspaceShard, err error)
	Get(name string) (*tenancyv1alpha1.ClusterWorkspaceShard, error)
}

// clusterWorkspaceShardLister can list all ClusterWorkspaceShards inside a workspace.
type clusterWorkspaceShardLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all ClusterWorkspaceShards in the indexer for a workspace.
func (s *clusterWorkspaceShardLister) List(selector labels.Selector) (ret []*tenancyv1alpha1.ClusterWorkspaceShard, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*tenancyv1alpha1.ClusterWorkspaceShard)
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

// Get retrieves the ClusterWorkspaceShard from the indexer for a given workspace and name.
func (s *clusterWorkspaceShardLister) Get(name string) (*tenancyv1alpha1.ClusterWorkspaceShard, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(tenancyv1alpha1.Resource("ClusterWorkspaceShard"), name)
	}
	return obj.(*tenancyv1alpha1.ClusterWorkspaceShard), nil
}
