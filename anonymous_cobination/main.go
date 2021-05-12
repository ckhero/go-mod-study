/**
 *@Description
 *@ClassName main
 *@Date 2021/5/10 下午8:42
 *@Author ckhero
 */

package main

import "fmt"

type Test struct {
	Test1
	Test2
}

type Test1 struct {

}
func(*Test1) Test() {
	fmt.Println("test1")
}
type Test2 struct {

}

func(*Test2) Test() {
	fmt.Println("test1")
}

func main() {
   var a rune
	//t := Test{}
	//t.Test()
}