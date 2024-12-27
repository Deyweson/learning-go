package main

import (
	"fmt"

	"github.com/deyweson/learning-go/alura-api-rest/database"
	"github.com/deyweson/learning-go/alura-api-rest/models"
	"github.com/deyweson/learning-go/alura-api-rest/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "nome 1", Historia: "historia 1"},
		{Id: 2, Nome: "nome 2", Historia: "historia 2"},
	}

	database.ConectaBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
