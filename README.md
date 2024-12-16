# Aprendendo GO

## Tipos de variaveis
Variaveis devem ser utilizadas e não podem ser deixadas sem uso, isso gera erro de compilação
como := conseguimos inicialziar e inferir o tipo de uma variavel

- float
- float32
- float64
- int

- string

## Controle de fluxo
if não precisa de parenteses em volta da condição
A condição do if precisa ser uma booleana true / false
if:
if condição {
  code
}else if condição{
  code
}else {
  code
}


## Funções

func funcao() tipo {
	return tipo
}

função comduplo valor de returno
na hora de receber caso não deseja usar um dos tipos só colcoar o _ para ignorar ele
func funcao() (tipo1, tipo2) {
	return tipo1, tipo2
}

switch:
O switch nao precisa utilizar o break é opcional
switch option {
	case 1:
		code
	case 2:
		code
	case 0:
		code
	default:
		code
	}

## Comando para executr o código
- go run filename.go -> executa o código
- go build filename.go -> gera o bin do código

## Misc.

- fmt -> pacote para formatações (ver mais)
- reflect 
  TypeOf -> verificar tipo da var
- os -> comunica com sistema operacional
  Exit() -> encerra o programa