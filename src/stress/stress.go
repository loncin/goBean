package stress

import (
	"fmt"
	"math/rand"
	"request"
	"request/req"
	"security"
)

func StressGo() error {
	conf, err := loadConf()
	if err != nil {
		return err
	}

	searchParams, err := loadSearchParam()
	if err != nil {
		return err
	}

	// 循环轮次
	for i := 0; i < conf.Stress.StressRound; i++ {
		fmt.Printf("stress test round:%d begin", i)

		threadCount := conf.Stress.ThreadCount
		chs := make([]chan int, threadCount)

		// 协程数
		for j := 0; j < threadCount; j++ {
			chs[j] = make(chan int)
			go requestSearch(j, &conf, &searchParams, chs[j])
		}

		for _, ch := range(chs) {
			fmt.Printf("goroutine:%d", <-ch)
		}

		fmt.Printf("stress test round:%d end", i)
	}

	return nil
}

func requestSearch(threadIndex int, conf *Configuration, searchParams *SearchParams, ch chan int) {
	fmt.Printf("stress test thread:%d begin", threadIndex)

	for i := 0; i < conf.Stress.ReqCountPerThread; i++ {
		reqParam := req.ReqParam{
			ServiceCode:  req.Search,
			User:         conf.User,
			Url:          conf.Url,
			BusinessType: 1,
		}

		// 随机取查询参数
		paramIndex := rand.Intn(len(searchParams.Params) - 1)
		reqParam.RequestId = security.UniqueId()
		reqParam.Params = searchParams.Params[paramIndex]

		resp, err := request.Request(reqParam)
		if err != nil {
			fmt.Printf("request search error, paramsIndex:%d, err msg:%s\n", paramIndex, err.Error())
		}

		fmt.Printf("thread:%d, reqIndex:%d, paramsIndex:%d, resp msg:%s\n", threadIndex, i, paramIndex, resp)
	}

	fmt.Printf("stress test thread:%d end", threadIndex)

	ch <- 1
}
