package main

import (
	"fmt"
	"math"
)

type Money int64

func (m Money) string() string {
	rub := m / 100
	kop := m % 100
	return fmt.Sprintf("%d руб. %02d коп.", rub, kop)
}

func (m Money) add(money Money) Money {
	return m + money
}

type Circle struct {
	Radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	price1 := Money(30030)
	fmt.Println("price1 =", price1.string())
	price2 := Money(2020)
	fmt.Println("price2 =", price2.string())
	totalPrice := price1.add(price2)
	fmt.Println("totalPrice =", totalPrice.string())

	circle := &Circle{Radius: 10}
	fmt.Println("Area:", circle.area())
	fmt.Println("Perimeter:", circle.perimeter())
}

