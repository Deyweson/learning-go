package main

import (
	"fmt"

	"github.com/deyweson/learning-go/poo/clientes"
	contas "github.com/deyweson/learning-go/poo/package"
)

type verificarConta interface {
	Sacar(valor float64) string
}

func boleto(conta verificarConta, valor float64) {

}

func main() {

	// // Exemplo utilizando os nomes do atributos

	// var conta contas.ContaCorrente = contas.ContaCorrente{
	// 	Titular: "Deyweson",
	// 	Agencia: 123,
	// 	Conta:   123,
	// 	Saldo:   1000,
	// }

	// conta.Sacar(10)

	// // Exemplo de criação rapida, sem os nomes do atribuitos
	// conta2 := contas.ContaCorrente{"Deyve", 124, 123, 1001}
	// conta2.Agencia = 1010

	// conta.Transferir(10, &conta2)

	// // Exemplo de ponteiro, aonde podemos preencher apenas os atribuitos necessarios
	// var conta3 *contas.ContaCorrente
	// conta3 = new(contas.ContaCorrente)
	// conta3.Agencia = 123
	// fmt.Println(*conta3)

	// fmt.Println(conta.Agencia, conta2)

	// Criando conta com o exemplo de tipo dentro de tipo
	conta := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Deyweson",
		CPF:       "12345567900",
		Profissao: "Desenvolvedor",
	}, Agencia: 1010, Conta: 1010, Saldo: 1010}

	fmt.Println(conta)
}
