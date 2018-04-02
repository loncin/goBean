package model

// 退改规则
type Rule struct {
	BaggageInfoList []Baggage `json:"baggageInfoList"`
	ChangesInfoList []Changes `json:"changesInfoList"`
	RefundInfoList  []Refund  `json:"refundInfoList"`
	Note            string    `json:"note"`
}

// 行李规则
type Baggage struct {
	AdultBaggage string `json:"adultBaggage"`
	ChildBaggage string `json:"childBaggage"`
	SegmentNo    string `json:"segmentNo"`
}

// 改签规则
type Changes struct {
	ChangesType   string `json:"changesType"`
	ChangesStatus string `json:"changesStatus"`
	ChangesFee    string `json:"changesFee"`
	Currency      string `json:"currency"`
	CnRemark      string `json:"cnRemark"`
	EnRemark      string `json:"enRemark"`
}

// 退票规则
type Refund struct {
	RefundType   string `json:"refundType"`
	RefundStatus string `json:"refundStatus"`
	RefundFee    string `json:"refundFee"`
	Currency     string `json:"currency"`
	CnRemark     string `json:"cnRemark"`
	EnRemark     string `json:"enRemark"`
}
