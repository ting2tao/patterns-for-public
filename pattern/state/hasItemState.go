package main

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (i *HasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
} //添加货物
func (i *HasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Item requestd\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
} //xxZE商品
func (i *HasItemState) insertMoney(money int) error {
	return fmt.Errorf("please select item first")
} //插入钱
func (i *HasItemState) dispenseItem() error {
	return fmt.Errorf("please select item first")
} //提供商品
