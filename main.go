package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/chriscowley/terraform-provider-ovirt/ovirt"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: ovirt.Provider,
        })
}
