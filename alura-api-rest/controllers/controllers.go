package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deyweson/learning-go/alura-api-rest/database"
	"github.com/deyweson/learning-go/alura-api-rest/models"
	"github.com/gorilla/mux"
)

// Controlador
// W é o response e R é o request, parecido com req e res do node
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPesonalidades(w http.ResponseWriter, r *http.Request) {
	// Define o tipo do retorno da api
	// w.Header().Set("Content-type", "application/json")

	// Realiza o encode para enviar os dados via json na resposta da api
	// Cria a variavel para receber os valores
	var p []models.Personalidade
	// Buscar no banco de dados com base no endereço de memória da variavel
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var Personalidade models.Personalidade
	database.DB.First(&Personalidade, id)
	json.NewEncoder(w).Encode(Personalidade)
}

func CriarPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Create(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}

func DeletarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var Personalidade models.Personalidade
	database.DB.Delete(&Personalidade, id)
	json.NewEncoder(w).Encode(Personalidade)
}

func EditarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var Personalidade models.Personalidade

	database.DB.First(&Personalidade, id)

	json.NewDecoder(r.Body).Decode(&Personalidade)

	database.DB.Save(&Personalidade)

}
