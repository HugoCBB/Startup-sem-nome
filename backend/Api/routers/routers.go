package routers

import (
	"backend/api/api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Middleware para permitir CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*
			Função que permite o CORS (Cross-Origin Resource Sharing) para que a aplicação frontend possa acessar a API.
			Para isso, é necessário configurar os headers da resposta HTTP para permitir o acesso de origens diferentes.
		*/
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Configuração das rotas e inicialização do servidor
func HandleRequest() {
	r := mux.NewRouter()

	r.Use(enableCORS)

	r.HandleFunc("/", controllers.GetAllUser).Methods("GET")

	log.Println("Servidor rodando na porta 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
