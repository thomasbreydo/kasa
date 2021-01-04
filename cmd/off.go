package cmd

import (
	"github.com/cgreenhalgh/hs1xxplug"
	"github.com/spf13/cobra"
	"github.com/thomasbreydo/kasa/device"
)

var OffCmd = &cobra.Command{
	Use:   "off <ip>",
	Short: "Turn off a device",
	Long:  "Turn off a device using its IP address or hostname.",
	Args:  OnlyIP,
	RunE: func(cmd *cobra.Command, args []string) error {
		plug := hs1xxplug.Hs1xxPlug{IPAddress: args[0]}
		return device.TurnOff(&plug)
	},
}
