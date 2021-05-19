/**
 *@Description
 *@ClassName leak
 *@Date 2021/5/13 下午7:12
 *@Author ckhero
 */

package main

import (
	"fmt"
	_ "net/http/pprof"
	"net/url"
)

func main() {
	test:= "http://www.361way.com/golang-urlencode-urldecode/6390.html"
	fmt.Println(url.QueryEscape(test))
	//// 开启pprof
	//go func() {
	//	ip := "0.0.0.0:6061"
	//	if err := http.ListenAndServe(ip, nil); err != nil {
	//		fmt.Printf("start pprof failed on %s\n", ip)
	//		os.Exit(1)
	//	}
	//}()
	//go func() {
	//	for  {
	//		for i :=0;i < 100;i++ {
	//			url := "test"
	//			logger, _ := zap.NewProduction()
	//			defer logger.Sync() // flushes buffer, if any
	//			sugar := logger.Sugar()
	//			sugar.Infow("failed to fetch URL",
	//				// Structured context as loosely typed key-value pairs.
	//				"url", url,
	//				"attempt", 3,
	//				"backoff", time.Second,
	//			)
	//			sugar.Infof("Failed to fetch URL: %s", url)
	//		}
	//	}
	//}()
 //ch := make(chan int)
 //<- ch
}
