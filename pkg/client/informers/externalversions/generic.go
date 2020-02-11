/*
Copyright The Flagger Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1beta1 "github.com/weaveworks/flagger/pkg/apis/appmesh/v1beta1"
	flaggerv1beta1 "github.com/weaveworks/flagger/pkg/apis/flagger/v1beta1"
	v1 "github.com/weaveworks/flagger/pkg/apis/gloo/v1"
	v1alpha3 "github.com/weaveworks/flagger/pkg/apis/istio/v1alpha3"
	projectcontourv1 "github.com/weaveworks/flagger/pkg/apis/projectcontour/v1"
	v1alpha1 "github.com/weaveworks/flagger/pkg/apis/smi/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=appmesh.k8s.aws, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("meshes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Appmesh().V1beta1().Meshes().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("virtualnodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Appmesh().V1beta1().VirtualNodes().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("virtualservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Appmesh().V1beta1().VirtualServices().Informer()}, nil

		// Group=flagger.app, Version=v1beta1
	case flaggerv1beta1.SchemeGroupVersion.WithResource("alertproviders"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Flagger().V1beta1().AlertProviders().Informer()}, nil
	case flaggerv1beta1.SchemeGroupVersion.WithResource("canaries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Flagger().V1beta1().Canaries().Informer()}, nil
	case flaggerv1beta1.SchemeGroupVersion.WithResource("metrictemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Flagger().V1beta1().MetricTemplates().Informer()}, nil

		// Group=gloo.solo.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("upstreamgroups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gloo().V1().UpstreamGroups().Informer()}, nil

		// Group=networking.istio.io, Version=v1alpha3
	case v1alpha3.SchemeGroupVersion.WithResource("destinationrules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().DestinationRules().Informer()}, nil
	case v1alpha3.SchemeGroupVersion.WithResource("virtualservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().VirtualServices().Informer()}, nil

		// Group=projectcontour.io, Version=v1
	case projectcontourv1.SchemeGroupVersion.WithResource("httpproxies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcontour().V1().HTTPProxies().Informer()}, nil

		// Group=split.smi-spec.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("trafficsplits"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Split().V1alpha1().TrafficSplits().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
