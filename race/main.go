/**
 *@Description
 *@ClassName main
 *@Date 2021/4/22 下午4:46
 *@Author ckhero
 */

package main

import (
	"fmt"
	)

var a, b int

func f() {
	a = 1
	b = 2
}

func g() {
	print(b)
	print(a)
}

func main() {
	go f()
	g()
}
func Go(fn func(...interface{})) {
	go func() {
		fn()
	}()
}
func race1() {
	a := 1;
	go func() {
		a = 2
	}()

	a = 3
	fmt.Println(a)
}