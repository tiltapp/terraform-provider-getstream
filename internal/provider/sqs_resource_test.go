package provider

import (
	"fmt"
)

//func TestAccExampleResource(t *testing.T) {
//	resource.Test(t, resource.TestCase{
//		PreCheck:                 func() { testAccPreCheck(t) },
//		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
//		Steps: []resource.TestStep{
//			// Create and Read testing
//			{
//				Config: testAccSqsConfig("one", "two", "three"),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					resource.TestCheckResourceAttr("getstream_sqs.test", "sqs_url", "one"),
//					resource.TestCheckResourceAttr("getstream_sqs.test", "sqs_access_key", "two"),
//					resource.TestCheckResourceAttr("getstream_sqs.test", "sqs_secret_key", "three"),
//				),
//			},
//			// ImportState testing
//			{
//				ResourceName:      "getstream_sqs.test",
//				ImportState:       true,
//				ImportStateVerify: true,
//				// This is not normally necessary, but is here because this
//				// example code does not have an actual upstream service.
//				// Once the Read method is able to refresh information from
//				// the upstream service, this can be removed.
//				ImportStateVerifyIgnore: []string{"sqs_url", "sqs_access_key", "sqs_secret_key"},
//			},
//			// Update and Read testing
//			{
//				Config: testAccSqsConfig("one", "two", "three"),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					resource.TestCheckResourceAttr("getstream_sqs.test", "sqs_url", "one"),
//					resource.TestCheckResourceAttr("getstream_sqs.test", "sqs_access_key", "two"),
//					resource.TestCheckResourceAttr("getstream_sqs.test", "sqs_secret_key", "three")),
//			},
//			// Delete testing automatically occurs in TestCase
//		},
//	})
//}

func testAccSqsConfig(sqsUrl string, sqsAccessKey string, sqsSecretKey string) string {
	return fmt.Sprintf(`
provider "getstream" {
  api_key = "test"
  api_secret = "test"
}
resource "getstream_sqs" "test" {
  sqs_url = %[1]q
  sqs_access_key = %[2]q
  sqs_secret_key = %[3]q
}
`, sqsUrl, sqsAccessKey, sqsSecretKey)
}
