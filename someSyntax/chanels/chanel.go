package main

import (
	"fmt"
	"sync"
)

var waitG = sync.WaitGroup{}

func main() {
	fmt.Println("chan")

	chanel := make(chan int, 10) // cream un chanel ce ia valori int, un chanel cu buffer de 10 valori

	waitG.Add(3)
	go func(chanel <-chan int) { // gorutina asta primeste date de la channel
		for val := range chanel { // se deosebire de un for normla,range returneaza doar valoare nu si indexul
			fmt.Println(val)
		}
		waitG.Done()
	}(chanel)

	// go func(chanel <-chan int) { // to access values one by one
	// 	fmt.Println(<-chanel, "first value")
	// 	fmt.Println(<-chanel, "scnd value")
	// 	fmt.Println(<-chanel, "third value")
	// 	waitG.Done()
	// }(chanel)

	go func(chanel chan<- int) { // asta baga date in cahnnel
		chanel <- 10
		chanel <- 33
		chanel <- 123
		close(chanel) // inchide canalul ca sa stie for-ul de mai sus sa nu mai asculte la canalul asta
		// daca incerci sa bagi valoare in chanel dupa asta iti da eroare
		waitG.Done()
	}(chanel)
	waitG.Wait()

	// asta e util cand ai un IOT ca face un burst de infos
	// si trimite multe si trebuie procesate

	// pentru a verifica ca chanelul nu e inchis poti proceda ca la map
	/*
		if val, ok :=  <-channel ; ok {
			fmt.println(val)
		}
	*/

	fmt.Println("Done")
}
