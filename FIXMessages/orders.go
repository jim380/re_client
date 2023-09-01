package fixmessages

import (
	"fmt"
	"os"
	"path/filepath"

	queries "github.com/jim380/re_client/internal/Queries"
	"github.com/spf13/cobra"
)

var outputFile string

var CmdOrders = &cobra.Command{
	Use:   "orders [address]",
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
		if outputFile != "" {
			dir := filepath.Dir(outputFile)
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				os.MkdirAll(dir, 0755)
			}
		}
		if outputFile != "" {
			err := os.WriteFile(outputFile, []byte(outputText), 0644) // <-- Changed this line
			if err != nil {
				fmt.Println("Failed to write to file:", err)
				return
			}
			fmt.Printf("Output saved to %s\n", outputFile)
		} else {
			fmt.Print(outputText)
		}
	},
}

func init() {
	CmdOrders.Flags().StringVarP(&outputFile, "output", "o", "", "Output file")
}
