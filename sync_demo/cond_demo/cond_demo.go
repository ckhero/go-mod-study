/**
 *@Description
 *@ClassName cond_demo
 *@Date 2021/5/19 下午7:09
 *@Author ckhero
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

type Command struct {
	sync.Mutex
}

func NewCommand() *Command {
	return &Command{}
}

func testCond() {
	cmd := NewCommand()
	cmdCond := sync.NewCond(cmd)

	fmt.Println("notify")
	cmdCond.Signal()
	time.Sleep(time.Second)
	go func() {
		cmd.Lock()
		cmdCond.Wait()
		fmt.Println(222222)
		cmd.Unlock()
	}()

	time.Sleep(time.Second)
}
func main() {

	a := 1
	go func() {
		a =2
		fmt.Println(a)
	}()
	go func() {
		a =23
		fmt.Println(a)
	}()
	time.Sleep(time.Second)
}
