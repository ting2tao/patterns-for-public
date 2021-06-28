package pattern

import "math"

// 策略模式
// 根据不同的策略获取不同的价格
const (
	ORIGINAL = iota
	DISCOUNT
	FullPER
	FullDISCOUNT
)

// 获取价格的接口
type PriceInterface interface {
	GetPrice(price float64) float64
}

// 原价
type Original struct {
}

func (ori *Original) GetPrice(price float64) float64 {
	return price
}

// 折扣
type Discount struct {
	Percent float64 `json:"percent"`
}

func (dis *Discount) GetPrice(price float64) float64 {
	return price * dis.Percent
}

// 每满减 不按比列折扣
type Reduction struct {
	Full     float64 `json:"full"`
	Discount float64 `json:"discount"`
}

func (red *Reduction) GetPrice(price float64) float64 {
	if price < red.Full || price == 0 || red.Full == 0 {
		return price
	}
	return price - red.Discount*math.Ceil(price/red.Full)
}

// 仅满减
type FullReduction struct {
	Full     float64 `json:"full"`
	Discount float64 `json:"discount"`
}

func (red *FullReduction) GetPrice(price float64) float64 {
	if price < red.Full || price == 0 {
		return price
	}
	return price - red.Discount
}

// 价格工厂
type PriceFactory struct {
}

func (fa PriceFactory) GetOrderPrice(offType int) PriceInterface {
	var priceInterface PriceInterface
	switch offType {
	case ORIGINAL:
		priceInterface = &Original{}
	case DISCOUNT:
		priceInterface = &Discount{Percent: 0.8}
	case FullPER:
		priceInterface = &Reduction{
			Full:     100,
			Discount: 5,
		}
	case FullDISCOUNT:
		priceInterface = &FullReduction{
			Full:     200,
			Discount: 20,
		}
	}
	return priceInterface
}
