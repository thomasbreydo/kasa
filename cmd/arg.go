package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func OnlyIP(cmd *cobra.Command, args []string) error {
	switch len(args) {
	case 0:
		return errors.New(`"ip" must be specified`)
	case 1:
		return nil
	default:
		return errors.New(`only "ip" can be specified"`)
	}
}
