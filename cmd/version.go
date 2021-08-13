package cmd

import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Generate version",
	Run:   Runner.seed,
}

func (c *command) version(cmd *cobra.Command, args []string) {

}
