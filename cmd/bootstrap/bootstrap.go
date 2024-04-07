package bootstrap

import (
	"GoJwtCreate/cmd/bootstrap/app"
	"GoJwtCreate/internal/health"
	"GoJwtCreate/kit/config"
	"GoJwtCreate/kit/logger"
	"github.com/gorilla/mux"
	"net/http"
)

const TitleBootstrap = "---- BOOTSTRAP ----"

func Run() {
	// Inicializar configuraciones y logger de servidor
	newConfig := config.NewConfig("application.yaml")
	newLogger := logger.NewLogger()

	serverName, found := newConfig.GetString("server.name")
	if !found {
		newLogger.Fatal(TitleBootstrap, "error", "server name not found")
	}

	port, found := newConfig.GetString("server.port")
	if !found {
		newLogger.Fatal(TitleBootstrap, "error", "server port not found")
	}

	// Crear instancia de Router Mux
	router := mux.NewRouter()

	// Endpoint Health Check
	router.Handle("/health", health.NewHealthChecker(serverName).CheckHandlerCustom()).Methods(http.MethodGet)

	// Endpoint CORS JWT Encrypt RSA
	app.RunEndpointRSAEncrypt(router, newConfig, newLogger)

	// Endpoint CORS JWT Decrypt RSA
	app.RunEndpointRSADecrypt(router, newConfig)

	// Funcion de encedido de servidor
	ServerTurnOn(router, serverName, port, newLogger)
}
