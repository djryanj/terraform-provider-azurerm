package dashboard

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/azure"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/dashboard/validate"
	networkValidate "github.com/hashicorp/terraform-provider-azurerm/internal/services/network/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

// "context"
// "fmt"
// "regexp"
// "time"

// "github.com/hashicorp/go-azure-helpers/lang/response"
// "github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
// "github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
// "github.com/hashicorp/go-azure-helpers/resourcemanager/location"
// "github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2023-09-01/managedprivateendpoints"
// "github.com/hashicorp/terraform-provider-azurerm/internal/services/dashboard/parse"
//
// "github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
// "github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
// "github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
// "github.com/hashicorp/terraform-provider-azurerm/utils"

var _ sdk.ResourceWithUpdate = ManagedPrivateEndpointResource{}

type ManagedPrivateEndpointModel struct {
	Name             string            `tfschema:"name"`
	Location         string            `tfschema:"location"`
	Tags             map[string]string `tfschema:"tags"`
	ManagedGrafanaId string            `tfschema:"managed_grafana_id"`
	TargetResourceId string            `tfschema:"target_resource_id"`
}

func (r ManagedPrivateEndpointResource) ModelObject() interface{} {
	return &ManagedPrivateEndpointModel{}
}

type ManagedPrivateEndpointResource struct{}

func (r ManagedPrivateEndpointResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"name": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"location": commonschema.Location(),

		"tags": commonschema.Tags(),

		"managed_grafana_id": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validate.ValidateGrafanaID,
		},

		"type": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
			ValidateFunc: validation.StringInSlice([]string{
				"Microsoft.Kusto/clusters", "Microsoft.Insights/privatelinkscopes", "Microsoft.Monitor/accounts", "Microsoft.Sql/servers", "Microsoft.DocumentDB/databaseAccounts", "Microsoft.Network/privateLinkServices",
			}, false),
		},

		"target_resource_id": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: azure.ValidateResourceID,
		},

		"subresource_name": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: networkValidate.PrivateLinkSubResourceName,
		},
	}
}

func (ManagedPrivateEndpointResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (ManagedPrivateEndpointResource) ResourceType() string {
	return "azurerm_dashboard_grafana_managed_private_endpoint"
}
