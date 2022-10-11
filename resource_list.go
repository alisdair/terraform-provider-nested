package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func newResourceList() resource.Resource {
	return &resourceList{}
}

func (r *resourceList) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"name": {
				Required: true,
				Type:     types.StringType,
			},
			"values": {
				Optional:   true,
				Attributes: tfsdk.ListNestedAttributes(valueAttributes),
			},
			"sensitive_values": {
				Optional:   true,
				Sensitive:  true,
				Attributes: tfsdk.ListNestedAttributes(valueAttributes),
			},
		},
	}, nil
}

func (r *resourceList) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_list"
}

type resourceList struct{}

type List struct {
	Name            string  `tfsdk:"name"`
	Values          []Value `tfsdk:"values"`
	SensitiveValues []Value `tfsdk:"sensitive_values"`
}

func (r resourceList) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var f List
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

func (r *resourceList) Read(_ context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
}

func (r *resourceList) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var f List
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

func (r *resourceList) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
