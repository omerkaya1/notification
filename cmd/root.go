package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd only starts the app according to selected child command
var RootCmd = &cobra.Command{
	Use:   "notification",
	Short: "simple notification service",
}

func init() {
	RootCmd.AddCommand()
}
