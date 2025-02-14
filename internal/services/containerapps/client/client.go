// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/certificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/containerapps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/containerappsrevisions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/daprcomponents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/managedenvironments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/managedenvironmentsstorages"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
)

type Client struct {
	CertificatesClient         *certificates.CertificatesClient
	ContainerAppClient         *containerapps.ContainerAppsClient
	ContainerAppRevisionClient *containerappsrevisions.ContainerAppsRevisionsClient
	DaprComponentsClient       *daprcomponents.DaprComponentsClient
	ManagedEnvironmentClient   *managedenvironments.ManagedEnvironmentsClient
	StorageClient              *managedenvironmentsstorages.ManagedEnvironmentsStoragesClient
}

func NewClient(o *common.ClientOptions) *Client {
	certificatesClient := certificates.NewCertificatesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&certificatesClient.Client, o.ResourceManagerAuthorizer)

	containerAppsClient := containerapps.NewContainerAppsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&containerAppsClient.Client, o.ResourceManagerAuthorizer)

	containerAppsRevisionsClient := containerappsrevisions.NewContainerAppsRevisionsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&containerAppsRevisionsClient.Client, o.ResourceManagerAuthorizer)

	managedEnvironmentClient := managedenvironments.NewManagedEnvironmentsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&managedEnvironmentClient.Client, o.ResourceManagerAuthorizer)

	managedEnvironmentStoragesClient := managedenvironmentsstorages.NewManagedEnvironmentsStoragesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&managedEnvironmentStoragesClient.Client, o.ResourceManagerAuthorizer)

	daprComponentClient := daprcomponents.NewDaprComponentsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&daprComponentClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		CertificatesClient:         &certificatesClient,
		ContainerAppClient:         &containerAppsClient,
		ContainerAppRevisionClient: &containerAppsRevisionsClient,
		DaprComponentsClient:       &daprComponentClient,
		ManagedEnvironmentClient:   &managedEnvironmentClient,
		StorageClient:              &managedEnvironmentStoragesClient,
	}
}
