package main

import "fmt"

// 业务逻辑层 -> 依赖工厂模式  -> 依赖基础模块
// ----------------抽象层---------------------------------
type AbstractApple interface {
	ShowApple()
}
type AbstractBanana interface {
	ShowBanana()
}

type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
}

// ----------------实现层---------------------------------

type ChinaApple struct {
}

func (chinaApple *ChinaApple) ShowApple() {
	fmt.Println("我是中国苹果")
}

type ChinaBanana struct {
}

func (chinaBanana *ChinaBanana) ShowBanana() {
	fmt.Println("我是中国香蕉")
}

type ChinaFactory struct {
}

func (chinaFactory *ChinaFactory) CreateApple() AbstractApple {
	//面向接口编程
	var chinaApple AbstractApple
	chinaApple = &ChinaApple{}
	return chinaApple
}
func (ChinaFactory *ChinaFactory) CreateBanana() AbstractBanana {
	var chinaBanana AbstractBanana
	chinaBanana = &ChinaBanana{}
	return chinaBanana
}

// 父类指着指向子类对象
func main() {

}
