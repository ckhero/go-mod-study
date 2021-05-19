/**
 *@Description
 *@ClassName race_demo
 *@Date 2021/5/17 下午4:23
 *@Author ckhero
 */

package main

import "fmt"

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
)
var Enbale = false
func main() {
	fmt.Println(mutexLocked)
	fmt.Println(mutexWoken)
	fmt.Println(mutexStarving|mutexLocked)
	a := 1

	go test(&a)

	go test2(&a)
	Enbale = true

}

func test(a *int) {
	if Enbale {
		*a = 2
	}
}

func test2(a *int) {
	if Enbale {
		*a = 3
	}
}