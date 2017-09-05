package azurerm

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceAzureRMFirstRunConfig_Daily_Today(t *testing.T) {
	dataSourceName := "data.azurerm_first_run_config.test"

	now := time.Now().UTC()
	config := testAccDataSourceAzureRMFirstRunConfig_Daily_Today(now)

	expectedTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+1, 0, 0, 0, time.UTC)
	formattedExpectedTime := expectedTime.Format(time.RFC3339)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "hour", strconv.Itoa(now.Hour()+1)),
					resource.TestCheckResourceAttr(dataSourceName, "minute", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "second", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "frequency", "Day"),
					resource.TestCheckResourceAttr(dataSourceName, "first_run_time", formattedExpectedTime),
				),
			},
		},
	})
}

func TestAccDataSourceAzureRMFirstRunConfig_Daily_Tomorrow(t *testing.T) {
        dataSourceName := "data.azurerm_first_run_config.test"

        now := time.Now().UTC()
        config := testAccDataSourceAzureRMFirstRunConfig_Daily_Tomorrow(now)

        expectedTime := time.Date(now.Year(), now.Month(), now.Day()+1, now.Hour()-1, 0, 0, 0, time.UTC)
        formattedExpectedTime := expectedTime.Format(time.RFC3339)

        resource.Test(t, resource.TestCase{
                PreCheck:  func() { testAccPreCheck(t) },
                Providers: testAccProviders,
                Steps: []resource.TestStep{
                        {
                                Config: config,
                                Check: resource.ComposeTestCheckFunc(
                                        resource.TestCheckResourceAttr(dataSourceName, "hour", strconv.Itoa(now.Hour()-1)),
                                        resource.TestCheckResourceAttr(dataSourceName, "minute", "0"),
                                        resource.TestCheckResourceAttr(dataSourceName, "second", "0"),
                                        resource.TestCheckResourceAttr(dataSourceName, "frequency", "Day"),
                                        resource.TestCheckResourceAttr(dataSourceName, "first_run_time", formattedExpectedTime),
                                ),
                        },
                },
        })
}

func testAccDataSourceAzureRMFirstRunConfig_Daily_Today(now time.Time) string {
	scheduletime := now.Add(time.Duration(1) * time.Hour)

	return fmt.Sprintf(`
data "azurerm_first_run_config" "test" {
	"hour" = "%d"		
        "minute" = "0"		
        "second" = "0"
	"frequency" = "Day"
}
`, scheduletime.Hour())
}

func testAccDataSourceAzureRMFirstRunConfig_Daily_Tomorrow(now time.Time) string {
        scheduletime := now.Add(time.Duration(-1) * time.Hour)

        return fmt.Sprintf(`
data "azurerm_first_run_config" "test" {
        "hour" = "%d"
        "minute" = "0"
        "second" = "0"
        "frequency" = "Day"
}
`, scheduletime.Hour())
}
