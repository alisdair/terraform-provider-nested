package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type resourceSetType struct{}

func (r resourceSetType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"name": {
				Required: true,
				Type:     types.StringType,
			},
			"values": {
				Optional:   true,
				Attributes: tfsdk.SetNestedAttributes(valueAttributes, tfsdk.SetNestedAttributesOptions{}),
			},
			"sensitive_values": {
				Optional:   true,
				Sensitive:  true,
				Attributes: tfsdk.SetNestedAttributes(valueAttributes, tfsdk.SetNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (r resourceSetType) NewResource(_ context.Context, _ tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceSet{}, nil
}

type resourceSet struct{}

type Set struct {
	Name            string  `tfsdk:"name"`
	Values          []Value `tfsdk:"values"`
	SensitiveValues []Value `tfsdk:"sensitive_values"`
}

func (r resourceSet) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
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

func (r resourceSet) Read(_ context.Context, _ tfsdk.ReadResourceRequest, _ *tfsdk.ReadResourceResponse) {
}

func (r resourceSet) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
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

func (r resourceSet) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	resp.State.RemoveResource(ctx)
}

func (r resourceSet) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStateNotImplemented(ctx, "", resp)
}
