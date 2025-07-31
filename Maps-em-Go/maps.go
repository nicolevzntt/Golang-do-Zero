package main

import "fmt"

//Maps: Heterogêneos
//pode misturar tipos
// estrutura chave - valor
//[key] = value
// chave tem um tipo e o valor pode ter outro
//map[k]v -> k = chave, v = valor

func main() {
	idade := map[string]int{}
	idade["nick"] = 24
	idade["lineu"] = 5
	fmt.Println(idade)
	fmt.Println(idade["nick"])
	fmt.Println(idade["lineu"])

	anoNasc := map[string]int{
		"nick":  2000,
		"lineu": 2020,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["nick"])
	fmt.Println(anoNasc["lineu"])
	anoNasc["lupita"] = 2024
	fmt.Println(anoNasc)

}
