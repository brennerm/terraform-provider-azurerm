package azurerm

import (
	"fmt"
	"time"

	appconf "github.com/Azure/azure-sdk-for-go/services/appconfiguration/mgmt/2019-10-01/appconfiguration"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tags"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func dataSourceArmAppConfigurationValue() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceArmAppConfigurationValueRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"app_conf_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: azure.ValidateResourceID,
			},

			"key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"label": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"content_type": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"tags": tags.SchemaDataSource(),
		},
	}
}

func dataSourceArmAppConfigurationValueRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).AppConfiguration.AppConfigurationsClient
	ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
	defer cancel()

	appConfID := d.Get("app_conf_id").(string)
	id, err := azure.ParseAzureResourceID(appConfID)
	if err != nil {
		return fmt.Errorf("Failed to parse App Configuration ID %s: %+v", appConfID, err)
	}

	resourceGroupName := id.ResourceGroup
	appConfName := id.Path["configurationStores"]

	key := d.Get("key").(string)
	label := d.Get("label").(string)
	listKeyValueParam := appconf.ListKeyValueParameters{
		Key:   &key,
		Label: &label,
	}

	resp, err := client.ListKeyValue(ctx, resourceGroupName, appConfName, listKeyValueParam)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return fmt.Errorf("Key %q with Label %q (App Configuration %q) does not exist", key, label, appConfName)
		}
		return fmt.Errorf("Error making Read request on Azure App Configuration Key %s: %+v", key, err)
	}

	d.SetId(*resp.ETag)
	d.Set("value", resp.Value)
	d.Set("content_type", resp.ContentType)

	return tags.FlattenAndSet(d, resp.Tags)
}
