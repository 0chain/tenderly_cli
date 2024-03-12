package config

import (
	"github.com/spf13/cobra"
)

// commandRoot represents root command used for command handlers registration.
var commandRoot = &cobra.Command{
	Use:   "tenderly_cli",
	Short: "Use TenderlyCLI to interact with Tenderly API",
}

func GetCommandRoot() *cobra.Command {
	return commandRoot
}
