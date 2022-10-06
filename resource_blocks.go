package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type resourceBlocksType struct{}

func (r resourceBlocksType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Blocks: map[string]tfsdk.Block{
			"list": {
				NestingMode: tfsdk.BlockNestingModeList,
				Attributes:  valueContainerAttributes,
			},
			"set": {
				NestingMode: tfsdk.BlockNestingModeSet,
				Attributes:  valueContainerAttributes,
			},
		},
		Attributes: map[string]tfsdk.Attribute{
			"name": {
				Required: true,
				Type:     types.StringType,
			},
		},
	}, nil
}

func (r resourceBlocksType) NewResource(_ context.Context, _ tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceBlocks{}, nil
}

type resourceBlocks struct{}

type ValueContainer struct {
	Value          *Value `tfsdk:"value"`
	SensitiveValue *Value `tfsdk:"sensitive_value"`
}

type Blocks struct {
	Name  string           `tfsdk:"name"`
	Lists []ValueContainer `tfsdk:"list"`
	Sets  []ValueContainer `tfsdk:"set"`
}

func (r resourceBlocks) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var f Blocks
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

func (r resourceBlocks) Read(_ context.Context, _ tfsdk.ReadResourceRequest, _ *tfsdk.ReadResourceResponse) {
}

func (r resourceBlocks) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var f Blocks
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

func (r resourceBlocks) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	resp.State.RemoveResource(ctx)
}

func (r resourceBlocks) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStateNotImplemented(ctx, "", resp)
}
