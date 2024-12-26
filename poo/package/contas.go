package contas

import "github.com/deyweson/learning-go/poo/clientes"

type ContaCorrente struct {
	Titular clientes.Titular
	Agencia int
	Conta   int
	Saldo   float64
}

// Método do type ContaCorrente
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.Saldo && valorDoSaque > 0
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Pode Sacar"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valor float64) {
	if valor > 0 {
		c.Saldo += valor
	}
}

func (c *ContaCorrente) Transferir(valor float64, conta *ContaCorrente) {
	if valor > 0 && c.Saldo >= valor {
		c.Saldo -= valor
		conta.Saldo += valor
	}
}
