package main

import (
	"fmt"
)

func main() {
	fmt.Println("--Inteface--")
	// interface nu descriu date, ele descriu behaiviours
	// trebui numita dupa ce face, daca in interfata se gaseste mteoda Wirte, interfata se va numi Writer

	var inc Kkt = new(IncData) // daca metodele utilizeaza pointeri ex (val *IncData)
	// initializezi cu un pointer like &IncData(0) daca nu initializezi normal IncData(0)

	inc.Increment()
	inc.Increment()
	inc.Decrement()

	fmt.Println(inc.GetValue())
	inc.GetKKT()

	// var test interface {} = .... // asta este un empty interface si in ea poti baga orice
	// best practices: https://youtu.be/YS4e4q9oBaU?t=19437

}

type Incrementer interface {
	Increment()
	Decrement()
	GetValue() int
}

type Kkt interface {
	Incrementer
	GetKKT()
}

type IncData int

func (val *IncData) Increment() {
	*val = *val + 1
}

func (val *IncData) Decrement() {
	*val = *val - 1
}

func (val *IncData) GetValue() int {
	return int(*val)
}

func (val *IncData) GetKKT() {
	fmt.Println("KKT")
}
