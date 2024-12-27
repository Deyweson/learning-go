package condicionais

import (
	"fmt"
	"math"
)

func condicionais() {

	// Declaração do if junto a declaração de uma var
	if x := math.Sqrt(4); x < 1 {
		fmt.Println(x)
	} else if x > 0 {
		fmt.Println(x)
	} else {
		fmt.Println(x)
	}

}
