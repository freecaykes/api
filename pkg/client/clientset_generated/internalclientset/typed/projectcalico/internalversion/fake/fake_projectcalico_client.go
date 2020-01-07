// Copyright (c) 2020 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	internalversion "github.com/tigera/api/pkg/client/clientset_generated/internalclientset/typed/projectcalico/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeProjectcalico struct {
	*testing.Fake
}

func (c *FakeProjectcalico) GlobalReportTypes(namespace string) internalversion.GlobalReportTypeInterface {
	return &FakeGlobalReportTypes{c, namespace}
}

func (c *FakeProjectcalico) LicenseKeys(namespace string) internalversion.LicenseKeyInterface {
	return &FakeLicenseKeys{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeProjectcalico) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
