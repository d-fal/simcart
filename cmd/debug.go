package cmd

import (
	"simcart/config"

	"github.com/spf13/cobra"
)

var (
	modeCmd = &cobra.Command{
		Use:   "mode",
		Short: "debug mode [true|false] default is false",
		Run:   Runner.mode,
	}
)

func init() {
	modeCmd.Flags().BoolVarP(&debug, "debug", "d", false, "debug mode")
	modeCmd.Flags().IntVarP(&counter, "level", "l", 0, "Set log level")
	modeCmd.Flags().StringVarP(&logPath, "logpath", "p", ".", "set log path")
}

func (c *command) mode(cmd *cobra.Command, args []string) {

	config.GetAppConfig().Debug(debug)
	c.run(cmd, args)

}
