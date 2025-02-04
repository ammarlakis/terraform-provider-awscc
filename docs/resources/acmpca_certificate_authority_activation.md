---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_acmpca_certificate_authority_activation Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Used to install the certificate authority certificate and update the certificate authority status.
---

# awscc_acmpca_certificate_authority_activation (Resource)

Used to install the certificate authority certificate and update the certificate authority status.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **certificate** (String) Certificate Authority certificate that will be installed in the Certificate Authority.
- **certificate_authority_arn** (String) Arn of the Certificate Authority.

### Optional

- **certificate_chain** (String) Certificate chain for the Certificate Authority certificate.
- **status** (String) The status of the Certificate Authority.

### Read-Only

- **complete_certificate_chain** (String) The complete certificate chain, including the Certificate Authority certificate.
- **id** (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_acmpca_certificate_authority_activation.example <resource ID>
```
