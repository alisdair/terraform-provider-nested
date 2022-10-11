package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var valueContainerBlocks = map[string]tfsdk.Block{
	"nested_list": {
		NestingMode: tfsdk.BlockNestingModeList,
		Attributes:  valueContainerAttributes,
		Blocks:      valueContainerBlocks1,
	},
	"nested_set": {
		NestingMode: tfsdk.BlockNestingModeSet,
		Attributes:  valueContainerAttributes,
		Blocks:      valueContainerBlocks1,
	},
}

var valueContainerBlocks1 = map[string]tfsdk.Block{
	"nested_list": {
		NestingMode: tfsdk.BlockNestingModeList,
		Attributes:  valueContainerAttributes,
		Blocks:      valueContainerBlocks2,
	},
	"nested_set": {
		NestingMode: tfsdk.BlockNestingModeSet,
		Attributes:  valueContainerAttributes,
		Blocks:      valueContainerBlocks2,
	},
}

var valueContainerBlocks2 = map[string]tfsdk.Block{
	"nested_list": {
		NestingMode: tfsdk.BlockNestingModeList,
		Attributes:  valueContainerAttributes,
		Blocks:      valueContainerBlocks3,
	},
	"nested_set": {
		NestingMode: tfsdk.BlockNestingModeSet,
		Attributes:  valueContainerAttributes,
		Blocks:      valueContainerBlocks3,
	},
}

var valueContainerBlocks3 = map[string]tfsdk.Block{
	"nested_list": {
		NestingMode: tfsdk.BlockNestingModeList,
		Attributes:  valueContainerAttributes,
		Blocks:      valueContainerBlocks4,
	},
	"nested_set": {
		NestingMode: tfsdk.BlockNestingModeSet,
		Attributes:  valueContainerAttributes,
		Blocks:      valueContainerBlocks4,
	},
}

var valueContainerBlocks4 = map[string]tfsdk.Block{
	"nested_list": {
		NestingMode: tfsdk.BlockNestingModeList,
		Attributes:  valueContainerAttributes,
	},
	"nested_set": {
		NestingMode: tfsdk.BlockNestingModeSet,
		Attributes:  valueContainerAttributes,
	},
}

func (r *resourceBlocks) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Blocks: map[string]tfsdk.Block{
			"list": {
				NestingMode: tfsdk.BlockNestingModeList,
				Attributes:  valueContainerAttributes,
				Blocks:      valueContainerBlocks,
			},
			"set": {
				NestingMode: tfsdk.BlockNestingModeSet,
				Attributes:  valueContainerAttributes,
				Blocks:      valueContainerBlocks,
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

type ValueContainerNested1 struct {
	NestedLists    []ValueContainerNested2 `tfsdk:"nested_list"`
	NestedSets     []ValueContainerNested2 `tfsdk:"nested_set"`
	Value          *Value                  `tfsdk:"value"`
	SensitiveValue *Value                  `tfsdk:"sensitive_value"`
}

type ValueContainerNested2 struct {
	NestedLists    []ValueContainerNested3 `tfsdk:"nested_list"`
	NestedSets     []ValueContainerNested3 `tfsdk:"nested_set"`
	Value          *Value                  `tfsdk:"value"`
	SensitiveValue *Value                  `tfsdk:"sensitive_value"`
}

type ValueContainerNested3 struct {
	NestedLists    []ValueContainerNested4 `tfsdk:"nested_list"`
	NestedSets     []ValueContainerNested4 `tfsdk:"nested_set"`
	Value          *Value                  `tfsdk:"value"`
	SensitiveValue *Value                  `tfsdk:"sensitive_value"`
}

type ValueContainerNested4 struct {
	Value          *Value `tfsdk:"value"`
	SensitiveValue *Value `tfsdk:"sensitive_value"`
}

type ValueContainer struct {
	NestedLists    []ValueContainerNested1 `tfsdk:"nested_list"`
	NestedSets     []ValueContainerNested1 `tfsdk:"nested_set"`
	Value          *Value                  `tfsdk:"value"`
	SensitiveValue *Value                  `tfsdk:"sensitive_value"`
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
