package cmd

import (
	"github.com/cgreenhalgh/hs1xxplug"
	"github.com/spf13/cobra"
	"github.com/thomasbreydo/kasa/device"
)

var OnCmd = &cobra.Command{
	Use:   "on <ip>",
	Short: "Turn on a device",
	Long:  "Turn on a device using its IP address or hostname.",
	Args:  OnlyIP,
	RunE: func(cmd *cobra.Command, args []string) error {
		plug := hs1xxplug.Hs1xxPlug{IPAddress: args[0]}
		return device.TurnOn(&plug)
	},
}
