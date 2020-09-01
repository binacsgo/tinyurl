package main

import (
	"fmt"

	"github.com/binacsgo/tinyurl"
)

func main() {
	urls := []string{
		"https://binacs.cn",
		"https://prometheus.binacs.cn",
		"https://grafana.binacs.cn",
		"https://jenkins.binacs.cn",
		"https://docs.binacs.cn",
		"https://kiki.binacs.cn",
	}
	for _, url := range urls {
		fmt.Println(tinyurl.GetMD5Impl().EncodeURL(url))
	}
}
