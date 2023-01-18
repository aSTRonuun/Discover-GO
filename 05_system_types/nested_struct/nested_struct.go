package main

import "fmt"

type item struct {
	productId int
	quantity  int
	price     float64
}

type order struct {
	userId int
	itens  []item
}

func (o order) totalValue() float64 {
	total := 0.0
	for _, item := range o.itens {
		total += item.price * float64(item.quantity)
	}
	return total
}

func main() {

	order := order{
		userId: 1,
		itens: []item{
			item{1, 2, 199.90},
			item{2, 2, 299.90},
			item{3, 1, 499.90},
		},
	}

	fmt.Printf("The total value of order is %.2f", order.totalValue())
}
