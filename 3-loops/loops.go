package loops

import "fmt"

func loops() {

	// Declaração de um for simples / clássico
	for i := 0; i < 10; i++ {

	}

	// Declaração de um for infinito
	for {
		// Encerra o loop
		break
	}

	// Declaração de for usando range
	// range passa por todos os elementos de um array
	// ele retorna o index da vez e o elemento
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, elemento := range array {
		fmt.Println(index, elemento)
	}

}
