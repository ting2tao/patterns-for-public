package main

import "fmt"

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) addItem(count int) error {
	return fmt.Errorf(" item has rerequested")
} //添加货物
func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf(" Item Dispense in progress")
} //xxZE商品
func (i *ItemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
} //插入钱
func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("please select item first")
} //提供商品
