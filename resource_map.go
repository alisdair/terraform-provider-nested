package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type resourceMapType struct{}

func (r resourceMapType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"name": {
				Required: true,
				Type:     types.StringType,
			},
			"values": {
				Optional:   true,
				Attributes: tfsdk.MapNestedAttributes(valueAttributes, tfsdk.MapNestedAttributesOptions{}),
			},
			"sensitive_values": {
				Optional:   true,
				Sensitive:  true,
				Attributes: tfsdk.MapNestedAttributes(valueAttributes, tfsdk.MapNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (r resourceMapType) NewResource(_ context.Context, _ tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceMap{}, nil
}

type resourceMap struct{}

type Map struct {
	Name            string           `tfsdk:"name"`
	Values          map[string]Value `tfsdk:"values"`
	SensitiveValues map[string]Value `tfsdk:"sensitive_values"`
}

func (r resourceMap) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
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

func (r resourceMap) Read(_ context.Context, _ tfsdk.ReadResourceRequest, _ *tfsdk.ReadResourceResponse) {
}

func (r resourceMap) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
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

func (r resourceMap) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	resp.State.RemoveResource(ctx)
}

func (r resourceMap) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStateNotImplemented(ctx, "", resp)
}
