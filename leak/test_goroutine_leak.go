/**
 *@Description
 *@ClassName test_goroutine_leak
 *@Date 2021/5/13 下午8:08
 *@Author ckhero
 */

package main

import (
	"fmt"
	"time"
)

func testGoroutinue() {
	outCh := make(chan int)
	// 死代码，永不读取
	go func() {
		if false {
			<-outCh
		}
		select {}
	}()

	// 每s起100个goroutine，goroutine会阻塞，不释放内存
	tick := time.Tick(time.Second / 1000)
	i := 0
	for range tick {
		i++
		fmt.Println(i)
		alloc1(outCh)
	}
}

func alloc1(outCh chan<- int) {
	go alloc2(outCh)
}

func alloc2(outCh chan<- int) {
	func() {
		defer fmt.Println("alloc-fm exit")
		// 分配内存，假用一下
		buf := make([]byte, 1024*1024*10)
		_ = len(buf)
		fmt.Println("alloc done")

		outCh <- 0 // 53行
	}()
}