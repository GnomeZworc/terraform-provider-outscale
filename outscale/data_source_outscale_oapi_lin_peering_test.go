package outscale

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDataSourceOutscaleOAPILinPeeringConnection_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			skipIfNoOAPI(t)
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceOutscaleOAPILinPeeringConnectionConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceOutscaleOAPILinPeeringConnectionCheck("outscale_net_peering.net_peering"),
				),
			},
		},
	})
}

func testAccDataSourceOutscaleOAPILinPeeringConnectionCheck(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", name)
		}

		pcxRs, ok := s.RootModule().Resources["outscale_net_peering.net_peering"]
		if !ok {
			return fmt.Errorf("can't find outscale_net_peering.net_peering in state")
		}

		attr := rs.Primary.Attributes

		if attr["id"] != pcxRs.Primary.Attributes["id"] {
			return fmt.Errorf(
				"id is %s; want %s",
				attr["id"],
				pcxRs.Primary.Attributes["id"],
			)
		}

		return nil
	}
}

const testAccDataSourceOutscaleOAPILinPeeringConnectionConfig = `
	resource "outscale_net" "net" {
		ip_range = "10.10.0.0/24"
	}

	resource "outscale_net" "net2" {
		ip_range = "10.11.0.0/24"
	}

	resource "outscale_net_peering" "net_peering" {
		accepter_net_id = "${outscale_net.net.net_id}"
		source_net_id   = "${outscale_net.net2.net_id}"
	}

	data "outscale_net_peering" "net_peering_data" {
		filter {
			name   = "net_peering_ids"
			values = ["${outscale_net_peering.net_peering.net_peering_id}"]
		}
	}
`
