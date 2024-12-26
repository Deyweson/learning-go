package variaveis

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 65
	f := float64(x)

	// s := string(x)
	s := strconv.FormatInt(int64(x), 10)
	fmt.Println(f, s)
}
