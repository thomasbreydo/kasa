package cmd

import (
	"fmt"

	"github.com/cgreenhalgh/hs1xxplug"
	"github.com/spf13/cobra"
	"github.com/thomasbreydo/kasa/addflag"
	"github.com/thomasbreydo/kasa/quiet"
	"github.com/thomasbreydo/kasa/smartdevice"
)

var StatusCmd = &cobra.Command{
	Use:          "status",
	Short:        "Show the status (on/off) of devices",
	Long:         "Show the status (on/off) of devices given their IP addresses or hostnames.",
	Args:         cobra.NoArgs,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, deviceIP := range DeviceIPs {
			device := hs1xxplug.Hs1xxPlug{IPAddress: deviceIP}
			resetStreams := quiet.SetupAndGetTeardown()
			isOff, err := smartdevice.IsOff(&device)
			resetStreams()
			if err == nil {
				switch isOff {
				case true:
					fmt.Printf("The device at %s is off.\n", deviceIP)
				case false:
					fmt.Printf("The device at %s is on.\n", deviceIP)
				}
			} else if !IgnoreErrors {
				return err
			}
		}
		return nil
	},
}

func init() {
	addflag.DeviceIP(&DeviceIPs, StatusCmd)
	addflag.IgnoreErrors(&IgnoreErrors, StatusCmd)
}
