package main

type State interface {
	addItem(count int) error     //添加货物
	requestItem() error          //xxZE商品
	insertMoney(money int) error //插入钱
	dispenseItem() error         //提供商品
}
