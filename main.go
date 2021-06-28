package main

import (
	"fmt"
	"github.com/slovty/patterns-for-public/pattern"
)

func main() {
	fmt.Println("设计模式")
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
