package main

import (
	"github.com/jim380/re_client/FIXMessages/orders"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "Re_client"}

	rootCmd.AddCommand(orders.CmdOrders)
	rootCmd.AddCommand(orders.CmdOrdersAll)
	rootCmd.Execute()
}
