package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccExampleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccExampleResourceConfig("one", "two", "three"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("getstreamio_sqs.test", "sqs_url", "one"),
					resource.TestCheckResourceAttr("getstreamio_sqs.test", "sqs_access_key", "two"),
					resource.TestCheckResourceAttr("getstreamio_sqs.test", "sqs_secret_key", "three"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "getstreamio_sqs.test",
				ImportState:       true,
				ImportStateVerify: true,
				// This is not normally necessary, but is here because this
				// example code does not have an actual upstream service.
				// Once the Read method is able to refresh information from
				// the upstream service, this can be removed.
				ImportStateVerifyIgnore: []string{"sqs_url", "sqs_access_key", "sqs_secret_key"},
			},
			// Update and Read testing
			{
				Config: testAccExampleResourceConfig("one", "two", "three"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("getstreamio_sqs.test", "sqs_url", "one"),
					resource.TestCheckResourceAttr("getstreamio_sqs.test", "sqs_access_key", "two"),
					resource.TestCheckResourceAttr("getstreamio_sqs.test", "sqs_secret_key", "three")),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccExampleResourceConfig(sqsUrl string, sqsAccessKey string, sqsSecretKey string) string {
	return fmt.Sprintf(`
provider "getstreamio" {
  api_key  = "123"
  api_secret = "123"
}
resource "getstreamio_sqs" "test" {
  sqs_url = %[1]q
  sqs_access_key = %[2]q
  sqs_secret_key = %[3]q
}
`, sqsUrl, sqsAccessKey, sqsSecretKey)
}
