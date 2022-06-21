package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceEncryptedFile(t *testing.T) {

	resource.UnitTest(t, resource.TestCase{
		// PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceEncryptedFile,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.gitcrypt_encrypted_file.example", "envs", regexp.MustCompile("^.*")),
					resource.TestCheckResourceAttr("data.gitcrypt_encrypted_file.example", "secrets.var1", "value1"),
					resource.TestCheckResourceAttr("data.gitcrypt_encrypted_file.example", "secrets.var2", "prd_value2"),
					resource.TestCheckResourceAttr("data.gitcrypt_encrypted_file.example", "secrets.var3", "prd_override_value3"),
					resource.TestCheckResourceAttr("data.gitcrypt_encrypted_file.example", "secrets.var4", "prd_override_value4"),
				),
			},
		},
	})
}

const testAccDataSourceEncryptedFile = `
provider "gitcrypt" {
    gitcrypt_key_base64 = "AEdJVENSWVBUS0VZAAAAAgAAAAAAAAABAAAABAAAAAAAAAADAAAAIP26dkCnQLES83htwVCWmAuBx1Kiq6llC6oDSqHTbQ/rAAAABQAAAECy6URjldOBe8HX9onc4D7bx4rizU7QScmDTWVJksb0h5JZGOpV0prhHmwedfqQE0xAvTKG4wpKD4HU1TKAqI00AAAAAA=="
}

data "gitcrypt_encrypted_file" "example" {
  envs = ["./test-data/.env", "./test-data/.env.prd", "./test-data/.env.prd.override"]
}
`
