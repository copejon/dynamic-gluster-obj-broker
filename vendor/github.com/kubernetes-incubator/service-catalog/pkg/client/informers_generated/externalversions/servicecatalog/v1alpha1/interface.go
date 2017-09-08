/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	internalinterfaces "github.com/kubernetes-incubator/service-catalog/pkg/client/informers_generated/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ServiceBrokers returns a ServiceBrokerInformer.
	ServiceBrokers() ServiceBrokerInformer
	// ServiceClasses returns a ServiceClassInformer.
	ServiceClasses() ServiceClassInformer
	// ServiceInstances returns a ServiceInstanceInformer.
	ServiceInstances() ServiceInstanceInformer
	// ServiceInstanceCredentials returns a ServiceInstanceCredentialInformer.
	ServiceInstanceCredentials() ServiceInstanceCredentialInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// ServiceBrokers returns a ServiceBrokerInformer.
func (v *version) ServiceBrokers() ServiceBrokerInformer {
	return &serviceBrokerInformer{factory: v.SharedInformerFactory}
}

// ServiceClasses returns a ServiceClassInformer.
func (v *version) ServiceClasses() ServiceClassInformer {
	return &serviceClassInformer{factory: v.SharedInformerFactory}
}

// ServiceInstances returns a ServiceInstanceInformer.
func (v *version) ServiceInstances() ServiceInstanceInformer {
	return &serviceInstanceInformer{factory: v.SharedInformerFactory}
}

// ServiceInstanceCredentials returns a ServiceInstanceCredentialInformer.
func (v *version) ServiceInstanceCredentials() ServiceInstanceCredentialInformer {
	return &serviceInstanceCredentialInformer{factory: v.SharedInformerFactory}
}
