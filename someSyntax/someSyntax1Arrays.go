package main

import "fmt"

func main() {
	arrSol := [...]int{1} // ... se refera ca arrayul e capabil sa prea tot ce i se da

	fmt.Println(arrSol)

	arrSol[0] = 2

	fmt.Println(arrSol)

	// in golang, arrayurile sunt tratate ca si date nu ca adrese de date
	// deci daca se face arr2 = arr1, se vor copia valorile
	// ex
	arr1 := [...]int{1, 2}
	arr2 := arr1
	arr2[1] = 0
	fmt.Println(arr1)
	fmt.Println(arr2)

	// slice

	slice1 := []int{1, 2, 3, 4}
	slice2 := slice1
	slice1[0] = 4
	fmt.Println(slice1)
	fmt.Println(slice2) // din print slice se comporta ca un array in python
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice3 := slice1[0 : len(slice1)-1]
	fmt.Println(slice3)

	// se poate folosi make pentru a crea un slice

	sl1 := make([]int, 5)
	fmt.Println(sl1)
	fmt.Println(len(sl1))
	fmt.Println(cap(sl1))

	sl2 := make([]int, 5, 50) // type, len, cap (args)
	fmt.Println(sl2)
	fmt.Println(len(sl2))
	fmt.Println(cap(sl2))
	// de ce e nevoie de capacitate
	// explicatie: sa zicem ca avem un slice gol
	// asta inseamna ca initial are len si cap = 0
	// dar daca i se adauga un elment sunt sanse ca acesta sa aiba
	// len 1 si cap > 1 pt ca este relocat intr-o zona de memorie ce permite
	// stocarea unui array de lungime min 1
	// Obsv. daca se face relocare multipla, (adica daca adaugam un element nou in
	// slice si len devine > decat cap), se consuma multe resurse
	// astfel cel mai bine este ca initial sa se seteze si cap (capacitatea max)
	// daca aceasta se stie
	sl2 = append(sl2, 1, 2, 3) // in append, dupa numele sliceului, poti adauga oricate valoir vrei sa introduci in slice
	fmt.Println(sl2)
	fmt.Println(len(sl2))
	fmt.Println(cap(sl2))
	// u cannot append 2 lices
	// dar poti face un workaround
	// cu spread operator
	// ex
	sl3 := append(sl2, sl1...)
	fmt.Println(sl2)
	fmt.Println(sl3)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))

	// eliminare element din slice
	// ex elimin al 6 lea elemement
	sl4 := append(sl3[0:5], sl3[6:]...)
	fmt.Println(sl4)
	fmt.Println(len(sl4))
	fmt.Println(cap(sl4))
}
