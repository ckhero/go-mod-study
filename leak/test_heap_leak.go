/**
 *@Description
 *@ClassName test_heap_leak
 *@Date 2021/5/13 下午8:05
 *@Author ckhero
 */

package main

import "time"

func testHeap() {
	tick := time.Tick(time.Second/ 1000)
	var buf []byte
	for range tick {
		buf = append(buf, make([]byte, 1024*1024)...)
	}
}
