package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareAccountMemberSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"email_address": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The email address of the user who you wish to manage. Following creation, this field becomes read only via the API and cannot be updated.",
		},

		"role_ids": {
			Type:        schema.TypeSet,
			Required:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of account role IDs that you want to assign to a member.",
		},
	}
}
