package main

import "fmt"

// 保证类非公有的 外界不能通过类创建实例
type singleton struct {
}

// golang 中没有常指针概念 只能通过指针私有化不让外部模块访问
// 饿汉模式
var instance *singleton = new(singleton)

func GetInstance() *singleton {
	return instance
}

func (singleton *singleton) Foo() {
	fmt.Println("--singleton-Foo()--")
}
func main() {
	s := GetInstance()
	s.Foo()

	s2 := GetInstance()
	s2.Foo()

	if s == s2 {
		fmt.Println("s == s2")
	} else {
		fmt.Println("s != s2")
	}
}
