package provider

import (
	"context"
	"os"

	gc "github.com/hixdevs/terraform-provider-gitcrypt/gitcrypt/internal/gitcrypt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joho/godotenv"
)

func dataSourceEncryptedFile() *schema.Resource {
	return &schema.Resource{
		// Data source for read and decrypt file encrypted by git-crypt
		Description: "Data source for read and decrypt file encrypted by git-crypt.",
		ReadContext: dataSourceEncryptedFileRead,
		Schema: map[string]*schema.Schema{
			"envs": {
				Description: "Paths to .env files encrypted by git-crypt",
				Required:    true,
				Type:        schema.TypeList,
				Elem:        schema.Schema{
					Type: schema.TypeString,
				},
			},
			"secrets": &schema.Schema{
				// Variables from ecrypted file after decryption
				Description: "Variables from ecrypted file after decryption",
				Type:        schema.TypeMap,
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

func dataSourceEncryptedFileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secrets := make(map[string]string)
	gitcryptKey := meta.(*gc.KeyData)
	paths := d.Get("envs").([]string)
	// raw := d.Get("envs").([]interface{})
	// paths := make([]string, len(raw))
	// for i, v := range raw {
	// 	l[i] = v.(string)
	// }
	// paths := d.Get("envs").([]string{})

	for _, path := range paths {
		variables := make(map[string]string)
		file, err := os.Open(path)
		if err != nil {
			return diag.FromErr(err)
		}
		switch pathType := file.(type) {
		case []byte:
			decrypted, err := gc.UnlockFile(path, *gitcryptKey)
			if err != nil {
				return diag.FromErr(err)
			}
			variables, err := godotenv.Parse(decrypted)
			if err != nil {
				return diag.FromErr(err)
			}
		case string:
			variables, err := godotenv.Read(path)
			if err != nil {
				return diag.FromErr(err)
			}
		}

		for variable, value := range variables {
			secrets[variable] = value
		}
	}
	
	err = d.Set("secrets", secrets)
	if err != nil {
		return diag.FromErr(err)
	}

	fileHMAC, err := gc.GetFileHMAC(filePath)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fileHMAC)

	return nil
}
