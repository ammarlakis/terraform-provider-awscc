// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package nimblestudio

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_nimblestudio_streaming_image", streamingImageDataSourceType)
}

// streamingImageDataSourceType returns the Terraform awscc_nimblestudio_streaming_image data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::NimbleStudio::StreamingImage resource type.
func streamingImageDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"ec_2_image_id": {
			// Property: Ec2ImageId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"encryption_configuration": {
			// Property: EncryptionConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "KeyArn": {
			//       "type": "string"
			//     },
			//     "KeyType": {
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "KeyType"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"key_arn": {
						// Property: KeyArn
						Type:     types.StringType,
						Computed: true,
					},
					"key_type": {
						// Property: KeyType
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"eula_ids": {
			// Property: EulaIds
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"owner": {
			// Property: Owner
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"platform": {
			// Property: Platform
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"streaming_image_id": {
			// Property: StreamingImageId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"studio_id": {
			// Property: StudioId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::NimbleStudio::StreamingImage",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NimbleStudio::StreamingImage").WithTerraformTypeName("awscc_nimblestudio_streaming_image")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":              "Description",
		"ec_2_image_id":            "Ec2ImageId",
		"encryption_configuration": "EncryptionConfiguration",
		"eula_ids":                 "EulaIds",
		"key_arn":                  "KeyArn",
		"key_type":                 "KeyType",
		"name":                     "Name",
		"owner":                    "Owner",
		"platform":                 "Platform",
		"streaming_image_id":       "StreamingImageId",
		"studio_id":                "StudioId",
		"tags":                     "Tags",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
