package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *resourceSet) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"name": {
				Required: true,
				Type:     types.StringType,
			},
			"values": {
				Optional:   true,
				Attributes: tfsdk.SetNestedAttributes(valueAttributes),
			},
			"sensitive_values": {
				Optional:   true,
				Sensitive:  true,
				Attributes: tfsdk.SetNestedAttributes(valueAttributes),
			},
		},
	}, nil
}

type resourceSet struct{}

type Set struct {
	Name            string  `tfsdk:"name"`
	Values          []Value `tfsdk:"values"`
	SensitiveValues []Value `tfsdk:"sensitive_values"`
}

func newResourceSet() resource.Resource {
	return &resourceSet{}
}

func (r *resourceSet) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_set"
}

func (r *resourceSet) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var f Set
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

func (r *resourceSet) Read(_ context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
}

func (r *resourceSet) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var f Set
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

func (r *resourceSet) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
