package commands

import "github.com/spf13/cobra"

var (
	cmd = &cobra.Command{
		Use:   "config-play",
		Short: "config play",
		Long:  "config playground",
	}
)

func init() {

}

func NewConfigPlayCmd() *cobra.Command {
	return cmd
}
