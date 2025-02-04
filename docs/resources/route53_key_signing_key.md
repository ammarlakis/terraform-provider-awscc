---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_route53_key_signing_key Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Represents a key signing key (KSK) associated with a hosted zone. You can only have two KSKs per hosted zone.
---

# awscc_route53_key_signing_key (Resource)

Represents a key signing key (KSK) associated with a hosted zone. You can only have two KSKs per hosted zone.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **hosted_zone_id** (String) The unique string (ID) used to identify a hosted zone.
- **key_management_service_arn** (String) The Amazon resource name (ARN) for a customer managed key (CMK) in AWS Key Management Service (KMS). The KeyManagementServiceArn must be unique for each key signing key (KSK) in a single hosted zone.
- **name** (String) An alphanumeric string used to identify a key signing key (KSK). Name must be unique for each key signing key in the same hosted zone.
- **status** (String) A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE.

### Read-Only

- **id** (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_route53_key_signing_key.example <resource ID>
```
