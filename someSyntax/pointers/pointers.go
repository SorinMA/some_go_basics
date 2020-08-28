package main

import (
	"fmt"
)

func main() {
	fmt.Println("Porinters")

	// basics
	a := 2
	b := &a
	fmt.Println(&a, a, b, *b)
	*b = 4
	fmt.Println(&a, a, b, *b)
	// chestia de mai sus e echivalenta cu:
	var aa int = 2
	var bb *int = &aa // * se refera ca este un pointer iar & ia adresa de memorie
	// pentru a extrage valoarea din pointerul; bb folosim *bb (adica operatorul de
	// deregerentiere)
	fmt.Println(&aa, aa, bb, *bb)
	*bb = 4
	fmt.Println(&aa, aa, bb, *bb)

	// --- arithmetics
	// go nu te lasa sa faci c := &a + 4 (adica nu te lasa sa faci operatii artimetice pe adrese de memorie)
	// by default, dar daca folosesti pachetul 'unsafe' poti face asta

	/*

		Structuri
		var someS *myStruct // aici someS are valoare nil (ca nu e initializat)
		someS = new(myStruct) // asta init default cu 0 tot
		 sau someS = &myStruct{trash:3} // pt initializare
		apoi:
		someS.trash = 4 este echivalent cu (*someS).trash = 4
	*/
	// daca avem doua silceuri a si b si facem b:= a, b va pointa la prima vloare din sliceul a (adica face o deplasare prin memorie)
	// la fel si la map

	//test

	a1 := []int{1, 2, 3}
	b1 := append(a1)
	b2 := append(a1, -1)
	b3 := append([]int{}, a1...)
	c := a1
	a1[0] = 22
	fmt.Println(a1, b1, c) //all ar the same ,b1 and c1 point to a1
	// but if i do, pt ca pare ca pointeaza la noua zona de memorie unde se afla noul slice b2
	// creat cu elementele din a1
	fmt.Println(a1, b1, c, b2)

	// inca ceva
	fmt.Println(b3) // asa pare o metoda buna de a copia valorile dintr-un slice in altul

}
