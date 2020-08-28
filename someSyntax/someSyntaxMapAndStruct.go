package main

import (
	"fmt"
)

func main() {
	fmt.Println("structs and map")

	countryCap := map[string]string{
		"RO": "Bucharest",
		"FR": "Paris",
	}

	fmt.Println(countryCap) // un array este valid key dar un slice nu
	// poti folosi make pt a crea un map
	newMap := make(map[int]int)
	newMap[2] = 5
	fmt.Println(newMap)
	delete(newMap, 2) // sterge se  se afla la keya 2
	fmt.Println(newMap)

	// dar daca e sa printez newMap[2] asta mir-ar returna 0
	// pentru a verifica ca o keye exista in dict
	// folosim sintaxa cu ok ex
	valueTwo, ok := newMap[2]
	fmt.Println(valueTwo, ok) // flase spune ca nu e in map
	// purtem folosi functia len
	fmt.Println(len(countryCap))
	// maps sunt pasate prin referinta la fel ca la slices
	// https://youtu.be/YS4e4q9oBaU?t=8988

	// define structure

	type OverLord struct {
		lvl    uint8
		name   string
		enemys []OverLord
	}

	overlord1 := OverLord{
		lvl:    2,
		name:   "Kax8",
		enemys: []OverLord{},
	}
	fmt.Println(overlord1)
	overlod2 := OverLord{
		lvl:  4,
		name: "Krax01",
		enemys: []OverLord{
			overlord1,
		},
	}
	fmt.Println(overlod2)
	overlord1.enemys = append(overlord1.enemys, overlod2)
	fmt.Println(overlord1)
	fmt.Println(overlod2)

	// structury anonime

	boss1 := struct {
		name string
		lvl  int
	}{
		name: "Gaspar",
		lvl:  23,
	}

	fmt.Println(boss1)

	boss2 := boss1
	boss2.lvl = 33

	fmt.Println("SomeChanges??: ", boss1, boss2)
	// aparent struct sunt ca array, independent data structures
	// adica boss2 este un nou boss

	// in GO nu exista inheritance
	// dar exista compozitie
	// adica sa zicem ca avem structura de date
	// Animal {...}
	// si structura de date Bird
	// pentru ca Bird e un Animal
	// putem face ceva de genu
	/*
		Bird struct{
			Animal
			speed int
			canFly bool
		}

		De ce s-a introdus Animal si nu 'Animal Animal', pentru ca nu cream un camp nou,
		ci preluam campurile din Animal, (ca si cum ai avea Animal... dereferentiere)
		https://youtu.be/YS4e4q9oBaU?t=9624
		Ca obsv, compozitia nu face ca Bird sa fie un tip de animal, Bird este Bird,
		si doar preia atributele unui animal general
	*/

	// IN GOLANG  la structuri poti adauga in definita unui camp un tag
	// ex: name string `requiered max: "100"`
	// acest tag pt a putea fi accesat e nevoie de pachetul "reflect"
	// iar pentru o validare trebui o librarie intermediara
	//https://youtu.be/YS4e4q9oBaU?t=9890

}
