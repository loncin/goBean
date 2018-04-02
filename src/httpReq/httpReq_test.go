package httpReq

import (
	"testing"
)

func TestRequestUrl(t *testing.T) {
	method := GET
	url := "http://www.baidu.com"
	header := map[string]string {"Content-Type" : "text/plain; charset=UTF8"}
	body := ""

	resp, err := RequestUrl(method, url, header, body)
	if err != nil {
		t.Errorf("RequestUrl error: %s", err.Error())
	}
	
	t.Errorf("RequestUrl response:%s", resp)
}