package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type resourceListType struct{}

func (r resourceListType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"name": {
				Required: true,
				Type:     types.StringType,
			},
			"values": {
				Optional: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"string": {
						Type:     types.StringType,
						Optional: true,
					},
					"number": {
						Type:     types.NumberType,
						Optional: true,
					},
					"bool": {
						Type:     types.BoolType,
						Optional: true,
					},
					"sensitive": {
						Type:      types.StringType,
						Optional:  true,
						Sensitive: true,
					},
					"nested": {
						Optional: true,
						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
							"string": {
								Type:     types.StringType,
								Optional: true,
							},
							"number": {
								Type:     types.NumberType,
								Optional: true,
							},
							"bool": {
								Type:     types.BoolType,
								Optional: true,
							},
							"sensitive": {
								Type:      types.StringType,
								Optional:  true,
								Sensitive: true,
							},
						}),
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
			"sensitive_values": {
				Optional:  true,
				Sensitive: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"string": {
						Type:     types.StringType,
						Optional: true,
					},
					"number": {
						Type:     types.NumberType,
						Optional: true,
					},
					"bool": {
						Type:     types.BoolType,
						Optional: true,
					},
					"sensitive": {
						Type:      types.StringType,
						Optional:  true,
						Sensitive: true,
					},
					"nested": {
						Optional: true,
						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
							"string": {
								Type:     types.StringType,
								Optional: true,
							},
							"number": {
								Type:     types.NumberType,
								Optional: true,
							},
							"bool": {
								Type:     types.BoolType,
								Optional: true,
							},
							"sensitive": {
								Type:      types.StringType,
								Optional:  true,
								Sensitive: true,
							},
						}),
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (r resourceListType) NewResource(_ context.Context, _ tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceList{}, nil
}

type resourceList struct{}

type Nested struct {
	String    string  `tfsdk:"string"`
	Number    float64 `tfsdk:"number"`
	Bool      bool    `tfsdk:"bool"`
	Sensitive string  `tfsdk:"sensitive"`
}

type Value struct {
	String    string  `tfsdk:"string"`
	Number    float64 `tfsdk:"number"`
	Bool      bool    `tfsdk:"bool"`
	Sensitive string  `tfsdk:"sensitive"`
	Nested    Nested  `tfsdk:"nested"`
}

type List struct {
	Name            string  `tfsdk:"name"`
	Values          []Value `tfsdk:"values"`
	SensitiveValues []Value `tfsdk:"sensitive_values"`
}

func (r resourceList) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
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

func (r resourceList) Read(_ context.Context, _ tfsdk.ReadResourceRequest, _ *tfsdk.ReadResourceResponse) {
}

func (r resourceList) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
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

func (r resourceList) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	resp.State.RemoveResource(ctx)
}

func (r resourceList) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStateNotImplemented(ctx, "", resp)
}
