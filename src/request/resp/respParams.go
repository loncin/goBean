package resp

import (
	"request/model"
)

// 响应基类
type BaseResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// 查询响应
type SearchResp struct {
	BaseResp
	Routings []model.Routing `json:"routings"`
}

// 验价响应
type VerifyResp struct {
	BaseResp
	SessionId string        `json:"sessionId"`
	MaxSeats  int           `json:"maxSeats"`
	Routing   model.Routing `json:"routing"`
	Rule      model.Rule    `json:"rule"`
}

// 下单响应
type OrderResp struct {
	BaseResp
	SessionId string        `json:"sessionId"`
	OrderNo   string        `json:"orderNo"`
	PnrCode   string        `json:"pnrCode"`
	Routing   model.Routing `json:"routing"`
}

// 支付校验响应
type PayCheckResp struct {
	BaseResp
	SessionId string        `json:"sessionId"`
	OrderNo   string        `json:"orderNo"`
	PnrCode   string        `json:"pnrCode"`
	Routing   model.Routing `json:"routing"`
}
