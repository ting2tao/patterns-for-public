package main

import (
	"flag"
	"fmt"
	"github.com/slovty/patterns-for-public/pattern"
	"github.com/slovty/patterns-for-public/pattern/observer"
	"github.com/slovty/patterns-for-public/pattern/state"
	"github.com/slovty/patterns-for-public/singleton"
	"log"
)

var patternName *string

func init() {
	patternName = flag.String("patternName", "Observer", "设计模式的名称")
}
func main() {
	flag.Parse()
	switch *patternName {
	case "Observer":
		Observer()
	case "Factory":
		Factory()
	case "Singleton":
		Singleton()
	case "State":
		State()
	default:
		fmt.Println("No pattern has imp")
	}

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

func Observer() {
	fmt.Println("观察者模式")
	observer.ObserverClient()

}

func State() {
	fmt.Println("状态模式")
	vendingMachine := state.NewVendingMachine(1, 10)

	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
