package outscale

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-outscale/osc/fcu"
)

func TestAccOutscaleImageTask_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOutscaleImageTaskDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccOutscaleImageTaskConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleImageTaskExists("outscale_image_task.test"),
				),
			},
		},
	})
}

func testAccCheckOutscaleImageTaskDestroy(s *terraform.State) error {

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "outscale_image_register" {
			continue
		}
		amiId := rs.Primary.ID
		conn := testAccProvider.Meta().(*OutscaleClient).FCU
		diReq := &fcu.DescribeImagesInput{
			ImageIds: []*string{aws.String(amiId)},
		}

		var diRes *fcu.DescribeImagesOutput

		err := resource.Retry(5*time.Minute, func() *resource.RetryError {
			var err error
			diRes, err = conn.VM.DescribeImages(diReq)
			if err != nil {
				if strings.Contains(err.Error(), "RequestLimitExceeded:") {
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})

		if err != nil {
			if strings.Contains(fmt.Sprint(err), "InvalidAMIID.NotFound") {
				return nil
			}
			return fmt.Errorf("[DEBUG TES] Error register image %s", err)
		}

	}

	return nil
}

func testAccCheckOutscaleImageTaskExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Role name is set")
		}

		return nil
	}
}

var testAccOutscaleImageTaskConfig = `
resource "outscale_vm" "outscale_vm" {
    count = 1

    image_id                    = "ami-880caa66"
    instance_type               = "c4.large"
    key_name                    = "integ_sut_keypair"
    security_group              = ["sg-c73d3b6b"]

    provisioner "local-exec" {
        command = "date; who -b"
    }
}

resource "outscale_image" "outscale_image" {
    name            = "image_${outscale_vm.outscale_vm.id}"
    instance_id     = "${outscale_vm.outscale_vm.id}"
    #no_reboot      = "false"                 # default value
}

resource "outscale_image_tasks" "outscale_image_tasks" {
    count = 1

    image_id = "${outscale_image.outscale_image.image_id}"
}
`
