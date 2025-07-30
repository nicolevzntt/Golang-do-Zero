package main

import "fmt"

func main() {

	// Array- tamanho fixo
	var array [2]string // define um array com 2 strings
	array[0] = "hello"  // define o índice de uma das strings
	array[1] = "world"
	// fmt.Println(array[0])
	// fmt.Println(array[1])
	// fmt.Println(array[0], array[1])
	// fmt.Println(array)

	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[0:4]) // printar da posição 0 até a posição 3 (antes da posição 3)
	fmt.Println(numPrimos[1])
	fmt.Println(numPrimos[2])
	fmt.Println(numPrimos[3])
	fmt.Println(numPrimos[4])
	fmt.Println(numPrimos[5])

}

//Array
// Tamanho fixo, de zero ou mais elementos do mesmo tipo
// acessamos os valores com índice a[0], a[1]..
// função embutida len() retorna o tamanho do array
// por conta do tamanho fixo, não é tão usado. só em casos específicos
