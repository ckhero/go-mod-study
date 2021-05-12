/**
 *@Description
 *@ClassName main
 *@Date 2021/5/10 下午2:44
 *@Author ckhero
 */

package main

import "fmt"

func main() {
	fmt.Println(22222)
	<- GetCh()
	fmt.Println(3333)
}

func GetCh() chan int {
	return nil
}

func TestNewNonEscape() {
	a := new(int)
	*a = 1
}


func TestNewEscape() *int {
	a := new(int)
	*a = 1
	return a
}

func TestMakeNonEscape() {
	a := make(map[string]string)
	a["test"] = "2"
}

func TestMakeEscape() map[string]string {
	a := make(map[string]string)
	a["test"] = "2"
	return a
}