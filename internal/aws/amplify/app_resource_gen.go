// Code generated by generators/resource/main.go; DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_amplify_app", appResourceType)
}

// appResourceType returns the Terraform awscc_amplify_app resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Amplify::App resource type.
func appResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"access_token": {
			// Property: AccessToken
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
			// AccessToken is a write-only property.
		},
		"app_id": {
			// Property: AppId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 20,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"app_name": {
			// Property: AppName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"auto_branch_creation_config": {
			// Property: AutoBranchCreationConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AutoBranchCreationPatterns": {
			//       "items": {
			//         "maxLength": 2048,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "BasicAuthConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "EnableBasicAuth": {
			//           "type": "boolean"
			//         },
			//         "Password": {
			//           "maxLength": 255,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "Username": {
			//           "maxLength": 255,
			//           "minLength": 1,
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "BuildSpec": {
			//       "maxLength": 25000,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "EnableAutoBranchCreation": {
			//       "type": "boolean"
			//     },
			//     "EnableAutoBuild": {
			//       "type": "boolean"
			//     },
			//     "EnablePerformanceMode": {
			//       "type": "boolean"
			//     },
			//     "EnablePullRequestPreview": {
			//       "type": "boolean"
			//     },
			//     "EnvironmentVariables": {
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Name": {
			//             "maxLength": 255,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "Value": {
			//             "maxLength": 5500,
			//             "pattern": "",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Name",
			//           "Value"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "PullRequestEnvironmentName": {
			//       "maxLength": 20,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "Stage": {
			//       "enum": [
			//         "EXPERIMENTAL",
			//         "BETA",
			//         "PULL_REQUEST",
			//         "PRODUCTION",
			//         "DEVELOPMENT"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"auto_branch_creation_patterns": {
						// Property: AutoBranchCreationPatterns
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayForEach(validate.StringLenBetween(1, 2048)),
						},
					},
					"basic_auth_config": {
						// Property: BasicAuthConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enable_basic_auth": {
									// Property: EnableBasicAuth
									Type:     types.BoolType,
									Optional: true,
								},
								"password": {
									// Property: Password
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 255),
									},
								},
								"username": {
									// Property: Username
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 255),
									},
								},
							},
						),
						Optional: true,
					},
					"build_spec": {
						// Property: BuildSpec
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 25000),
						},
					},
					"enable_auto_branch_creation": {
						// Property: EnableAutoBranchCreation
						Type:     types.BoolType,
						Optional: true,
					},
					"enable_auto_build": {
						// Property: EnableAutoBuild
						Type:     types.BoolType,
						Optional: true,
					},
					"enable_performance_mode": {
						// Property: EnablePerformanceMode
						Type:     types.BoolType,
						Optional: true,
					},
					"enable_pull_request_preview": {
						// Property: EnablePullRequestPreview
						Type:     types.BoolType,
						Optional: true,
					},
					"environment_variables": {
						// Property: EnvironmentVariables
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(255),
									},
								},
								"value": {
									// Property: Value
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(5500),
									},
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
					"pull_request_environment_name": {
						// Property: PullRequestEnvironmentName
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(20),
						},
					},
					"stage": {
						// Property: Stage
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"EXPERIMENTAL",
								"BETA",
								"PULL_REQUEST",
								"PRODUCTION",
								"DEVELOPMENT",
							}),
						},
					},
				},
			),
			Optional: true,
			// AutoBranchCreationConfig is a write-only property.
		},
		"basic_auth_config": {
			// Property: BasicAuthConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "EnableBasicAuth": {
			//       "type": "boolean"
			//     },
			//     "Password": {
			//       "maxLength": 255,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "Username": {
			//       "maxLength": 255,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"enable_basic_auth": {
						// Property: EnableBasicAuth
						Type:     types.BoolType,
						Optional: true,
					},
					"password": {
						// Property: Password
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
					"username": {
						// Property: Username
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
				},
			),
			Optional: true,
			// BasicAuthConfig is a write-only property.
		},
		"build_spec": {
			// Property: BuildSpec
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 25000,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 25000),
			},
		},
		"custom_headers": {
			// Property: CustomHeaders
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 25000,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 25000),
			},
		},
		"custom_rules": {
			// Property: CustomRules
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Condition": {
			//         "maxLength": 2048,
			//         "minLength": 0,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Source": {
			//         "maxLength": 2048,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Status": {
			//         "maxLength": 7,
			//         "minLength": 3,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Target": {
			//         "maxLength": 2048,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Target",
			//       "Source"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"condition": {
						// Property: Condition
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 2048),
						},
					},
					"source": {
						// Property: Source
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 2048),
						},
					},
					"status": {
						// Property: Status
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(3, 7),
						},
					},
					"target": {
						// Property: Target
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 2048),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"default_domain": {
			// Property: DefaultDomain
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(1000),
			},
		},
		"enable_branch_auto_deletion": {
			// Property: EnableBranchAutoDeletion
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
		},
		"environment_variables": {
			// Property: EnvironmentVariables
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Name": {
			//         "maxLength": 255,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 5500,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Name",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(255),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(5500),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"iam_service_role": {
			// Property: IAMServiceRole
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 1000),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
		},
		"oauth_token": {
			// Property: OauthToken
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(1000),
			},
			// OauthToken is a write-only property.
		},
		"repository": {
			// Property: Repository
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "insertionOrder": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The AWS::Amplify::App resource creates Apps in the Amplify Console. An App is a collection of branches.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Amplify::App").WithTerraformTypeName("awscc_amplify_app")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_token":                  "AccessToken",
		"app_id":                        "AppId",
		"app_name":                      "AppName",
		"arn":                           "Arn",
		"auto_branch_creation_config":   "AutoBranchCreationConfig",
		"auto_branch_creation_patterns": "AutoBranchCreationPatterns",
		"basic_auth_config":             "BasicAuthConfig",
		"build_spec":                    "BuildSpec",
		"condition":                     "Condition",
		"custom_headers":                "CustomHeaders",
		"custom_rules":                  "CustomRules",
		"default_domain":                "DefaultDomain",
		"description":                   "Description",
		"enable_auto_branch_creation":   "EnableAutoBranchCreation",
		"enable_auto_build":             "EnableAutoBuild",
		"enable_basic_auth":             "EnableBasicAuth",
		"enable_branch_auto_deletion":   "EnableBranchAutoDeletion",
		"enable_performance_mode":       "EnablePerformanceMode",
		"enable_pull_request_preview":   "EnablePullRequestPreview",
		"environment_variables":         "EnvironmentVariables",
		"iam_service_role":              "IAMServiceRole",
		"key":                           "Key",
		"name":                          "Name",
		"oauth_token":                   "OauthToken",
		"password":                      "Password",
		"pull_request_environment_name": "PullRequestEnvironmentName",
		"repository":                    "Repository",
		"source":                        "Source",
		"stage":                         "Stage",
		"status":                        "Status",
		"tags":                          "Tags",
		"target":                        "Target",
		"username":                      "Username",
		"value":                         "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/AccessToken",
		"/properties/BasicAuthConfig",
		"/properties/OauthToken",
		"/properties/AutoBranchCreationConfig",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
