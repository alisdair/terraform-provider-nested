---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "Nested Provider"
subcategory: ""
description: |-
  
---

# Nested Provider

This provider is intended to be used for testing Terraform's implementation of nested attribute types, available in [terraform-plugin-framework](https://github.com/hashicorp/terraform-plugin-framework).

## Example

```terraform
terraform {
  required_providers {
    nested = {
      source = "alisdair/nested"
    }
  }
}

resource "nested_list" "example" {
  name = "example"
  values = [
    {
      string = "foo"
      number = 1
      bool = true
    },
    {
      string = "bar"
      number = 2
      bool = false
      sensitive = "boop"
    },
    {
      string = "baz"
      number = 3
      nested = {
        string = "bazfoo"
        number = 9
        sensitive = "honk"
      }
    },
  ]
}
```