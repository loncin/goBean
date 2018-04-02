package httpReq

import (
	"io/ioutil"
	"fmt"
	"errors"
	"net/http"
	"strings"
)

const (
	GET = iota
	POST
)

var reqeustMethodText = map[int]string{
	GET:  "GET",
	POST: "POST",
}

func RequestUrl(method int, url string, header map[string]string, body string) (string, error) {
	fmt.Printf("request body:%s\n", body)

	bodyReader := strings.NewReader(body)
	req, err := http.NewRequest(reqeustMethodText[method], url, bodyReader)

	var out string
	if err != nil {
		return out, err
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return out, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := errors.New("request status:" + resp.Status)
		return out, err
	}

	contentByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return out, err
	}

	respStr := string(contentByte)

	fmt.Printf("response:%s\n", respStr)

	return respStr, nil
}
