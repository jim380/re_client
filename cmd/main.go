package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "Re_client"}

	rootCmd.AddCommand(cmdOrders)
	rootCmd.Execute()
}
