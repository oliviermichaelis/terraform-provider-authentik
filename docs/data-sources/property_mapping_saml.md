---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "authentik_property_mapping_saml Data Source - terraform-provider-authentik"
subcategory: "Property Mappings"
description: |-
  Get SAML Property mappings
---

# authentik_property_mapping_saml (Data Source)

Get SAML Property mappings

## Example Usage

```terraform
# To get the ID of a SAML Property mapping

data "authentik_property_mapping_saml" "test" {
  managed = "goauthentik.io/providers/saml/upn"
}

# Then use `data.authentik_property_mapping_saml.test.id`

# Or, to get the IDs of multiple mappings

data "authentik_property_mapping_saml" "test" {
  managed_list = [
    "goauthentik.io/providers/saml/upn",
    "goauthentik.io/providers/saml/name"
  ]
}

# Then use data.authentik_property_mapping_saml.test.ids
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **expression** (String)
- **friendly_name** (String)
- **id** (String) The ID of this resource.
- **ids** (List of String) List of ids when `managed_list` is set.
- **managed** (String)
- **managed_list** (List of String) Retrive multiple property mappings
- **name** (String)
- **saml_name** (String)


