# Aprendendo GO

## Tipos de variaveis
Variaveis devem ser utilizadas e não podem ser deixadas sem uso, isso gera erro de compilação
como := conseguimos inicialziar e inferir o tipo de uma variavel

- float
- float32
- float64
- int

- string

- nil -> valor nulo

Podemos criar variaveis com valor constante tambem, utilizando a palavra reservada const 
depois de atribuida um valor ele não poderar ser alterado

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

## Loops
For infinito:
for {}

For tradicional:
for i:=0 ;  i < len(array); i++{}

For range:
o range retorna a posição do item no array e o seu valor 
for posição,item := range items{}


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

## Coleção de dados

Array:
Arrays são definidos com seu tamanho e ele não pode ser alterado
sintax: var array [4]string

Slice: 
Simplificando são arrays, porem com habilidae de se adaptarem a quantidade de dados que voce quise rinseir tendo um tamanho dinamico
sintax: array := int[]{1,2,3}
sintax: array := int[]{}


cap(array) -> retorna capacidaed do array
len(array) -> retorna a quantidade de elementos do array

## Lendo arquivos

É possivel ler os arquivos de varias maneiras

uamd elas é utilizanfo
o.Open(filename) -> retorna o endereço em memoria e um erro caso tenha
os.ReadFile(filename) -> retorna os bits do arquivo e um erro caso tenha

para ler as linhas de um arquivo podemos fazer da seguinte forma 
acessamos o arquivo com:
arquivo, erro := os.Open(filename)

e depois usamos um leitor com o pacote bufio:
leitor := bufio.NewReader(arquivo)

com esse leitor podemos ler as linhas do arquivo com o ReadString:
linha, erro := leitor.ReadString('\n')
O ReadString precisa de um bit para indicar ate aonde ele vai ler a linha e ele retorna o texto da linha e um erro caso tenha
Depois de ler a ultima linha ele retorna o erro END OF FILE

## Comando para executr o código
- go run filename.go -> executa o código
- go build filename.go -> gera o bin do código

## Misc.

- fmt -> pacote para formatações (ver mais)
- reflect 
  TypeOf -> verificar tipo da var
- os -> comunica com sistema operacional
  Exit() -> encerra o programa
- time -> manipulaçao de tempo
	Sleep -> faz o programa esperar por n tempo