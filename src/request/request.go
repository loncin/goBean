package request

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"httpReq"
	"security"
	"strconv"
	"time"
	"request/req"
)

var serviceCodeText = map[int]string{
	req.Search:   "Search",
	req.Verify:   "Verify",
	req.Order:    "Order",
	req.PayCheck: "PayCheck",
	req.Ticket:   "Ticket",
}

type jsonReq struct {
	ServiceCode  string `json:"serviceCode"`
	Pid          string `json:"pid"`
	Sign         string `json:"sign"`
	ReqeustID    string `json:"requestID"`
	TimeStamp    string `json:"timestamp"`
	BusinessType string `json:"businessType"`
	Params       string `json:"params"`
}

// api请求
func Request(reqParam req.ReqParam) (string, error) {
	var out string
	paramByte, err := json.Marshal(reqParam.Params)
	if err != nil {
		return out, err
	}

	paramStr := string(paramByte)

	serviceCode := serviceCodeText[reqParam.ServiceCode]
	param := jsonReq{
		ServiceCode:  serviceCode,
		Pid:          reqParam.User.Pid,
		BusinessType: strconv.Itoa(reqParam.BusinessType),
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	param.Sign = getSign(reqParam.User.Pid, reqParam.User.SecurityCode, timestamp, paramStr)
	param.ReqeustID = security.UniqueId()
	param.TimeStamp = timestamp

	// 生单参数先AES加密再base64
	if reqParam.ServiceCode == req.Order {
		securityByte, err := security.AesEncrypt([]byte(paramByte), []byte(reqParam.User.SecurityCode))
		if err != nil {
			return out, err
		}

		base64Str := base64.StdEncoding.EncodeToString(securityByte)
		paramByte = []byte(base64Str)
	}

	gzipByte, err := httpReq.GzipEncode(paramByte)
	if err != nil {
		return out, err
	}

	param.Params = base64.StdEncoding.EncodeToString(gzipByte)
	jsonParam, err := json.Marshal(param)
	if err != nil {
		return out, err
	}

	header := make(map[string]string)
	header["Content-Type"] = "text/plain; charset=UTF8"

	resp, err := httpReq.RequestUrl(httpReq.POST, reqParam.Url, header, string(jsonParam))
	if err != nil {
		return resp, err
	}

	fmt.Print(resp)

	// 出参先base64解码再解压缩
	base64Decode, err := base64.StdEncoding.DecodeString(resp)
	if err != nil {
		return out, err
	}

	unGzip, err := httpReq.GzipDecode(base64Decode)
	if err != nil {
		return out, err
	}

	return string(unGzip), nil
}

// 获取sign
func getSign(pid string, securityCode string, timestamp string, params string) string {
	var buffer bytes.Buffer
	buffer.WriteString(params)
	buffer.WriteString(timestamp)
	buffer.WriteString(securityCode)
	buffer.WriteString(pid)

	return security.GetMD5Hash(buffer.String())
}

// 获取下单加密数据
func getOrdeSecurityData(securityCode string, params string) (string, error) {
	encrypt, err := security.AesEncrypt([]byte(params), []byte(securityCode))

	var output string

	if err != nil {
		return output, err
	}

	base64str := base64.StdEncoding.EncodeToString(encrypt)

	return base64str, nil
}
