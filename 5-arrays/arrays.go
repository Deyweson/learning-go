package arrays

import "fmt"

func arrays() {
	// Declaração de arrays
	// os [] indica o que é um array e é
	// seguido pelo tipo que vai ser o array
	// o numero entre os [] significa o tamanho dele
	// o tamanho do array precisa ser constantes, não pode ser uma var
	arr := [3]int{}

	// Declaração de um array já com os elementos
	arr2 := [3]int{1, 2, 3}

	// Declaração de um array definindo valos em posições escolhidas
	arr3 := [10]int{5: 100, 6: 200}

	fmt.Println(arr, arr2, arr3)
}
