---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "authentik_stage_authenticator_duo Resource - terraform-provider-authentik"
subcategory: "Flows & Stages"
description: |-

---

# authentik_stage_authenticator_duo (Resource)



## Example Usage

```terraform
# Create a duo setup stage

resource "authentik_stage_authenticator_duo" "name" {
  name          = "duo-setup"
  client_id     = "foo"
  client_secret = "bar"
  api_hostname  = "http://foo.bar.baz"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **api_hostname** (String)
- **client_id** (String)
- **client_secret** (String, Sensitive)
- **name** (String)

### Optional

- **configure_flow** (String)
- **id** (String) The ID of this resource.


