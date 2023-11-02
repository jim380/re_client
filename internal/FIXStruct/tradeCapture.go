package fixstruct

type TradeCapture struct {
	SessionID          string `json:"sessionID"`
	TradeCaptureReport struct {
		Header               Header  `json:"header"`
		TradeReportID        string  `json:"tradeReportID"`
		TradeReportTransType string  `json:"tradeReportTransType"`
		TradeReportType      string  `json:"tradeReportType"`
		TrdType              string  `json:"trdType"`
		TrdSubType           string  `json:"trdSubType"`
		Side                 string  `json:"side"`
		OrderQty             string  `json:"orderQty"`
		LastQty              string  `json:"lastQty"`
		LastPx               string  `json:"lastPx"`
		GrossTradeAmt        string  `json:"grossTradeAmt"`
		ExecID               string  `json:"execID"`
		OrderID              string  `json:"orderID"`
		TradeID              string  `json:"tradeID"`
		OrigTradeID          string  `json:"origTradeID"`
		Symbol               string  `json:"symbol"`
		SecurityID           string  `json:"securityID"`
		SecurityIDSource     string  `json:"securityIDSource"`
		TradeDate            string  `json:"tradeDate"`
		TransactTime         string  `json:"transactTime"`
		SettlType            string  `json:"settlType"`
		SettlDate            string  `json:"settlDate"`
		Trailer              Trailer `json:"trailer"`
	} `json:"tradeCaptureReport"`
	TradeCaptureReportAcknowledgement struct {
		Header                  Header  `json:"header"`
		TradeReportID           string  `json:"tradeReportID"`
		TradeID                 string  `json:"tradeID"`
		SecondaryTradeID        string  `json:"secondaryTradeID"`
		TradeReportType         string  `json:"tradeReportType"`
		TrdType                 string  `json:"trdType"`
		TrdSubType              string  `json:"trdSubType"`
		ExecType                string  `json:"execType"`
		TradeReportRefID        string  `json:"tradeReportRefID"`
		SecondaryTradeReportID  string  `json:"secondaryTradeReportID"`
		TradeReportStatus       string  `json:"tradeReportStatus"`
		TradeTransType          string  `json:"tradeTransType"`
		TradeReportRejectReason string  `json:"tradeReportRejectReason"`
		Text                    string  `json:"text"`
		Trailer                 Trailer `json:"trailer"`
	} `json:"tradeCaptureReportAcknowledgement"`
	TradeCaptureReportRejection interface{} `json:"tradeCaptureReportRejection"`
}

type Pagination struct {
	NextKey string `json:"next_key"`
	Total   string `json:"total"`
}

type TradeCaptureResponse struct {
	TradeCapture []TradeCapture `json:"TradeCapture"`
	Pagination   Pagination     `json:"pagination"`
}
