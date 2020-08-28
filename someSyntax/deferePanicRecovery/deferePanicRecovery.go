package main

import (
	"fmt"
)

func main() {
	fmt.Println("--defere--")
	// defere is executed LIFO
	// adica daca avem:
	/*
		defere fct1()
		defere fct2()
		defere fct3()
		se va apela
		fct3()
		fct2()
		fct1()
	*/
	// cum merger defere
	/*
			fct1()
			defere fct2()
			fct3()
			se va executa:
			fct1()
			fct3()
			fct2()

			defere e bunt pentru a inchide resurse si se executra functiile in mod LIFO
			pentru a inchide resursele simteric fata de cum au fost deschise
		big brain time
		!!!! bucatile de cod taguite cu defere se vor executa la finalul functieii in care se afla
	*/
	fmt.Println("--panic--")
	/*
		sa zicem ca avem operatia 1 / 0, aceasta operatie va panica procesul si va opri functia din run
		keywordul panic are acelasi efect si se pune in cazul in care avem o eroare
		ceva de genu panic fmt.println(err)

		daca avem niste blocuri defere pana in momentul cand se face panic,
		acestea se vor executa
	*/
	fmt.Println("--recovery--")

	// scenariu:
	/*
		sa zicem ca avem callul
		fct1()
		defere func() {
			if err := recover() ; err != nil {
				log.Println(err)
			}
		}() // asta e o functie anonima
		panic("ceva eroare")
		...

		se executa fct1
		api se intra in panic
		pentru ca s-a apelat panic
		se executa derefurile de dinaintea lui;
		cand functia anonimca se executa aceasta preia eroarea semnalata de panic
		cu recover() si poate face ce vrea cu ea xds
		big brain
	*/
}
