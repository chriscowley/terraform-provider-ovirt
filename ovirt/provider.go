package ovirt

import (
        "github.com/hashicorp/terraform/helper/schema"
        "github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
        return &schema.Provider{
            Schema: map[string]*schema.Schema{
                "ovirt_url": &schema.Schema{
                    Type: schema.TypeString,
                    Optional: false,
                },
                "ovirt_user": &schema.Schema{
                    Type: schema.TypeString,
                    Optional: false,
                    Default: "admin",
                },
                "ovirt_password": &schema.Schema{
                    Type: schema.TypeString,
                    Optional: false,
                },
            },

            ResourcesMap: map[string]*schema.Resource{
                "ovirt_vm": resourceOVirtVM(),
            },

            ConfigureFunc: configureProvider,
        }
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
    config := Config{
        OVirtURL:      d.Get("ovirt_url").(string),
        OVirtUser:     d.Get("ovirt_user").(string),
        OVirtPassword: d.Get("ovirt_password").(string),
    }

    return &config, nil
}

