package entities

type Cotacao struct {
	USDBRL USDBRL `json:"USDBRL,omitempty"`
}

type USDBRL struct {
	Code       string `json:"code,omitempty"`
	CodeIn     string `json:"codein,omitempty"`
	Name       string `json:"name,omitempty"`
	High       string `json:"high,omitempty"`
	Low        string `json:"low,omitempty"`
	VarBid     string `json:"varBid,omitempty"`
	PctChange  string `json:"pctChange,omitempty"`
	Bid        string `json:"bid,omitempty"`
	Ask        string `json:"ask,omitempty"`
	Timestamp  string `json:"timestamp,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}
