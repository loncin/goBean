package request

import (
	"encoding/json"
	"request/resp"
	"request/req"
	"config"
	"security"
	"testing"
)

func TestRequest(t *testing.T) {
	reqParam := req.ReqParam{
		ServiceCode:  req.Search,
		Url:          "http://tcflightopenapi.qa.17usoft.com/flightdistributeapi/dispatcher/api",
		BusinessType: 1,
	}

	reqParam.RequestId = security.UniqueId()
	reqParam.User = config.UserInfo{
		Pid:          "dbceea59d2024cf7bfd22220e04f9ca2",
		SecurityCode: "ff84853857884aa1",
	}

	reqParam.Params = req.SearchReqParam{
		ProductType: 2,
		TripType:    "1",
		FromCity:    "SHA",
		ToCity:      "HKG",
		FromDate:    "20180409",
		RetDate:     "",
		AdultNumber: 1,
		ChildNumber: 1,
	}

	respStr, err := Request(reqParam)
	if err != nil {
		t.Errorf("TestRequest request error:%s", err.Error())
	}

	t.Errorf("TestRequest response:%s", respStr)

	var searchResp resp.SearchResp
	err = json.Unmarshal([]byte(respStr), &searchResp)
	if err != nil {
		t.Errorf("TestRequest request error:%s", err.Error())
	}

	t.Errorf("TestRequest response msg:%s", searchResp.Msg)
}
