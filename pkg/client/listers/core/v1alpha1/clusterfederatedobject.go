// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubewharf/kubeadmiral/pkg/apis/core/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterFederatedObjectLister helps list ClusterFederatedObjects.
// All objects returned here must be treated as read-only.
type ClusterFederatedObjectLister interface {
	// List lists all ClusterFederatedObjects in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterFederatedObject, err error)
	// Get retrieves the ClusterFederatedObject from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterFederatedObject, error)
	ClusterFederatedObjectListerExpansion
}

// clusterFederatedObjectLister implements the ClusterFederatedObjectLister interface.
type clusterFederatedObjectLister struct {
	indexer cache.Indexer
}

// NewClusterFederatedObjectLister returns a new ClusterFederatedObjectLister.
func NewClusterFederatedObjectLister(indexer cache.Indexer) ClusterFederatedObjectLister {
	return &clusterFederatedObjectLister{indexer: indexer}
}

// List lists all ClusterFederatedObjects in the indexer.
func (s *clusterFederatedObjectLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterFederatedObject, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterFederatedObject))
	})
	return ret, err
}

// Get retrieves the ClusterFederatedObject from the index for a given name.
func (s *clusterFederatedObjectLister) Get(name string) (*v1alpha1.ClusterFederatedObject, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterfederatedobject"), name)
	}
	return obj.(*v1alpha1.ClusterFederatedObject), nil
}