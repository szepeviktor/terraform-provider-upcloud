package upcloud

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccDataSourceUpCloudIPAddresses_basic(t *testing.T) {
	var providers []*schema.Provider

	resourceName := "data.upcloud_ip_addresses.empty"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(&providers),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUpCloudIPAddressesConfig_empty(),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceUpCloudIPAddressesCheck(resourceName),
				),
			},
		},
	})
}

func testAccDataSourceUpCloudIPAddressesCheck(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("root module has no resource called %s", resourceName)
		}

		_, ipAddressesOk := rs.Primary.Attributes["addresses.#"]

		if !ipAddressesOk {
			return fmt.Errorf("addresses attribute is missing.")
		}

		return nil
	}
}

func testAccDataSourceUpCloudIPAddressesConfig_empty() string {
	return `
data "upcloud_ip_addresses" "empty" {}
`
}
