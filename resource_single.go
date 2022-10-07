package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *resourceSingle) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"name": {
				Required: true,
				Type:     types.StringType,
			},
			"value": {
				Optional:   true,
				Attributes: tfsdk.SingleNestedAttributes(valueAttributes),
			},
			"sensitive_value": {
				Optional:   true,
				Sensitive:  true,
				Attributes: tfsdk.SingleNestedAttributes(valueAttributes),
			},
		},
	}, nil
}

type resourceSingle struct{}

func newResourceSingle() resource.Resource {
	return &resourceSingle{}
}

type Single struct {
	Name           string `tfsdk:"name"`
	Value          *Value `tfsdk:"value"`
	SensitiveValue *Value `tfsdk:"sensitive_value"`
}

func (r *resourceSingle) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_single"
}

func (r *resourceSingle) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var f Single
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

func (r *resourceSingle) Read(_ context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
}

func (r *resourceSingle) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var f Single
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

func (r *resourceSingle) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
