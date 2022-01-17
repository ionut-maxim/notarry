package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(webhookCmd)
}

var (
	webhookCmd = &cobra.Command{
		Use:   "webhook",
		Short: "Start a webhook",
		Long:  "sasas",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Test")
		},
	}
)
