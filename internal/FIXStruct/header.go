package fixstruct

type Header struct {
	BeginString  string `json:"beginString"`
	BodyLength   string `json:"bodyLength"`
	MsgType      string `json:"msgType"`
	SenderCompID string `json:"senderCompID"`
	TargetCompID string `json:"targetCompID"`
	MsgSeqNum    string `json:"msgSeqNum"`
	SendingTime  string `json:"sendingTime"`
	ChainID      string `json:"chainID"`
}
