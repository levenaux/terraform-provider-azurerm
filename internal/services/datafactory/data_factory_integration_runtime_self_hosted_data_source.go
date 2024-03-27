// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactory

import (
	"fmt"
	"regexp"
	"time"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/integrationruntimes"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azurerm/internal/timeouts"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

func dataSourceDataFactoryIntegrationRuntimeSelfHosted() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Read: dataSourceDataFactoryIntegrationRuntimeSelfHostedRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"name": {
				Type:     pluginsdk.TypeString,
				Required: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile(`^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`),
					`Invalid name for Integration Runtime: minimum 3 characters, must start and end with a number or a letter, may only consist of letters, numbers and dashes and no consecutive dashes.`,
				),
			},
			"data_factory_name": {
				Type:     pluginsdk.TypeString,
				Required: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile(`^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`),
					`Invalid name for Data Factory, see https://docs.microsoft.com/azure/data-factory/naming-rules`,
				),
			},
			"resource_group_name": commonschema.ResourceGroupNameForDataSource(),
		},
	}
}

func dataSourceDataFactoryIntegrationRuntimeSelfHostedRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DataFactory.IntegrationRuntimesClient
	subscriptionId := meta.(*clients.Client).DataFactory.IntegrationRuntimesClient.SubscriptionID
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id := integrationruntimes.NewIntegrationRuntimeID(subscriptionId, d.Get("resource_group_name").(string), d.Get("data_factory_name").(string), d.Get("name").(string))
	resp, err := client.Get(ctx, id.ResourceGroupName, id.FactoryName, id.IntegrationRuntimeName, "")
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return fmt.Errorf("%s was not found", id)
		}
		return fmt.Errorf("retreiving %s: %+v", id, err)
	}

	d.SetId(id.ID())

	return nil
}
