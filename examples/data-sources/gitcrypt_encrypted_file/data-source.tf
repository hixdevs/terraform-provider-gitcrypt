terraform {
  required_providers {
    gitcrypt = {
      source = "hixdevs/gitcrypt"
      version = "0.0.1"
    }
  }
}

# init provider
provider "gitcrypt" {
    gitcrypt_key_base64 = "AEdJVENSWVBUS0VZAAAAAgAAAAAAAAABAAAABAAAAAAAAAADAAAAIP26dkCnQLES83htwVCWmAuBx1Kiq6llC6oDSqHTbQ/rAAAABQAAAECy6URjldOBe8HX9onc4D7bx4rizU7QScmDTWVJksb0h5JZGOpV0prhHmwedfqQE0xAvTKG4wpKD4HU1TKAqI00AAAAAA=="
    # this key can be set as an ENV variable GIT_CRYPT_KEY_BASE64 or KEY_BASE64
}

# decrypt and parse git-crypt encrypted file
data "gitcrypt_encrypted_file" "some_file" {
  envs = ["./test-data/.env", "./test-data/.env.prd", "./test-data/.env.prd.override"]
}

# use secrets to define your own resources
# outputs are used as an example in order not to import any other providers
output "var1" {
  value = data.gitcrypt_encrypted_file.some_file.secrets.var1
  # or can be used in this format:
  # value = data.git_crypt_encrypted_file.yml_variables.vars["var1"]
}

output "all_vars" {
  value = data.gitcrypt_encrypted_file.some_file.secrets
  # or can be used in this format:
  # value = data.git_crypt_encrypted_file.yml_variables.secrets[*]
}
