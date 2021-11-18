package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type resourceSingleType struct{}

func (r resourceSingleType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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

func (r resourceSingleType) NewResource(_ context.Context, _ tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceSingle{}, nil
}

type resourceSingle struct{}

type Single struct {
	Name           string `tfsdk:"name"`
	Value          *Value `tfsdk:"value"`
	SensitiveValue *Value `tfsdk:"sensitive_value"`
}

func (r resourceSingle) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
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

func (r resourceSingle) Read(_ context.Context, _ tfsdk.ReadResourceRequest, _ *tfsdk.ReadResourceResponse) {
}

func (r resourceSingle) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
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

func (r resourceSingle) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	resp.State.RemoveResource(ctx)
}

func (r resourceSingle) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStateNotImplemented(ctx, "", resp)
}
