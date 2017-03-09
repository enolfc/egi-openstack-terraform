package main

import (
	"github.com/enolfc/egi-openstack-terraform/openstack"
	"github.com/hashicorp/terraform/plugin"
)

// The main program entry-point.
func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: openstack.Provider,
	})
}
