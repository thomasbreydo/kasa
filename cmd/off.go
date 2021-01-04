package cmd

import (
	"github.com/cgreenhalgh/hs1xxplug"
	"github.com/spf13/cobra"
	"github.com/thomasbreydo/kasa/addflag"
	"github.com/thomasbreydo/kasa/quiet"
	"github.com/thomasbreydo/kasa/smartdevice"
)

var OffCmd = &cobra.Command{
	Use:          "off",
	Short:        "Turn off a device",
	Long:         "Turn off a device given its IP address or hostname.",
	Args:         cobra.NoArgs,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		teardownFunc := quiet.SetupAndGetTeardown()
		defer teardownFunc()
		plug := hs1xxplug.Hs1xxPlug{IPAddress: DeviceIP}
		return smartdevice.TurnOff(&plug)
	},
}

func init() {
	addflag.IP(&DeviceIP, OffCmd)
}
