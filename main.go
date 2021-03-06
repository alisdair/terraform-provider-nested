package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func main() {
	if err := tfsdk.Serve(context.Background(), func() tfsdk.Provider { return &provider{} }, tfsdk.ServeOpts{
		Name: "nested",
	}); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start provider: %v\n", err)
		os.Exit(1)
	}
}

type provider struct{}

func (p *provider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"dummy": {
				Type:     types.StringType,
				Optional: true,
			},
		},
	}, nil
}

func (p *provider) Configure(_ context.Context, _ tfsdk.ConfigureProviderRequest, _ *tfsdk.ConfigureProviderResponse) {
}

func (p *provider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"nested_list":   resourceListType{},
		"nested_map":    resourceMapType{},
		"nested_set":    resourceSetType{},
		"nested_single": resourceSingleType{},
	}, nil
}

func (p *provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{}, nil
}
