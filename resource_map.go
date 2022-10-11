package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *resourceMap) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"name": {
				Required: true,
				Type:     types.StringType,
			},
			"values": {
				Optional:   true,
				Attributes: tfsdk.MapNestedAttributes(valueAttributes),
			},
			"sensitive_values": {
				Optional:   true,
				Sensitive:  true,
				Attributes: tfsdk.MapNestedAttributes(valueAttributes),
			},
		},
	}, nil
}

type resourceMap struct{}

type Map struct {
	Name            string           `tfsdk:"name"`
	Values          map[string]Value `tfsdk:"values"`
	SensitiveValues map[string]Value `tfsdk:"sensitive_values"`
}

func newResourceMap() resource.Resource {
	return &resourceMap{}
}

func (r *resourceMap) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_map"
}

func (r *resourceMap) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var f Map
	diags := req.Plan.Get(ctx, &f)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, f)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *resourceMap) Read(_ context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
}

func (r *resourceMap) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var f Map
	diags := req.Plan.Get(ctx, &f)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, f)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *resourceMap) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
