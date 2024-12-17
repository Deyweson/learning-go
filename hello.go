package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
	sites := lerArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testarURL(site)
		}
		time.Sleep(delay)
	}

}

func testarURL(url string) {
	response, erro := http.Get(url)

	if erro != nil {
		fmt.Println("Error:", erro)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", url, "foi carregado com sucesso")
	} else {
		fmt.Println("Site:", url, "esta com problemas. Status Code:", response.StatusCode)
	}
}

func lerArquivo() []string {

	sites := []string{}

	// arquivo, erro := os.Open("sites.txt")
	// arquivo, erro := os.ReadFile("sites.txt")
	arquivo, erro := os.Open("sites.txt")

	if erro != nil {
		fmt.Println("Error: ", erro)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, erro := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if erro == io.EOF {
			break
		}

	}
	arquivo.Close()

	return sites
}
