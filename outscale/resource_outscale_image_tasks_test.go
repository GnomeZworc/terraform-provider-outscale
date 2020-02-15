package outscale

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccOutscaleOAPIImageTask_basic(t *testing.T) {
	t.Skip()

	omi := os.Getenv("OUTSCALE_IMAGEID")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccOutscaleOAPIImageTaskConfig(omi),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleOAPIImageTaskExists("outscale_image_tasks.outscale_image_tasks"),
				),
			},
		},
	})
}

func testAccCheckOutscaleOAPIImageTaskExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No image task id is set")
		}

		return nil
	}
}

func testAccOutscaleOAPIImageTaskConfig(omi string) string {
	return fmt.Sprintf(`
	resource "outscale_vm" "outscale_vm" {
		count = 1

		image_id = "%s"
		type     = "c4.large"
	}

	resource "outscale_image" "outscale_image" {
		name  = "image_${outscale_vm.outscale_vm.id}"
		vm_id = "${outscale_vm.outscale_vm.id}"
	}

	resource "outscale_image_tasks" "outscale_image_tasks" {
		count = 1

		osu_export {
			disk_image_format = "raw"
			osu_bucket        = "test"
		}

		image_id = "${outscale_image.outscale_image.image_id}"
	}
`, omi)
}
