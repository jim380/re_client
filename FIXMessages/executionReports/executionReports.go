package executionreports

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	queries "github.com/jim380/re_client/internal/Queries"
	"github.com/jim380/re_client/utils"
	"github.com/spf13/cobra"
)

var CmdExecutionReports = &cobra.Command{
	Use:   "execution-report [chainID] [address]",
	Short: "Fetch and convert execution reports to FIX format",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		argChainID := args[0]
		argAddress := args[1]

		executionReports, err := queries.FetchExecutionReports(argChainID, argAddress)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		var outputText string

		for _, executionReport := range executionReports.OrdersExecutionReport {
			executionReportMessage := fmt.Sprintf(
				"8=%s|9=%s|35=%s|49=%s|56=%s|34=%s|52=%s|11=%s|37=%s|17=%s|150=%s|39=%s|55=%s|54=%s|44=%s|59=%s|32=%s|31=%s|151=%s|14=%s|6=%s|58=%s|60=%s|10=%s|",
				executionReport.Header.BeginString,
				executionReport.Header.BodyLength,
				executionReport.Header.MsgType,
				executionReport.Header.SenderCompID,
				executionReport.Header.TargetCompID,
				executionReport.Header.MsgSeqNum,
				executionReport.Header.SendingTime,
				executionReport.ClOrdID,
				executionReport.OrderID,
				executionReport.ExecID,
				executionReport.ExecType,
				executionReport.OrdStatus,
				executionReport.Symbol,
				executionReport.Side,
				executionReport.Price,
				executionReport.TimeInForce,
				executionReport.LastQty,
				executionReport.LastPx,
				executionReport.LeavesQty,
				executionReport.CumQty,
				executionReport.AvgPx,
				executionReport.Text,
				executionReport.TransactTime,
				executionReport.Trailer.CheckSum)

			outputText += " " + executionReportMessage + "\n\n"
		}

		// Ensure the folder exists
		if utils.OutputFile != "" {
			dir := filepath.Dir(utils.OutputFile)
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				if err := os.MkdirAll(dir, 0o755); err != nil {
					log.Fatalf("Error creating directory %s: %s", dir, err)
				}
			}
		}
		if utils.OutputFile != "" {
			err := os.WriteFile(utils.OutputFile, []byte(outputText), 0o644)
			if err != nil {
				fmt.Println("Failed to write to file:", err)
				return
			}
			fmt.Printf("Output saved to %s\n", utils.OutputFile)
		} else {
			fmt.Print(outputText)
		}
	},
}

func init() {
	CmdExecutionReports.Flags().StringVarP(&utils.OutputFile, "output", "o", "", "Output file")
}
