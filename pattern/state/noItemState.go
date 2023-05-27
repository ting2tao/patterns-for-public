package main

import "fmt"

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (no *NoItemState) addItem(count int) error {
	no.vendingMachine.incrementItemCount(count)
	no.vendingMachine.setState(no.vendingMachine.hasItem)
	return nil
} //添加货物
func (no *NoItemState) requestItem() error {
	return fmt.Errorf("item out of stock")
} //xxZE商品
func (no *NoItemState) insertMoney(money int) error {
	return fmt.Errorf("item out of stock")
} //插入钱
func (no *NoItemState) dispenseItem() error {
	return fmt.Errorf("item out of stock")
} //提供商品
