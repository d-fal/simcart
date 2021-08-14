package cmd

import (
	"context"
	"fmt"
	"simcart/app/scaffold"
	"simcart/pkg/logger"

	"github.com/spf13/cobra"
)

type CommandLine interface {
	migrate(cmd *cobra.Command, args []string)
	seed(cmd *cobra.Command, args []string)
	version(cmd *cobra.Command, args []string)
	run(cmd *cobra.Command, args []string)
	mode(cmd *cobra.Command, args []string)
	extensions(cmd *cobra.Command, args []string)
	createModels(cmd *cobra.Command, args []string)
}
type command struct{}

var (
	debug   bool
	counter int
	path    string
	logPath string
	Runner  CommandLine = &command{}
	rootCmd             = &cobra.Command{
		Use:              "simcart",
		Short:            "simcart app",
		Run:              Runner.run,
		TraverseChildren: true,
	}
	s scaffold.Scaffold

	systemwideContext context.Context
	cancelFunc        context.CancelFunc
)

func init() {

	systemwideContext, cancelFunc = context.WithCancel(context.Background())

	s = scaffold.Prepare(systemwideContext)

	rootCmd.Flags().StringVarP(&path, "path", "p",
		".",
		"Set base path for the config file. Default is . ")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(seedCmd)
	rootCmd.AddCommand(migrateCMD)
	rootCmd.AddCommand(modeCmd)

	s.SetConfigPath(path)

}

func Execute() error {
	return rootCmd.Execute()
}

func (c *command) run(cmd *cobra.Command, args []string) {

	// set logger
	lg := logger.NewPrototype(logger.WithDebug(debug), logger.WithPath(logPath))

	systemwideContext = context.WithValue(systemwideContext, "logger", lg)

	if err := s.NormalStart(systemwideContext); err != nil {
		fmt.Println("error: ", err)
	}
}
