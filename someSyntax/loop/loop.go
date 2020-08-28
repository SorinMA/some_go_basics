package main

import (
	"fmt"
)

func main() {
	fmt.Println("https://youtu.be/YS4e4q9oBaU?t=12077")

	// simple for

	for i := 0; i < 4; i += 2 {
		fmt.Println(i, "ex1")
	}

	// mai e optiunea cu for i, j
	for i, j := 0, 0; i < 4 && j < 10; i, j = i+j, 2*j+1 {
		fmt.Println(i, j, "ex2")
	}
	// merge pe principul de dereferentriere, sau o poti vedea ca pe o structura de date , idk lamoo

	// alte arome de for
	ij := 0

	for ; ij < 5; ij = ij + 1 {
		fmt.Println(ij, "ij")
	}

	for ij < 10 {
		fmt.Println(ij, "ijwhile")
		ij = ij + 1
	}

	for ij < 16 {
		fmt.Println(ij, "ijwhile;;")
		ij = ij + 1
	}

	for i := 0; i < 5; {
		fmt.Println(i, "just i")
		i = i + 1
	}

	ji := 0
	for {
		fmt.Println(ji, "ji")
		ji = ji + 1

		if ji > 5 {
			break
		}
	}

	// label for loop

CeLabelVreauEuPentruLoop:
	for i := 0; i < 10; i = i + 1 {
		for j := i + 1; j < 10; j = j + 1 {
			if i+j > 5 {
				fmt.Println("GONE LABEL")
				break CeLabelVreauEuPentruLoop
			}
		}
	}

	// lucru cu liste in for

	someSlice := []int{5, 6, 3}

	for key, value := range someSlice { // key value range
		fmt.Println(key, value, "kv")
	}

	someMap := map[int]int{
		0: 5,
		1: 6,
		2: 3,
	}

	for key, value := range someMap {
		fmt.Println(key, value, "kvMAP")
	}

	someString := "kkt cu sana"

	for _, letter := range someString { // folosim _ pentru ca nu ne pasa de key
		fmt.Print(string(letter))
	}

}
