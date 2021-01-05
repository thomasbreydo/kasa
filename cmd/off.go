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
	Short:        "Turn off devices",
	Long:         "Turn off devices given their IP addresses or hostname.",
	Args:         cobra.NoArgs,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		resetStreams := quiet.SetupAndGetTeardown()
		defer resetStreams()
		for _, deviceIP := range DeviceIPs {
			plug := hs1xxplug.Hs1xxPlug{IPAddress: deviceIP}
			err := smartdevice.TurnOff(&plug)
			if !IgnoreErrors && err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	addflag.DeviceIP(&DeviceIPs, OffCmd)
	addflag.IgnoreErrors(&IgnoreErrors, OffCmd)
}
