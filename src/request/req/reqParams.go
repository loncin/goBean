package req

import (
	"request/model"
)

// serviceCode
const (
	Search = iota
	Verify
	Order
	PayCheck
	Ticket
)

// businessType
const (
	Demostic = iota
	International
)

type UserInfo struct {
	Pid          string `json:"pid"`
	SecurityCode string `json:"securityCode"`
}

// 请求参数
type ReqParam struct {
	ServiceCode  int
	User         UserInfo
	Url          string
	BusinessType int // 0 国内 1国际
	RequestId    string
	Params       interface{}
}

// 查询参数
type SearchReqParam struct {
	ProductType int    `json:"productType"` // 产品类型 1优选 2特惠
	TripType    string `json:"tripType"`    // 行程类型 1单程 2往返
	FromCity    string `json:"fromCity"`    // 出发城市IATA三字码
	ToCity      string `json:"toCity"`      //到达城市IATA三字码
	FromDate    string `json:"fromDate"`    // 出发日期 yyyyMMdd
	RetDate     string `json:"retDate"`     // 到达日期 yyyyMMdd
	AdultNumber int    `json:"adultNumber"` // 成人数量
	ChildNumber int    `json:"childNumber"` // 儿童数量
	Carrier     string `json:"carrier"`     // 航司二字码
	CabinGrade  string `json:"cabinGrade"`  // 舱等 F头等舱 C商务舱 S超级经济舱 Y经济舱
}

// 验价参数
type VerifyReqParam struct {
	TripType    string        `json:"tripType"`
	AdultNumber int           `json:"adultNumber"`
	ChildNumber int           `json:"childNumber"`
	Routing     model.Routing `json:"routing"`
}

// 下单参数
type OrderReqParam struct {
	TripType        string            `json:"tripType"`
	SessionId       int               `json:"sessionId"`
	ExternalOrderNo int               `json:"externalOrderNo"`
	Routing         model.Routing     `json:"routing"`
	Passengers      []model.Passenger `json:"passengers"`
	Contract        model.Contract    `json:"contract"`
}

// 支付校验参数
type PayCheckReqParam struct {
	TripType        string        `json:"tripType"`
	SessionId       int           `json:"sessionId"`
	OrderNo         int           `json:"orderNo"`
	ExternalOrderNo int           `json:"externalOrderNo"`
	PnrCode         int           `json:"pnrCode"`
	Routing         model.Routing `json:"routing"`
}
