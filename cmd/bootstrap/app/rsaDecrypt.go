package app

import (
	"GoJwtCreate/internal/jwt/rsaDecrypt/operations"
	"GoJwtCreate/internal/jwt/rsaDecrypt/platform/handler"
	"GoJwtCreate/kit/config"
	"github.com/gorilla/mux"
	"net/http"
)

func RunEndpointRSADecrypt(router *mux.Router, config config.IConfig) {
	router.HandleFunc("/jwt/rsaDecrypt/rsa256", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Origins", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Connection", "Keep-Alive")
	}).Methods(http.MethodOptions)

	// Endpoint JWT Encrypt
	jwtCreateOperation := operations.NewJwtCreated(config)
	jwtCreateHandler := handler.NewJwtDecryptHandler(jwtCreateOperation)
	router.HandleFunc("/jwt/rsaDecrypt/rsa256", jwtCreateHandler.ServerHTTP).Methods(http.MethodPost)
}
