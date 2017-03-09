# EGI AAI + OpenStack Terraform Plugin

This is a Terraform provider plugin for OpenStack resources of the EGI Federated Cloud.
It has the same features of the internal OpenStack provider of Terraform, but adds a
new authentication method based on VOMS X.509 proxies

## Using the plugin

In order to avoid changes in existing deployments, the easiest way to use the plugin is
to override the internal OpenStack provider with the EGI one. For that you can add these
lines to your `$HOME/.terraformrc` file:

```
providers {
    openstack = "terraform-provider-egi-openstack"
}
```

And when configuring the provider add the `voms = true` option and your X.509 proxy in the
`cert`  and `key` options like this:
```
provider "openstack" {
    auth_url = "https://keystone.site.com:5000/v2.0"
    tenant_id = "8f5c5d426a164478b1c659965a0a3dfd"
    cert = "/tmp/x509up_u1000"
    key = "/tmp/x509up_u1000"
    voms = true
}
```

## Building the plugin

From a working golang environment:

- Get the code
```
go get github.com/enolfc/egi-openstack-terraform
cd $GOPATH/src/github.com/enolfc/egi-openstack-terraform
```
- Build the plugin (`dev` will just build for your current platform, use `bin` to build
  for other platforms)
```
make dev
```
- Put the binary on your `$PATH` and you are ready to go!
