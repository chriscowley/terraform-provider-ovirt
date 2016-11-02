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
