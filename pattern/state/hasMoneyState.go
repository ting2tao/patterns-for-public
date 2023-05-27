package main

import "fmt"

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) addItem(count int) error {

	return fmt.Errorf("Item dispense in progress ")
} //添加货物
func (i *HasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress ")
} //xxZE商品
func (i *HasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock ")
} //插入钱
func (i *HasMoneyState) dispenseItem() error {
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
} //提供商品
