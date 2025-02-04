// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_cloudformation_public_type_version", publicTypeVersionResourceType)
}

// publicTypeVersionResourceType returns the Terraform awscc_cloudformation_public_type_version resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CloudFormation::PublicTypeVersion resource type.
func publicTypeVersionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Number (ARN) of the extension.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Number (ARN) of the extension.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"log_delivery_bucket": {
			// Property: LogDeliveryBucket
			// CloudFormation resource type schema:
			// {
			//   "description": "A url to the S3 bucket where logs for the testType run will be available",
			//   "type": "string"
			// }
			Description: "A url to the S3 bucket where logs for the testType run will be available",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"public_type_arn": {
			// Property: PublicTypeArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Number (ARN) assigned to the public extension upon publication",
			//   "maxLength": 1024,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Number (ARN) assigned to the public extension upon publication",
			Type:        types.StringType,
			Computed:    true,
		},
		"public_version_number": {
			// Property: PublicVersionNumber
			// CloudFormation resource type schema:
			// {
			//   "description": "The version number of a public third-party extension",
			//   "maxLength": 64,
			//   "minLength": 5,
			//   "type": "string"
			// }
			Description: "The version number of a public third-party extension",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(5, 64),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"publisher_id": {
			// Property: PublisherId
			// CloudFormation resource type schema:
			// {
			//   "description": "The publisher id assigned by CloudFormation for publishing in this region.",
			//   "maxLength": 40,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The publisher id assigned by CloudFormation for publishing in this region.",
			Type:        types.StringType,
			Computed:    true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "The kind of extension",
			//   "enum": [
			//     "RESOURCE",
			//     "MODULE"
			//   ],
			//   "type": "string"
			// }
			Description: "The kind of extension",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"RESOURCE",
					"MODULE",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"type_name": {
			// Property: TypeName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"type_version_arn": {
			// Property: TypeVersionArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Number (ARN) of the extension with the versionId.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Number (ARN) of the extension with the versionId.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Test and Publish a resource that has been registered in the CloudFormation Registry.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::PublicTypeVersion").WithTerraformTypeName("awscc_cloudformation_public_type_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"log_delivery_bucket":   "LogDeliveryBucket",
		"public_type_arn":       "PublicTypeArn",
		"public_version_number": "PublicVersionNumber",
		"publisher_id":          "PublisherId",
		"type":                  "Type",
		"type_name":             "TypeName",
		"type_version_arn":      "TypeVersionArn",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithRequiredAttributesValidators(validate.OneOfRequired(
		validate.Required(
			"type_name",
			"type",
		),
		validate.Required(
			"arn",
		),
	),
	)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
