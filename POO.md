# Go - POO

## Criando um type de objeto 

Criação básica de um tipo/objeto
```go 
type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}
```
Definição de um método
No go antes do nome do método definimos um ponteiro para qual ele pertencera
Assim os objetos do tipo poderam usar esse método
```go
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar {
		c.saldo = c.saldo - valorDoSaque
		return "Pode Sacar"
	} else {
		return "Saldo insuficiente"
	}
}
```


Exemplos de uso:
```go
var conta ContaCorrente = ContaCorrente{
		titular: "Deyweson",
		agencia: 123,
		conta:   123,
		saldo:   1000,
	}

	conta.Sacar(10)
```
No arquivo poo.go tem mais exemplos da criação do objeto


