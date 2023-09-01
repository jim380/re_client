package orders

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	queries "github.com/jim380/re_client/internal/Queries"
	"github.com/jim380/re_client/utils"
	"github.com/spf13/cobra"
)

var CmdOrders = &cobra.Command{
	Use:   "order [address]",
	Short: "Fetch and convert orders to FIX format",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		address := args[0]

		orders, err := queries.FetchOrders(address)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		var outputText string

		for _, order := range orders.Orders {
			orderMessage := fmt.Sprintf("8=%s|9=%s|35=%s|49=%s|56=%s|34=%s|52=%s|11=%s|55=%s|54=%s|44=%s|59=%s|",
				order.Header.BeginString,
				order.Header.BodyLength,
				order.Header.MsgType,
				order.Header.SenderCompID,
				order.Header.TargetCompID,
				order.Header.MsgSeqNum,
				order.Header.SendingTime,
				order.ClOrdID,
				order.Symbol,
				order.Side,
				order.Price,
				order.TimeInForce)

			outputText += " " + orderMessage + "\n\n"
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
	CmdOrders.Flags().StringVarP(&utils.OutputFile, "output", "o", "", "Output file")
}
