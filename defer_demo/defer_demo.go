/**
 *@Description
 *@ClassName defer_demo
 *@Date 2021/5/17 下午2:16
 *@Author ckhero
 */

package main

import "fmt"

func main() {
	fmt.Println(f())
	//fmt.Println(f1())
}

//func f1() int {
//	a := 1
//	for i :=0;i < 5;i++ {
//		defer func() {
//			fmt.Println(i)
//		}()
//	}
//	return a
//}
func f() (r int) {
	t := 5
	defer func() {
		t = t + 5
		fmt.Println("t", t)
	}()
	return t
}
//
//func f1() (r int) {
//	defer func(r int) {
//		r = r + 5
//		fmt.Println("r", r)
//	}(r)
//	return 1
//}