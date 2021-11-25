package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tharun208/config-loader/pkg/config"
)

var initConfig = config.NewConfig()

// newRootCmd represents the base command when called without any subcommands.
func newRootCmd() *cobra.Command {
	// initConfig := config.NewConfig()
	cmd := &cobra.Command{
		Use:   "config-loader",
		Short: "Config-loader for loading configuration files and retrieval",
		Long:  `Config-loader for loading configuration files and retrieval.`,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			cmd.SilenceUsage = true
			return nil
		},
	}
	cmd.AddCommand(newLoadCommand())
	cmd.AddCommand(newGetKeyCommand())
	return cmd
}

func DefaultRootCmd() *cobra.Command {
	return newRootCmd()
}

func Execute() {
	if err := DefaultRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}

func newLoadCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "load",
		Short: "load configuration file(s)",
		Long:  `load configuration file(s)`,
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, file := range args {
				if err := initConfig.LoadFiles(file); err != nil {
					return err
				}
			}
			return nil
		},
	}
	return cmd
}

func newGetKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get value from configuration file",
		Long:  `get value from configuration file based key`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			value := initConfig.GetConfigValue(args[0])
			fmt.Fprintf(cmd.OutOrStdout(), "%v\n", value)
			return nil
		},
	}
	return cmd
}
