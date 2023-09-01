package fixstruct

type ExecutionReport struct {
	SessionID    string  `json:"sessionID"`
	Header       Header  `json:"header"`
	ClOrdID      string  `json:"clOrdID"`
	OrderID      string  `json:"orderID"`
	ExecID       string  `json:"execID"`
	OrdStatus    string  `json:"ordStatus"`
	ExecType     string  `json:"execType"`
	Symbol       string  `json:"symbol"`
	Side         string  `json:"side"`
	OrderQty     string  `json:"orderQty"`
	Price        string  `json:"price"`
	TimeInForce  string  `json:"timeInForce"`
	LastPx       string  `json:"lastPx"`
	LastQty      string  `json:"lastQty"`
	LeavesQty    string  `json:"leavesQty"`
	CumQty       string  `json:"cumQty"`
	AvgPx        string  `json:"avgPx"`
	Text         string  `json:"text"`
	TransactTime string  `json:"transactTime"`
	Trailer      Trailer `json:"trailer"`
}

type OrdersExecutionReportResponse struct {
	OrdersExecutionReport []ExecutionReport `json:"OrdersExecutionReport"`
}
