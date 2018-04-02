package main

import (
	"fmt"
	"stress"
)

func main() {
	err := stress.StressGo()

	if err != nil {
		fmt.Printf("stress go error:%s", err.Error())
	}
}