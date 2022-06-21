# Gitcrypt Provider

The gitcrypt provider is used to read files encrypted with [git-crypt](https://github.com/AGWA/git-crypt#git-crypt---transparent-file-encryption-in-git).

-> Keep your terraform state secure. Remember that all secrets which you use to define other resources will be stored in the terraform state as plain text.

## Example Usage

```hcl
# Configure the gitcrypt Provider
provider "gitcrypt" {
}

# Read encrypted .env files:
data "gitcrypt_encrypted_file" "example" {
  envs = ["./test-data/.env", "./test-data/.env.prd", "./test-data/.env.prd.override"]
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `gitcrypt_key_base64` - (Required) A git-crypt key for repository.
  Can be set as the `GIT_CRYPT_KEY_BASE64` or `KEY_BASE64` environment variable.

To get the value for argument `gitcrypt_key_base64`, you need to get a git-crypt key from your repository in base64 format:
```bash
    $ cd ./your-repository/
    $ git-crypt unlock
    $ base64 -i .git/git-crypt/keys/default
```
Output example:
```
AEdJVENSWVBUS0VZAAAAAgAAAAAAAAABAAAABAAAAAAAAAADAAAAIP26dkCnQLES83htwVCWmAuBx1Kiq6llC6oDSqHTbQ/rAAAABQAAAECy6URjldOBe8HX9onc4D7bx4rizU7QScmDTWVJksb0h5JZGOpV0prhHmwedfqQE0xAvTKG4wpKD4HU1TKAqI00AAAAAA==
```

!> It is strongly NOT recommended to set parameter `gitcrypt_key_base64` in the open. It is dangerous and not secure.
Anyone who knows the value of argument `gitcrypt_key_base64` can decrypt secret files.
