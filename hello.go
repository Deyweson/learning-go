package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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
			imprimirLogs()
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
		registraLog(url, true)
	} else {
		fmt.Println("Site:", url, "esta com problemas. Status Code:", response.StatusCode)
		registraLog(url, false)
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

func registraLog(url string, status bool) {
	// caminhodo arquivo , flag do que pode se fazer com ele de permissão
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + url + " - online: " + strconv.FormatBool(status) + "\n")
}

func imprimirLogs() {
	file, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(file))

}
