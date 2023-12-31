package fixstruct

type Order struct {
	SessionID   string `json:"sessionID"`
	Header      Header `json:"header"`
	ClOrdID     string `json:"clOrdID"`
	Symbol      string `json:"symbol"`
	Side        string `json:"side"`
	OrderQty    string `json:"orderQty"`
	OrdType     string `json:"ordType"`
	Price       string `json:"price"`
	TimeInForce string `json:"timeInForce"`
	Text        string `json:"text"`
}

type OrderResponse struct {
	Orders []Order `json:"Orders"`
}
