package app

import (
	"GoJwtCreate/internal/jwt/rsaEncrypt/operations"
	"GoJwtCreate/internal/jwt/rsaEncrypt/platform/handler"
	"GoJwtCreate/kit/config"
	"GoJwtCreate/kit/logger"
	"github.com/gorilla/mux"
	"net/http"
)

func RunEndpointRSAEncrypt(router *mux.Router, config config.IConfig, log logger.ILogger) {
	router.HandleFunc("/jwt/rsaEncrypt/rsa256", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Origins", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Connection", "Keep-Alive")
	}).Methods(http.MethodOptions)

	// Endpoint JWT Encrypt
	jwtCreateOperation := operations.NewJwtCreated(config)
	jwtCreateHandler := handler.NewJwtEncryptHandler(jwtCreateOperation, log)
	router.HandleFunc("/jwt/rsaEncrypt/rsa256", jwtCreateHandler.ServerHTTP).Methods(http.MethodPost)
}
