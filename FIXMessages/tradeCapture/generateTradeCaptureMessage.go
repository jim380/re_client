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
			"8=%s|9=%s|35=%s|49=%s|56=%s|34=%s|52=%s|571=%s|751=%s|754=%s|58=%s|10=%s",
			rejection.Header.BeginString,
			rejection.Header.BodyLength,
			rejection.Header.MsgType,
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
		"8=%s|9=%s|35=%s|49=%s|56=%s|34=%s|52=%s|571=%s|487=%s|856=%s|828=%s|829=%s|54=%s|38=%s|32=%s|31=%s|381=%s|17=%s|37=%s|1003=%s|1126=%s|55=%s|48=%s|22=%s|75=%s|60=%s|63=%s|64=%s|10=%s",
		tradeCaptureReport.Header.BeginString,
		tradeCaptureReport.Header.BodyLength,
		tradeCaptureReport.Header.MsgType,
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
			"8=%s|9=%s|35=%s|49=%s|56=%s|34=%s|52=%s|571=%s|1003=%s|1040=%s|856=%s|828=%s|829=%s|150=%s|572=%s|818=%s|939=%s|487=%s|751=%s|58=%s|10=%s",
			tradeCaptureAcknowledgement.Header.BeginString,
			tradeCaptureAcknowledgement.Header.BodyLength,
			tradeCaptureAcknowledgement.Header.MsgType,
			tradeCaptureAcknowledgement.Header.SenderCompID,
			tradeCaptureAcknowledgement.Header.TargetCompID,
			tradeCaptureAcknowledgement.Header.MsgSeqNum,
			tradeCaptureAcknowledgement.Header.SendingTime,
			tradeCaptureAcknowledgement.TradeReportID,
			tradeCaptureAcknowledgement.TradeID,
			tradeCaptureAcknowledgement.SecondaryTradeID,
			tradeCaptureAcknowledgement.TradeReportType,
			tradeCaptureAcknowledgement.TrdType,
			tradeCaptureAcknowledgement.TrdSubType,
			tradeCaptureAcknowledgement.ExecType,
			tradeCaptureAcknowledgement.TradeReportRefID,
			tradeCaptureAcknowledgement.SecondaryTradeReportID,
			tradeCaptureAcknowledgement.TradeReportStatus,
			tradeCaptureAcknowledgement.TradeTransType,
			tradeCaptureAcknowledgement.TradeReportRejectReason,
			tradeCaptureAcknowledgement.Text,
			tradeCaptureAcknowledgement.Trailer.CheckSum,
		)

		message += " " + acknowledgementMessage + "\n\n"
	}

	if tradeCaptureRejection != "" {
		message += " " + tradeCaptureRejection + "\n\n"
	}

	return message
}
