// Code generated by generators/resource/main.go; DO NOT EDIT.

package sso

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_sso_assignment", assignmentResourceType)
}

// assignmentResourceType returns the Terraform awscc_sso_assignment resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SSO::Assignment resource type.
func assignmentResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"instance_arn": {
			// Property: InstanceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The sso instance that the permission set is owned.",
			//   "maxLength": 1224,
			//   "minLength": 10,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The sso instance that the permission set is owned.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(10, 1224),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"permission_set_arn": {
			// Property: PermissionSetArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The permission set that the assignemt will be assigned",
			//   "maxLength": 1224,
			//   "minLength": 10,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The permission set that the assignemt will be assigned",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(10, 1224),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"principal_id": {
			// Property: PrincipalId
			// CloudFormation resource type schema:
			// {
			//   "description": "The assignee's identifier, user id/group id",
			//   "maxLength": 47,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The assignee's identifier, user id/group id",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 47),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"principal_type": {
			// Property: PrincipalType
			// CloudFormation resource type schema:
			// {
			//   "description": "The assignee's type, user/group",
			//   "enum": [
			//     "USER",
			//     "GROUP"
			//   ],
			//   "type": "string"
			// }
			Description: "The assignee's type, user/group",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"USER",
					"GROUP",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"target_id": {
			// Property: TargetId
			// CloudFormation resource type schema:
			// {
			//   "description": "The account id to be provisioned.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The account id to be provisioned.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"target_type": {
			// Property: TargetType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of resource to be provsioned to, only aws account now",
			//   "enum": [
			//     "AWS_ACCOUNT"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of resource to be provsioned to, only aws account now",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"AWS_ACCOUNT",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for SSO assignmet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSO::Assignment").WithTerraformTypeName("awscc_sso_assignment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"instance_arn":       "InstanceArn",
		"permission_set_arn": "PermissionSetArn",
		"principal_id":       "PrincipalId",
		"principal_type":     "PrincipalType",
		"target_id":          "TargetId",
		"target_type":        "TargetType",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
