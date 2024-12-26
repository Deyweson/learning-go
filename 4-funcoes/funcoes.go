package funcoes

// Função simples
func somar(a int, b int) int {
	return a + b
}

// Função simples omitindo tipo de varias iguais
func somar2(a, b int) int {
	return a + b
}

// Função simples definindo o retorn
func somar3(a, b int) (res int) {
	res = a + b
	return res
}

// Exemplo de função com retorno pelado
func somar4(a, b int) (res int) {
	res = a + b
	return
}

// Função com duplo retorno
func dividir(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return res, rem
}

// Exemplo função HigherOrder / Clojure
func somar5(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

// Função com entrada variaveis
func somar6(nums ...int) (res int) {
	for _, n := range nums {
		res += n
	}
	return res
}
