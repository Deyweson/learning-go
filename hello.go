package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for {
		menu()

		option := selectOption()

		switch option {
		case 1:
			fmt.Println("Monitorando...")
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido")
			os.Exit(-1)
		}
	}
}

func menu() {
	fmt.Println("1 - Monitorar")
	fmt.Println("2 - Logs")
	fmt.Println("0 - Sair	")
}

func selectOption() int {
	var option int
	fmt.Scan(&option)
	return option
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando o Monitoramento...")
	site := "https://alura.com.br"
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", response.StatusCode)
	}
}
