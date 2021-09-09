// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package networkfirewall_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSNetworkFirewallLoggingConfigurationDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::NetworkFirewall::LoggingConfiguration", "awscc_networkfirewall_logging_configuration", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSNetworkFirewallLoggingConfigurationDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::NetworkFirewall::LoggingConfiguration", "awscc_networkfirewall_logging_configuration", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
