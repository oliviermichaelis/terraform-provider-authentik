---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "authentik_stage_prompt Resource - terraform-provider-authentik"
subcategory: "Flows & Stages"
description: |-

---

# authentik_stage_prompt (Resource)



## Example Usage

```terraform
# Create a prompt stage with 1 field

resource "authentik_stage_prompt_field" "field" {
  field_key = "username"
  label     = "Username"
  type      = "username"
}
resource "authentik_stage_prompt" "name" {
  name = "test"
  fields = [
    resource.authentik_stage_prompt_field.field.id,
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **fields** (List of String)
- **name** (String)

### Optional

- **id** (String) The ID of this resource.
- **validation_policies** (List of String)


