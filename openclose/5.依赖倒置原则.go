package main

import "fmt"

// ---------------抽象层-----------------
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

//---------------实现层-----------------

type BenZ struct {
}

func (benz *BenZ) Run() {
	println("benz run....")
}

type Bmw struct {
}

func (bmw *Bmw) Run() {
	println("bmw run.....")
}

type Zhangsan struct {
}

func (zhangsan *Zhangsan) Drive(car Car) {
	fmt.Println("zhangsan is driving...")
	car.Run()
}

type LiSi struct {
}

func (lisi *LiSi) Drive(car Car) {
	fmt.Println("lisi is driving...")
	car.Run()
}

type wang5 struct {
}

func (wan5 *wang5) Drive(car Car) {
	fmt.Println("wang5 is driving...")
	car.Run()
}

//---------------业务逻辑层-----------------

func main() {
	//	张三开奔驰
	var zhangsan Driver
	//zhangsan = &Zhangsan{}
	zhangsan = new(Zhangsan)
	var benz Car
	benz = &BenZ{}
	zhangsan.Drive(benz)

	var c Car
	c = &Bmw{}
	zhangsan.Drive(c)

	//	李四开宝马
}
