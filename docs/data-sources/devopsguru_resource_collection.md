---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_devopsguru_resource_collection Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::DevOpsGuru::ResourceCollection
---

# awscc_devopsguru_resource_collection (Data Source)

Data Source schema for AWS::DevOpsGuru::ResourceCollection



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **resource_collection_filter** (Attributes) Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru. (see [below for nested schema](#nestedatt--resource_collection_filter))
- **resource_collection_type** (String) The type of ResourceCollection

<a id="nestedatt--resource_collection_filter"></a>
### Nested Schema for `resource_collection_filter`

Read-Only:

- **cloudformation** (Attributes) CloudFormation resource for DevOps Guru to monitor (see [below for nested schema](#nestedatt--resource_collection_filter--cloudformation))

<a id="nestedatt--resource_collection_filter--cloudformation"></a>
### Nested Schema for `resource_collection_filter.cloudformation`

Read-Only:

- **stack_names** (List of String) An array of CloudFormation stack names.


