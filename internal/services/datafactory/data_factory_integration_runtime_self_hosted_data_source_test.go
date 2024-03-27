// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactory_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
)

type IntegrationRuntimeDataSource struct{}

func TestAccDataFactoryIntegrationRuntimeSelfHostedDataSource_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_data_factory_integration_runtime_self_hosted", "test")
	r := IntegrationRuntimeDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check:  acceptance.ComposeTestCheckFunc(),
		},
	})
}

func (IntegrationRuntimeDataSource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

data "azurerm_data_factory_integration_runtime_self_hosted" "test" {
	name                = "acctestSIR"
	data_factory_name   =  azurerm_data_factory_integration_runtime_self_hosted.test.name
	resource_group_name = azurerm_resource_group.test.name
}
`, IntegrationRuntimeSelfHostedResource{}.basic(data))
}
