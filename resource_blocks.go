package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *resourceBlocks) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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

func newResourceBlock() resource.Resource {
	return &resourceBlocks{}
}

func (r *resourceBlocks) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_blocks"
}

func (r *resourceBlocks) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// NOOP
}

func (r *resourceBlocks) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
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

func (r *resourceBlocks) Read(_ context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
}

func (r *resourceBlocks) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
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

func (r *resourceBlocks) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
