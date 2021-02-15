/*
Copyright 2019 The OpenEBS Authors

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

package provisioner

import (
	menv "github.com/openebs/maya/pkg/env/v1alpha1"
)

//This file defines the environement variable names that are specific
// to this provisioner. In addition to the variables defined in this file,
// provisioner also uses the following:
//   OPENEBS_NAMESPACE
//   OPENEBS_SERVICE_ACCOUNT
//   OPENEBS_IO_K8S_MASTER
//   OPENEBS_IO_KUBE_CONFIG

const (
	// ProvisionerNFSServerType is the environment variable that
	// allows user to specify the default NFS Server type to be used.
	ProvisionerNFSServerType menv.ENVKey = "OPENEBS_IO_NFS_SERVER_TYPE"

	// ProvisionerExportsSC is the environment variable that provides the
	// default storage class to be used for exports PVC mount used by NFS Server.
	ProvisionerExportsSC menv.ENVKey = "OPENEBS_IO_EXPORTS_SC"

	// ProvisionerNFSServerUseClusterIP is the environment variable that
	// allows user to specify if ClusterIP should be used in NFS K8s Service
	ProvisionerNFSServerUseClusterIP menv.ENVKey = "OPENEBS_IO_NFS_SERVER_USE_CLUSTERIP"
)

var (
	defaultNFSServerType = "kernel"
	defaultExportsSC     = ""
)

func getOpenEBSNamespace() string {
	return menv.Get(menv.OpenEBSNamespace)
}
func getDefaultExportsSC() string {
	return menv.GetOrDefault(ProvisionerExportsSC, string(defaultExportsSC))
}

func getDefaultNFSServerType() string {
	return menv.GetOrDefault(ProvisionerNFSServerType, string(defaultNFSServerType))
}

func getOpenEBSServiceAccountName() string {
	return menv.Get(menv.OpenEBSServiceAccount)
}
