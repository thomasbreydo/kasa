package addflag

import "github.com/spf13/cobra"

func DeviceIP(p *[]string, command *cobra.Command) {
	command.PersistentFlags().StringSliceVar(p, "ip", []string{}, "device IPs or hostnames")
	_ = command.MarkPersistentFlagRequired("ip")
}

func IgnoreErrors(p *bool, command *cobra.Command) {
	command.PersistentFlags().BoolVar(
		p,
		"ignore-errors",
		false,
		"try all specified device IPs and hostnames, ignoring errors")
}
