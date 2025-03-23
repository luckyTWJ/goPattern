package main

import (
	"fmt"
	"sync"
)

// 保证类非公有的 外界不能通过类创建实例
type singleton1 struct {
}

// golang 中没有常指针概念 只能通过指针私有化不让外部模块访问
// 饿汉模式
var instance1 *singleton1

// 定义一个锁
var lock sync.Mutex

var once sync.Once

// 标记
var initialized uint32

func GetInstance1() *singleton1 {

	once.Do(func() {
		instance1 = new(singleton1)
	})
	//懒加载
	//if instance1 == nil {
	//	instance1 = new(singleton1)
	//	return instance1
	//}
	return instance1
}

func (singleton *singleton1) Foo() {
	fmt.Println("--singleton-Foo()--")
}
func main() {
	s := GetInstance1()
	s.Foo()

	s2 := GetInstance1()
	s2.Foo()

	if s == s2 {
		fmt.Println("s == s2")
	} else {
		fmt.Println("s != s2")
	}
}
