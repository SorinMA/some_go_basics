package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitGroup = sync.WaitGroup{} // creaza un wait group din care default face parte si main

func main() {
	fmt.Println("https://youtu.be/YS4e4q9oBaU?t=21199")
	fmt.Println(runtime.GOMAXPROCS(-1)) // afiseaza numarul de nuclee disponibile
	// cu runtime.GOMAXPROCS(value) setezi numarul de procs, value >=1
	fmt.Println("Goroutines Xd")
	// cand creezi o go rutina, ex
	/*
		go func(msg string) {

			...
		}(msg) // foloseste argumente pentru a apela functia
	*/

	sum1, sum2, sum3, sum4, sum5, sum6, sum7, sum8, sum9 := 0, 0, 0, 0, 0, 0, 0, 0, 0

	waitGroup.Add(9) // adaugam la waitgroup 9 pentru ca urmeaza sa procesam in paralel 9 functii
	go addInterval(0, 10, &sum1)
	go addInterval(0, 20, &sum2)
	go addInterval(0, 3, &sum3)
	go addInterval(0, 41, &sum4)
	go addInterval(0, 5, &sum5)
	go addInterval(0, 16, &sum6)
	go addInterval(0, 7, &sum7)
	go addInterval(0, 88, &sum8)
	go addInterval(0, 29, &sum9)
	waitGroup.Wait() // asteptam pana cele 9 functii au terminat treaba

	fmt.Println(sum1, sum2, sum3, sum4, sum5, sum6, sum7, sum8, sum9)

	var sum = 0

	waitGroup.Add(5)
	go addIntM(0, 1, &sum)
	go addIntM(0, 2, &sum)
	go addIntM(0, 3, &sum)
	go addIntM(0, 4, &sum)
	go addIntM(0, 5, &sum)
	waitGroup.Wait()

	fmt.Println(sum)
	// pentru testare pt a  avea o vizualiare a cum se executa functiile
	// poti rula go run -race app.go // face un data race
}

func addIntM(start, stop int, sum *int) {
	auxSum := 0
	for start < stop {
		auxSum = auxSum + start
		start = start + 1
	}

	*sum = *sum + auxSum
	waitGroup.Done()
}

func addInterval(startInt, stopInt int, sum *int) {
	*sum = 0

	for i := startInt; i < stopInt; i++ {
		*sum = *sum + 1
	}
	fmt.Println("done", startInt, stopInt)
	waitGroup.Done() // cand se face done, se decrementeaza waitgroupul
}
