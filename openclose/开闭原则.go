package main

import "fmt"

type IBanker interface {
	DoBusiness()
}
type SaveBanker struct {
}

// 存款业务员
func (saveBanker *SaveBanker) DoBusiness() {
	fmt.Println("存款业务...")
}

// 转账业务员
type TransferBanker struct {
}

// 转账业务员
func (tb *TransferBanker) DoBusiness() {
	fmt.Println("转账业务...")

}

// 股票业务员
type StockBanker struct {
}

func (sb *StockBanker) DoBusiness() {
	fmt.Println("股票业务...")
}

//实现架构层（基于抽象层进行业务封装 针对interface接口进行封装）

func BankBusiness(ibanker IBanker) {
	ibanker.DoBusiness()
}

func main() {
	//开闭原则 ：对扩展开放，对修改关闭

	sb := SaveBanker{}
	sb.DoBusiness()
	fmt.Println("-----------------------")

	tb := &TransferBanker{}
	tb.DoBusiness()
	fmt.Println("-----------------------")

	sb2 := &StockBanker{}
	sb2.DoBusiness()
	fmt.Println("-----------------------")

	BankBusiness(&SaveBanker{})
	BankBusiness(&TransferBanker{})
	BankBusiness(&StockBanker{})

}
