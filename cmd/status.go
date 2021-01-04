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
	Short:        "Show the status (on/off) of a device",
	Long:         "Show the status (on/off) of a device given its IP or hostname.",
	Args:         cobra.NoArgs,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		device := hs1xxplug.Hs1xxPlug{IPAddress: DeviceIP}
		resetStreams := quiet.SetupAndGetTeardown()
		isOff, err := smartdevice.IsOff(&device)
		resetStreams()
		if err != nil {
			return err
		}
		switch isOff {
		case true:
			fmt.Println("The device is off.")
		case false:
			fmt.Println("The device is on.")
		}
		return nil
	},
}

func init() {
	addflag.IP(&DeviceIP, StatusCmd)
}
