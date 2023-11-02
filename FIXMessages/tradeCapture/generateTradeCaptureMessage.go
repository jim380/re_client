package tradecapture

import (
	"fmt"

	fixstruct "github.com/jim380/re_client/internal/FIXStruct"
)

func GenerateTradeCaptureMessage(tradeCapture fixstruct.TradeCapture) string {
	tradeCaptureReport := tradeCapture.TradeCaptureReport
	tradeCaptureAcknowledgement := tradeCapture.TradeCaptureReportAcknowledgement
	var tradeCaptureRejection string
	if tradeCapture.TradeCaptureReportRejection != nil {
		rejection := *tradeCapture.TradeCaptureReportRejection

		tradeCaptureRejection = fmt.Sprintf(
			"8=%s|9=%s|35=j|49=%s|56=%s|34=%s|52=%s|11=%s|37=%s|17=%s|150=%s|39=%s|55=%s|54=%s|44=%s|59=%s|32=%s|31=%s|151=%s|14=%s|6=%s|58=%s|60=%s|10=%s",
			rejection.Header.BeginString,
			rejection.Header.BodyLength,
			rejection.Header.SenderCompID,
			rejection.Header.TargetCompID,
			rejection.Header.MsgSeqNum,
			rejection.Header.SendingTime,
			rejection.TradeReportID,
			rejection.TradeReportRejectReason,
			rejection.TradeReportRejectRefID,
			rejection.Text,
			rejection.Trailer.CheckSum,
		)
	}

	message := fmt.Sprintf(
		"8=%s|9=%s|35=AE|49=%s|56=%s|34=%s|52=%s|11=%s|37=%s|17=%s|150=%s|39=%s|55=%s|54=%s|44=%s|59=%s|32=%s|31=%s|151=%s|14=%s|6=%s|58=%s|60=%s|10=%s",
		tradeCaptureReport.Header.BeginString,
		tradeCaptureReport.Header.BodyLength,
		tradeCaptureReport.Header.SenderCompID,
		tradeCaptureReport.Header.TargetCompID,
		tradeCaptureReport.Header.MsgSeqNum,
		tradeCaptureReport.Header.SendingTime,
		tradeCaptureReport.TradeReportID,
		tradeCaptureReport.TradeReportTransType,
		tradeCaptureReport.TradeReportType,
		tradeCaptureReport.TrdType,
		tradeCaptureReport.TrdSubType,
		tradeCaptureReport.Side,
		tradeCaptureReport.OrderQty,
		tradeCaptureReport.LastQty,
		tradeCaptureReport.LastPx,
		tradeCaptureReport.GrossTradeAmt,
		tradeCaptureReport.ExecID,
		tradeCaptureReport.OrderID,
		tradeCaptureReport.TradeID,
		tradeCaptureReport.OrigTradeID,
		tradeCaptureReport.Symbol,
		tradeCaptureReport.SecurityID,
		tradeCaptureReport.SecurityIDSource,
		tradeCaptureReport.TradeDate,
		tradeCaptureReport.TransactTime,
		tradeCaptureReport.SettlType,
		tradeCaptureReport.SettlDate,
		tradeCaptureReport.Trailer.CheckSum,
	)

	message += "\n\n"

	if tradeCaptureAcknowledgement.TradeReportID != "" {
		acknowledgementMessage := fmt.Sprintf(
			"8=%s|9=%s|35=AR|49=%s|56=%s|34=%s|52=%s|11=%s|37=%s|17=%s|150=%s|39=%s|55=%s|54=%s|44=%s|59=%s|32=%s|31=%s|151=%s|14=%s|6=%s|58=%s|60=%s|10=%s",
			tradeCaptureAcknowledgement.Header.BeginString,
			tradeCaptureAcknowledgement.Header.BodyLength,
			tradeCaptureAcknowledgement.Header.SenderCompID,
			tradeCaptureAcknowledgement.Header.TargetCompID,
			tradeCaptureAcknowledgement.Header.MsgSeqNum,
			tradeCaptureAcknowledgement.Header.SendingTime,
			tradeCaptureAcknowledgement.TradeReportID,
			tradeCaptureAcknowledgement.OrderID,
			tradeCaptureAcknowledgement.OrderID,
			tradeCaptureAcknowledgement.TrdType,
			tradeCaptureAcknowledgement.TradeReportType,
			tradeCaptureAcknowledgement.Symbol,
			tradeCaptureAcknowledgement.Side,
			tradeCaptureAcknowledgement.Price,
			tradeCaptureAcknowledgement.TimeInForce,
			tradeCaptureAcknowledgement.LastQty,
			tradeCaptureAcknowledgement.LastPx,
			tradeCaptureAcknowledgement.LeavesQty,
			tradeCaptureAcknowledgement.CumQty,
			tradeCaptureAcknowledgement.AvgPx,
			tradeCaptureAcknowledgement.Text,
			tradeCaptureAcknowledgement.TransactTime,
			tradeCaptureAcknowledgement.Trailer.CheckSum,
		)

		message += " " + acknowledgementMessage + "\n\n"
	}

	if tradeCaptureRejection != "" {
		message += " " + tradeCaptureRejection + "\n\n"
	}

	return message
}
