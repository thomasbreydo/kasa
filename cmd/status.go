package cmd

import (
	"fmt"

	"github.com/cgreenhalgh/hs1xxplug"
	"github.com/thomasbreydo/kasa/device"

	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status <ip>",
	Short: "Show the status (on/off) of a device",
	Long:  "Show the status (on/off) of a device using its IP or hostname.",
	Args:  OnlyIP,
	RunE: func(cmd *cobra.Command, args []string) error {
		plug := hs1xxplug.Hs1xxPlug{IPAddress: args[0]}
		isOff, err := device.IsOff(&plug)
		switch isOff {
		case true:
			fmt.Println("The device is off.")
		case false:
			fmt.Println("The device is on.")
		}
		return err
	},
}
