/**
 *@Description
 *@ClassName _unsafe
 *@Date 2021/5/12 下午7:26
 *@Author ckhero
 */

package main

import (
	"fmt"
	"unsafe"
)

func main() {

	DemoFour()
}

// 测试 风险
func DemoOne() {
	a := uint64(32)
	p1 := &a
	// 创建一个 *uint8 的指针 去指向一个 uint64的变量。而这个变量是使用p1 去访问的。有风险
	p2 := (*uint8)(unsafe.Pointer(p1))
	fmt.Println("p1", *p1)
	fmt.Println("p2", *p2)
	// 22222211 超出 uint8 的范围，*p2 输出异常
	*p1 = 22222211
	fmt.Println("p1", *p1)
	fmt.Println("p2", *p2)

	*p1 = 100
	fmt.Println("p1", *p1)
	fmt.Println("p2", *p2)
}

// 利用 unsafe  获取数组中的数据, 以及修改
func DemoTwo() {
	arr := []int32{1, 2, 3}
	p1 := &arr[0]
	memoryAddress := uintptr(unsafe.Pointer(p1))
	for i := 0; i < len(arr); i ++ {
		p1 = (*int32)(unsafe.Pointer(memoryAddress))
		fmt.Println(*p1)
		*p1 += 1
		memoryAddress = uintptr(unsafe.Pointer(p1)) + unsafe.Sizeof(arr[0])

	}
	fmt.Println(arr)
}

type TypeA struct {
	Name string
	Name2 string
	name string
	age int
}

type TypeB struct {
	Name2 string
}

// 测试结构体的类型转换
func DemoThree() {
	a := TypeA{Name:"2222", Name2:"333333"}
	p1 := &a
	p2 := (*TypeB)(unsafe.Pointer(p1))
	fmt.Println(*p1)
	fmt.Println(*p2)
}

// 获取私有变量值
func DemoFour() {
	a := TypeA{
		//Name:  "Name",
		Name2: "Name2",
		name:  "name",
		age:   22,
	}
	fmt.Println("name align",unsafe.Alignof(a.Name))
	fmt.Println("name offset",unsafe.Offsetof(a.Name))
	fmt.Println("name2 align",unsafe.Alignof(a.Name2))
	fmt.Println("name2 offset",unsafe.Offsetof(a.Name2))
	p1 := unsafe.Pointer(&a)
	// 获取age
	agePtr := unsafe.Pointer(uintptr(p1) + unsafe.Offsetof(a.age))
	fmt.Println(*(*int)(agePtr))
	// 修改
	*(*int)(agePtr) = 100
	fmt.Println(a)
}
