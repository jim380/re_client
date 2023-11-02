package tradecapture

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	queries "github.com/jim380/re_client/internal/Queries"
	"github.com/jim380/re_client/utils"
	"github.com/spf13/cobra"
)

var CmdTradeCaptureAll = &cobra.Command{
	Use:   "trade-captures [chainID]",
	Short: "Fetch and convert trade captures to FIX format",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		argChainID := args[0]

		tradeCaptures, err := queries.FetchAllTradeCapture(argChainID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		var outputText string

		for _, tradeCapture := range tradeCaptures.TradeCapture {
			tradeCaptureMessage := GenerateTradeCaptureMessage(tradeCapture)
			outputText += " " + tradeCaptureMessage + "\n\n"
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
	CmdTradeCaptureAll.Flags().StringVarP(&utils.OutputFile, "output", "o", "", "Output file")
}
