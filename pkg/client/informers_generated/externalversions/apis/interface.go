// Copyright (c) 2020 Tigera, Inc. All rights reserved.

// Code generated by informer-gen. DO NOT EDIT.

package projectcalico

import (
	projectcalico "github.com/tigera/api/pkg/client/informers_generated/externalversions/apis/projectcalico"
	internalinterfaces "github.com/tigera/api/pkg/client/informers_generated/externalversions/internalinterfaces"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// Projectcalico provides access to shared informers for resources in Projectcalico.
	Projectcalico() projectcalico.Interface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Projectcalico returns a new projectcalico.Interface.
func (g *group) Projectcalico() projectcalico.Interface {
	return projectcalico.New(g.factory, g.namespace, g.tweakListOptions)
}
