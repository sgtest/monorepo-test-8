/*
Copyright 2016 The Kubernetes Authors.

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

package rest

import (
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-3/pkg/api"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-3/pkg/apis/certificates"
	certificatesapiv1beta1 "github.com/sourcegraph/monorepo-test-1/kubernetes-3/pkg/apis/certificates/v1beta1"
	certificatestore "github.com/sourcegraph/monorepo-test-1/kubernetes-3/pkg/registry/certificates/certificates/storage"
)

type RESTStorageProvider struct{}

func (p RESTStorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool) {
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(certificates.GroupName, api.Registry, api.Scheme, api.ParameterCodec, api.Codecs)

	if apiResourceConfigSource.AnyResourcesForVersionEnabled(certificatesapiv1beta1.SchemeGroupVersion) {
		apiGroupInfo.VersionedResourcesStorageMap[certificatesapiv1beta1.SchemeGroupVersion.Version] = p.v1beta1Storage(apiResourceConfigSource, restOptionsGetter)
		apiGroupInfo.GroupMeta.GroupVersion = certificatesapiv1beta1.SchemeGroupVersion
	}

	return apiGroupInfo, true
}

func (p RESTStorageProvider) v1beta1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
	version := certificatesapiv1beta1.SchemeGroupVersion

	storage := map[string]rest.Storage{}
	if apiResourceConfigSource.ResourceEnabled(version.WithResource("certificatesigningrequests")) {
		csrStorage, csrStatusStorage, csrApprovalStorage := certificatestore.NewREST(restOptionsGetter)
		storage["certificatesigningrequests"] = csrStorage
		storage["certificatesigningrequests/status"] = csrStatusStorage
		storage["certificatesigningrequests/approval"] = csrApprovalStorage
	}
	return storage
}

func (p RESTStorageProvider) GroupName() string {
	return certificates.GroupName
}
