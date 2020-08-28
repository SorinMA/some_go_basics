package main

import (
	"fmt"
)

func main() {
	fmt.Println("func")
	// orice ce incepe cu litera mare gen 'L' e public si daca nu este vizibil doar in packet
	str1 := "sori"
	str2 := "ggy"

	structRes := someShit(&str1, &str2, 10, 2, 3, 4)

	fmt.Println(structRes)

	// un pattern comun in go este ca daca ai o functie care stii ca poate sa aiba probleme
	// ge impartire la 0, sa ai inca un parametru de return numir err
	// ex

	rez, err := errEx(1, 2, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rez, "rez")
	}

	// functii variabile
	// ex

	var exFct func(int, int) int
	exFct = func(a, b int) int {
		return a + b
	}
	fmt.Println(exFct(1, 2))

	// mehtod

	sd := kkt{
		name: "sss",
		lvl:  2,
	}
	sd.getPoop()
	fmt.Println("setting pop to 4")
	sd.setPoop(4)
	sd.getPoop()
}

// method, aka functie pt structura

type kkt struct {
	name string
	lvl  int
}

func (k kkt) getPoop() { // aka e ca si cum ai avea functia asta pt structura kkt
	fmt.Println(k) // e ca la metoda la clasa
	// obsv: in cazul (k kkt) se face o copie a structurii ce apeleaza functia
	// daca se vrea a se folosi structura fara a fi copiata folosim
	// (k *kkt)
}

func (k *kkt) setPoop(t int) { // aka e ca si cum ai avea functia asta pt structura kkt
	k.lvl = t // e ca la metoda la clasa
	// obsv: in cazul (k kkt) se face o copie a structurii ce apeleaza functia
	// daca se vrea a se folosi structura fara a fi copiata folosim
	// (k *kkt)
}

func errEx(val1, val2, val3 float32) (float32, error) {
	if val3 == 2.0 {
		return -1.0, fmt.Errorf("u deaf monkey")
	}

	return 0.0, nil
}

func someShit(str1, str2 *string, vals ...int) interface{} {
	*str1 = *str1 + "_GAY"
	*str2 = *str2 + "Gay"
	for _, val := range vals {
		fmt.Print(val, " ")
	}

	return struct {
		name     string
		lastName string
		grades   []int
	}{
		name:     *str1,
		lastName: *str2,
		grades:   vals,
	}
}
