package stress

import (
	"io/ioutil"
	"encoding/json"
	"request/req"
)

type Stress struct {
	StressRound       int `json:"stressRound"`
	ThreadCount       int `json:"threadCount"`
	ReqCountPerThread int `json:"reqCountPerThread"`
}

type Configuration struct {
	Url    string       `json:"url"`
	User   req.UserInfo `json:"user"`
	Stress Stress       `json:"stress"`
}

type SearchParams struct {
	Params []req.SearchReqParam `json:"params"`
}

const confFile string = "conf.json"
const searchFile string = "search.json"

func loadConf() (conf Configuration, err error) {
	conf = Configuration{}

	err = loadDataFromFile(confFile, &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

func loadSearchParam() (searchParams SearchParams, err error) {
	searchParams = SearchParams{}

	err = loadDataFromFile(searchFile, &searchParams)
	if err != nil {
		return searchParams, err
	}

	return searchParams, nil
}

func loadDataFromFile(filePath string, v interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}