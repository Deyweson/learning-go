package routes

import (
	"log"
	"net/http"

	"github.com/deyweson/learning-go/alura-api-rest/controllers"
	"github.com/deyweson/learning-go/alura-api-rest/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Função handle para controlar os requests
func HandleRequest() {
	r := mux.NewRouter()

	r.Use(middleware.SetContentTypeMiddlewate)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPesonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriarPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletarPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditarPersonalidade).Methods("Put")

	// Inicia o server e log.Fatal retorna para a gente case de algum erro
	// Define o cors com a lib gorilla/handlers
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
