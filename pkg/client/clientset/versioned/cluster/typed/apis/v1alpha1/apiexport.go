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
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	apisv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/apis/v1alpha1"
	apisv1alpha1client "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/apis/v1alpha1"
)

// APIExportsClusterGetter has a method to return a APIExportClusterInterface.
// A group's cluster client should implement this interface.
type APIExportsClusterGetter interface {
	APIExports() APIExportClusterInterface
}

// APIExportClusterInterface can operate on APIExports across all clusters,
// or scope down to one cluster and return a apisv1alpha1client.APIExportInterface.
type APIExportClusterInterface interface {
	Cluster(logicalcluster.Name) apisv1alpha1client.APIExportInterface
	List(ctx context.Context, opts metav1.ListOptions) (*apisv1alpha1.APIExportList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type aPIExportsClusterInterface struct {
	clientCache kcpclient.Cache[*apisv1alpha1client.ApisV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *aPIExportsClusterInterface) Cluster(name logicalcluster.Name) apisv1alpha1client.APIExportInterface {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(name).APIExports()
}

// List returns the entire collection of all APIExports across all clusters.
func (c *aPIExportsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*apisv1alpha1.APIExportList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).APIExports().List(ctx, opts)
}

// Watch begins to watch all APIExports across all clusters.
func (c *aPIExportsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).APIExports().Watch(ctx, opts)
}
