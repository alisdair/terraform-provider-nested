---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nested_single Resource - terraform-provider-nested"
subcategory: ""
description: |-
  
---

# nested_single (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `sensitive_value` (Attributes, Sensitive) (see [below for nested schema](#nestedatt--sensitive_value))
- `value` (Attributes) (see [below for nested schema](#nestedatt--value))

<a id="nestedatt--sensitive_value"></a>
### Nested Schema for `sensitive_value`

Optional:

- `bool` (Boolean)
- `nested` (Attributes) (see [below for nested schema](#nestedatt--sensitive_value--nested))
- `number` (Number)
- `sensitive` (String, Sensitive)
- `string` (String)

<a id="nestedatt--sensitive_value--nested"></a>
### Nested Schema for `sensitive_value.nested`

Optional:

- `bool` (Boolean)
- `number` (Number)
- `sensitive` (String, Sensitive)
- `string` (String)



<a id="nestedatt--value"></a>
### Nested Schema for `value`

Optional:

- `bool` (Boolean)
- `nested` (Attributes) (see [below for nested schema](#nestedatt--value--nested))
- `number` (Number)
- `sensitive` (String, Sensitive)
- `string` (String)

<a id="nestedatt--value--nested"></a>
### Nested Schema for `value.nested`

Optional:

- `bool` (Boolean)
- `number` (Number)
- `sensitive` (String, Sensitive)
- `string` (String)


