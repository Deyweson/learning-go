package variaveis

import "fmt"

func variaveis() {

	// Tipos Básicos
	// inteiros com sinal: int int8 int 32 int64
	// inteiros sem sinal: uint uint8 uint32 uint64 uintptr
	//
	// byte
	// rune
	//
	// Numeros flutuantes: float32 float64
	//
	// Numeros complexos: complex64 complex128

	// Declaração de varivael
	var foo string = "foo"
	// Declaração curta de variavel
	bar := "bar"

	// Declarando duas variaveis de forma unica
	var nome, sobrenome string = "Deyweson", "Almeida"

	// Declarando variavesi sem indicar o tipo
	// O tipo será auto inferido
	var idade, altura = 24, 1.75

	// Declaração de duas variaveis
	// sem indicar tipo na forma curta
	quantidade, valor := 26, 14.40

	// Variaveis deve ser utilizadas
	// Em go não é permitido criar e não utilizar varivaveis
	fmt.Println(foo, bar)
	fmt.Println(nome, sobrenome)
	fmt.Println(idade, altura)
	fmt.Println(quantidade, valor)

	// Variaveis Constante
	const teste int = 10
	// Omitindo o tipo
	// A const vai adptar o tipo ao contexto
	// Nem sempre vai servir por exemplo:
	// um float nao vai ser possivel usar em int
	const teste2 = 20
	// Além disso as contanste nao é possivel usar a declaração curta
}
