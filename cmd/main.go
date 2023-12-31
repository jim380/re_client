package main

import (
	"log"

	executionreports "github.com/jim380/re_client/FIXMessages/executionReports"
	"github.com/jim380/re_client/FIXMessages/orders"
	tradecapture "github.com/jim380/re_client/FIXMessages/tradeCapture"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "Re_client"}

	// orders
	rootCmd.AddCommand(orders.CmdOrders)
	rootCmd.AddCommand(orders.CmdOrdersAll)

	// execution reports
	rootCmd.AddCommand(executionreports.CmdExecutionReports)
	rootCmd.AddCommand(executionreports.CmdExecutionReportsAll)

	// trade capture
	rootCmd.AddCommand(tradecapture.CmdTradeCapture)
	rootCmd.AddCommand(tradecapture.CmdTradeCaptureAll)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %s", err)
	}
}
