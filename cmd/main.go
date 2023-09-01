package main

import (
	fixmessages "github.com/jim380/re_client/FIXMessages"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "Re_client"}

	rootCmd.AddCommand(fixmessages.CmdOrders)
	rootCmd.Execute()
}
