/**
 *@Description
 *@ClassName methods
 *@Date 2021/5/10 下午8:02
 *@Author ckhero
 */

package main

import "fmt"

type Test struct {

}

func(*Test) TestP() {
	fmt.Println("test P")
}

func(Test) Test() {
	fmt.Println("test")

}

func main() {

	t := Test{}
	// 编译器不做任何处理
	t.Test()
	// 遍历会转化成 (&t).TestP()
	t.TestP()
	(&t).TestP()
}

