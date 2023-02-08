package main

import (
	"fmt"
	"github.com/slovty/patterns-for-public/pattern"
	"github.com/slovty/patterns-for-public/singleton"
)

func main() {
	Factory()
	Singleton()
}

// Factory 工厂模式
func Factory() {
	fmt.Println("工厂模式")
	var pf pattern.PriceFactory
	pi := pf.GetOrderPrice(pattern.ORIGINAL)
	fmt.Println("原价", pi.GetPrice(200))
	pi = pf.GetOrderPrice(pattern.DISCOUNT)
	fmt.Println("8折", pi.GetPrice(200))
	pi = pf.GetOrderPrice(pattern.FullPER)
	fmt.Println("每满减", pi.GetPrice(200))
	pi = pf.GetOrderPrice(pattern.FullDISCOUNT)
	fmt.Println("达满减", pi.GetPrice(200))
}

func Singleton() {
	fmt.Println("单例模式")
	singleton.ClientOnce()
	singleton.ClientLock()
}
