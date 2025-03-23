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
type ChinaBanana struct {
}

func (chinaApple *ChinaApple) ShowApple() {
	fmt.Println("我是中国苹果")
}
func (chinaBanana *ChinaBanana) ShowBanana() {
	fmt.Println("我是中国香蕉")
}

type ChinaFactory struct {
}

func (chinaFactory *ChinaFactory) CreateApple() AbstractApple {
	//面向接口编程
	var chinaApple AbstractApple
	chinaApple = new(ChinaApple)
	return chinaApple
}
func (ChinaFactory *ChinaFactory) CreateBanana() AbstractBanana {
	var chinaBanana AbstractBanana
	chinaBanana = &ChinaBanana{}
	return chinaBanana
}

type UsaApple struct {
}
type UsaBanana struct {
}

func (usaApple *UsaApple) ShowApple() {
	fmt.Println("我是美国苹果")
}

type UsaFactory struct {
}

func (usaBanana *UsaBanana) ShowBanana() {
	fmt.Println("我是美国香蕉...")
}
func (usaFactory *UsaFactory) CreateApple() AbstractApple {
	var usaApple AbstractApple
	usaApple = new(UsaApple)
	return usaApple
}
func (usaFactory *UsaFactory) CreateBanana() AbstractBanana {
	var usaBanana AbstractBanana
	usaBanana = &UsaBanana{}
	return usaBanana
}

// 父类指着指向子类对象
func main() {
	//	创建一个美国工厂
	var usaFactory AbstractFactory
	usaFactory = new(UsaFactory)

	//生产美国苹果
	var usaApple AbstractApple
	usaApple = usaFactory.CreateApple()
	usaApple.ShowApple()
	//生产美国香蕉
	var usaBanana AbstractBanana
	usaBanana = usaFactory.CreateBanana()
	usaBanana.ShowBanana()

	// 中国工厂
	var chainaFactory AbstractFactory
	chainaFactory = new(ChinaFactory)

	// 生产中国苹果   父类引用调用子类方法   多态
	var chinaApple AbstractApple
	chinaApple = chainaFactory.CreateApple()
	chinaApple.ShowApple()

}
