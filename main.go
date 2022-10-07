package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func main() {
	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/alisdair/nested",
	}

	err := providerserver.Serve(context.Background(), func() provider.Provider {
		return &nestedProvider{}
	}, opts)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

type nestedProvider struct{}

var _ provider.Provider = &nestedProvider{}

func (p *nestedProvider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"dummy": {
				Type:     types.StringType,
				Optional: true,
			},
		},
	}, nil
}

func (p *nestedProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "nested"
}

func (p *nestedProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
}

func (p *nestedProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		newResourceBlock,
		newResourceList,
		newResourceMap,
		newResourceSet,
		newResourceSingle,
	}
}

func (p *nestedProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}
