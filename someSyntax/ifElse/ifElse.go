package main

import (
	"fmt"
)

func someFuncToTestBook() int {
	fmt.Println("in func")
	return 2
}

func main() {
	fmt.Println("ifElse:")

	someMap := map[int]int{
		2: 3,
		5: 1,
	}

	fmt.Println(someMap)

	// if simplu

	if someMap[2] >= 0 { // in go e neaparat sa pui acolade
		fmt.Println("yes")
	}

	// if cu instantiere
	if val, ok := someMap[4]; ok {
		fmt.Println(val, ok)
	} else {
		fmt.Println("Not ok")
	}

	if val, ok := someMap[2]; ok || someFuncToTestBook() > 0 {
		fmt.Println("in main", val)
	} // se obsertva ca someFuncToTestBook() nu mai e apelata
	// asta pentru ca ok este true si daca s-a gasit o valoare de true
	// || (orul) nu mai continua executarea retrului de cod din conditie
	// same on &&, daca prima valoare din && este true o verifica si pe a doua
	// daca prima val e false face scurtcircuit si nu o mai verifica si pe a doua
	if someFuncToTestBook() == 0 {
		fmt.Print(1)
	} else if someFuncToTestBook() == 1 {
		fmt.Print(2)
	} else {
		fmt.Print(3)
	}

	// switch
	fmt.Println("---switch----")
	i := 23
	switch j := 4; j {
	case 1, 4, 5:
		i = i + 1
		fmt.Println(i) // in switch, nu trebui utilizat {} pt fiecare case,
		// caseurile sunt folosite ca delimitatori
		// si break este pus default dupa fiecare swithc
	case 9:
		fmt.Println(i)
	default:
		i = i - 1
		fmt.Println(i)
	}
	switch {
	case i < 10:
		fmt.Println(i)
	case i > 20:
		i = i - 10
		fmt.Println(i)
	default:
		i = i - 100
		fmt.Println(i)
	}

	fmt.Println("--typeswithc---")

	// interfatra
	// daca o var e interfata putem asigna orice cater ea
	var intr interface{}

	intr = struct {
		name string
		lvl  int
	}{
		name: "boss2",
		lvl:  2,
	}

	fmt.Println(intr)
	switch intr.(type) { // type switch
	case int32:
		fmt.Println("integer sama")
	case float32:
		fmt.Println("float chan")
	default:
		fmt.Println("struct san???")
	}

	// intr = []int{}
	// intr = append(intr, "u", "re", "gay")
	// fmt.Println(intr)
	// din ce am testat, nu merge sa reasignezi alt tip de data
	// pare sa fie ca la typescript
}
