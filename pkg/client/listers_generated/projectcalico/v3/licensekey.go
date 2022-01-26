// Copyright (c) 2022 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LicenseKeyLister helps list LicenseKeys.
// All objects returned here must be treated as read-only.
type LicenseKeyLister interface {
	// List lists all LicenseKeys in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.LicenseKey, err error)
	// Get retrieves the LicenseKey from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v3.LicenseKey, error)
	LicenseKeyListerExpansion
}

// licenseKeyLister implements the LicenseKeyLister interface.
type licenseKeyLister struct {
	indexer cache.Indexer
}

// NewLicenseKeyLister returns a new LicenseKeyLister.
func NewLicenseKeyLister(indexer cache.Indexer) LicenseKeyLister {
	return &licenseKeyLister{indexer: indexer}
}

// List lists all LicenseKeys in the indexer.
func (s *licenseKeyLister) List(selector labels.Selector) (ret []*v3.LicenseKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.LicenseKey))
	})
	return ret, err
}

// Get retrieves the LicenseKey from the index for a given name.
func (s *licenseKeyLister) Get(name string) (*v3.LicenseKey, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("licensekey"), name)
	}
	return obj.(*v3.LicenseKey), nil
}
