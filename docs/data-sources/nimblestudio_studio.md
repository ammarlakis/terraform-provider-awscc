---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_nimblestudio_studio Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::NimbleStudio::Studio
---

# awscc_nimblestudio_studio (Data Source)

Data Source schema for AWS::NimbleStudio::Studio



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **admin_role_arn** (String)
- **display_name** (String)
- **home_region** (String)
- **sso_client_id** (String)
- **studio_encryption_configuration** (Attributes) (see [below for nested schema](#nestedatt--studio_encryption_configuration))
- **studio_id** (String)
- **studio_name** (String)
- **studio_url** (String)
- **tags** (Map of String)
- **user_role_arn** (String)

<a id="nestedatt--studio_encryption_configuration"></a>
### Nested Schema for `studio_encryption_configuration`

Read-Only:

- **key_arn** (String)
- **key_type** (String)


