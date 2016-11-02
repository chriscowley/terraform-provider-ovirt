package ovirt

import (
    "github.com/hashicorp/terraform/helper/schema"
//    "log"
)

func resourceOVirtVM() *schema.Resource {
    return &schema.Resource{
        Create: resourceOVirtVMCreate,
        Read:   resourceOVirtVMRead,
        Update: resourceOVirtVMUpdate,
        Delete: resourceOVirtVMDelete,

        Schema: map[string]*schema.Schema{
            "name": &schema.Schema{
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
            },
            "template": &schema.Schema{
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
            },
            "vcpus": &schema.Schema{
                Type: schema.TypeString,
                Required: false,
                ForceNew: true,
                Default: 1,
            },
            "nic": &schema.Schema{
                Type: schema.TypeList,
                Required: false,
                ForceNew: true,
                Elem: &schema.Resource{
                    "profile": &schema.Schema{
                        Type: schema.TypeString,
                        Required: true
                    },
                },
            },
            "volumes": &schema.Schema{
                Type: schema.TypeList,
                Required: false,
                ForceNew: true,
                Elem: &schema.Resource{
                    "size": &schema.Schema{
                        Type: schema.TypeInt,
                        required: true,
                    },
                    "interface": &schema.Schema{
                        Type: schema.TypeString,
                        Optional: true,
                        Default: "VirtIO",
                        ForceNew: true,
                    },
                    "domain": &schema.Schema{
                        Type: schema.TypeString,
                        Required: true,
                    },
                    "thin": &schema.Schema{
                        Type: schema.TypeBool,
                        Optional: true,
                        Default: true,
                        ForceNew: true,
                    },
                    "profile": &schema.Schema{
                        Type: schema.TypeString,
                        Required: true,
                    },
                },
            },
        },
    }
}

func resourceOVirtVMCreate(d *schema.ResourceData,m interface{}) error {
    return nil
}

func resourceOVirtVMRead(d *schema.ResourceData,m interface{}) error {
    return nil
}

func resourceOVirtVMUpdate(d *schema.ResourceData,m interface{}) error {
    return nil
}

func resourceOVirtVMDelete(d *schema.ResourceData,m interface{}) error {
    return nil
}
