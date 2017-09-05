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
	scheduletime := now.Add(time.Duration(1) * time.Hour)

	config := testAccDataSourceAzureRMFirstRunConfig_Daily(scheduletime)

	expectedTime := time.Date(scheduletime.Year(), scheduletime.Month(), scheduletime.Day(), scheduletime.Hour(), 0, 0, 0, time.UTC)
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
	scheduletime := now.Add(time.Duration(-1) * time.Hour)
        config := testAccDataSourceAzureRMFirstRunConfig_Daily(scheduletime)

	expectedTime := time.Date(scheduletime.Year(), scheduletime.Month(), scheduletime.Day()+1, scheduletime.Hour(), 0, 0, 0, time.UTC)
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

func TestAccDataSourceAzureRMFirstRunConfig_Hourly_CurrentHour(t *testing.T) {
        dataSourceName := "data.azurerm_first_run_config.test"

        now := time.Now().UTC()
	scheduletime := now.Add(time.Duration(1) * time.Minute)

        config := testAccDataSourceAzureRMFirstRunConfig_Hourly(scheduletime)

        expectedTime := time.Date(scheduletime.Year(), scheduletime.Month(), scheduletime.Day(), scheduletime.Hour(), scheduletime.Minute(), 0, 0, time.UTC)
        formattedExpectedTime := expectedTime.Format(time.RFC3339)

        resource.Test(t, resource.TestCase{
                PreCheck:  func() { testAccPreCheck(t) },
                Providers: testAccProviders,
                Steps: []resource.TestStep{
                        {
                                Config: config,
                                Check: resource.ComposeTestCheckFunc(
                                        resource.TestCheckResourceAttr(dataSourceName, "hour", strconv.Itoa(scheduletime.Hour())),
                                        resource.TestCheckResourceAttr(dataSourceName, "minute", strconv.Itoa(scheduletime.Minute())),
                                        resource.TestCheckResourceAttr(dataSourceName, "second", "0"),
                                        resource.TestCheckResourceAttr(dataSourceName, "frequency", "Hour"),
                                        resource.TestCheckResourceAttr(dataSourceName, "first_run_time", formattedExpectedTime),
                                ),
                        },
                },
        })
}

func TestAccDataSourceAzureRMFirstRunConfig_Hourly_NextHour(t *testing.T) {
        dataSourceName := "data.azurerm_first_run_config.test"

        now := time.Now().UTC()
        scheduletime := now.Add(time.Duration(-1) * time.Minute)

        config := testAccDataSourceAzureRMFirstRunConfig_Hourly(scheduletime)

        expectedTime := time.Date(scheduletime.Year(), scheduletime.Month(), scheduletime.Day(), scheduletime.Hour()+1, scheduletime.Minute(), 0, 0, time.UTC)
        formattedExpectedTime := expectedTime.Format(time.RFC3339)

        resource.Test(t, resource.TestCase{
                PreCheck:  func() { testAccPreCheck(t) },
                Providers: testAccProviders,
                Steps: []resource.TestStep{
                        {
                                Config: config,
                                Check: resource.ComposeTestCheckFunc(
                                        resource.TestCheckResourceAttr(dataSourceName, "hour", strconv.Itoa(scheduletime.Hour())),
                                        resource.TestCheckResourceAttr(dataSourceName, "minute", strconv.Itoa(scheduletime.Minute())),
                                        resource.TestCheckResourceAttr(dataSourceName, "second", "0"),
                                        resource.TestCheckResourceAttr(dataSourceName, "frequency", "Hour"),
                                        resource.TestCheckResourceAttr(dataSourceName, "first_run_time", formattedExpectedTime),
                                ),
                        },
                },
        })
}

func testAccDataSourceAzureRMFirstRunConfig_Daily(scheduletime time.Time) string {
	return fmt.Sprintf(`
data "azurerm_first_run_config" "test" {
	"hour" = "%d"		
        "minute" = "0"		
        "second" = "0"
	"frequency" = "Day"
}
`, scheduletime.Hour())
}

func testAccDataSourceAzureRMFirstRunConfig_Hourly(scheduletime time.Time) string {
        return fmt.Sprintf(`
data "azurerm_first_run_config" "test" {
        "hour" = "%d"
        "minute" = "%d"
        "second" = "0"
        "frequency" = "Hour"
}
`, scheduletime.Hour(), scheduletime.Minute())
}
