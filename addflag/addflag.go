package addflag

import "github.com/spf13/cobra"

func IP(p *string, command *cobra.Command) {
	command.PersistentFlags().StringVar(p, "ip", "", "smartdevice IP or hostname")
	_ = command.MarkPersistentFlagRequired("ip")
}
