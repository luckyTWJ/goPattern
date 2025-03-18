package main

import "fmt"

//司机张三 李四 汽车 宝马 奔驰

type Benz struct {
}

func (ben *Benz) Run() {
	fmt.Println("benz is running...")
}

type BMW struct {
}

func (bmw *BMW) Run() {
	fmt.Println("bmw is running...")

}

type Zhang3 struct {
}

func (z3 *Zhang3) DriveBenz(benz *Benz) {
	fmt.Println("zhang3 is driving benz...")
	benz.Run()

}

type Li4 struct {
}

func (l4 *Li4) DriveBenz(benz *Benz) {
	fmt.Println("li4 is driving benz...")
	benz.Run()
}
func (li4 *Li4) DriveBMW(bmw *BMW) {
	fmt.Println("li4 is driving bmw...")
	bmw.Run()
}

func main() {
	//美增加一种车型 都要重新新增对应的 drive方法
	benz := Benz{}
	bmw := BMW{}
	zhang3 := Zhang3{}
	li4 := Li4{}
	zhang3.DriveBenz(&benz)
	li4.DriveBMW(&bmw)
	li4.DriveBenz(&benz)
}
