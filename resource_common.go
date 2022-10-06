package main

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Nested struct {
	String    *string  `tfsdk:"string"`
	Number    *float64 `tfsdk:"number"`
	Bool      *bool    `tfsdk:"bool"`
	Sensitive *string  `tfsdk:"sensitive"`
}

type Value struct {
	String    *string  `tfsdk:"string"`
	Number    *float64 `tfsdk:"number"`
	Bool      *bool    `tfsdk:"bool"`
	Sensitive *string  `tfsdk:"sensitive"`
	Nested    *Nested  `tfsdk:"nested"`
}

var valueAttributes = map[string]tfsdk.Attribute{
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
}

var valueContainerAttributes = map[string]tfsdk.Attribute{
	"value": {
		Optional:   true,
		Attributes: tfsdk.SingleNestedAttributes(valueAttributes),
	},
	"sensitive_value": {
		Optional:   true,
		Sensitive:  true,
		Attributes: tfsdk.SingleNestedAttributes(valueAttributes),
	},
}
