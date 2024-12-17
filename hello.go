package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const delay = 10 * time.Second
const monitoramentos = 5

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
	sites := []string{}

	sites = append(sites, "https://alura.com.br")
	sites = append(sites, "https://alura.com.br")
	sites = append(sites, "https://alura.com.br")

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testarURL(site)
		}
		time.Sleep(delay)
	}

}

func testarURL(url string) {
	response, _ := http.Get(url)

	if response.StatusCode == 200 {
		fmt.Println("Site:", url, "foi carregado com sucesso")
	} else {
		fmt.Println("Site:", url, "esta com problemas. Status Code:", response.StatusCode)
	}
}
