package cmd

import (
	"github.com/cgreenhalgh/hs1xxplug"
	"github.com/spf13/cobra"
	"github.com/thomasbreydo/kasa/addflag"
	"github.com/thomasbreydo/kasa/quiet"
	"github.com/thomasbreydo/kasa/smartdevice"
)

var OnCmd = &cobra.Command{
	Use:          "on",
	Short:        "Turn on a device",
	Long:         "Turn on a device given its IP address or hostname.",
	Args:         cobra.NoArgs,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		teardownFunc := quiet.SetupAndGetTeardown()
		defer teardownFunc()
		plug := hs1xxplug.Hs1xxPlug{IPAddress: DeviceIP}
		return smartdevice.TurnOn(&plug)
	},
}

func init() {
	addflag.IP(&DeviceIP, OnCmd)
}
