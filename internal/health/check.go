package health

import (
	"github.com/dimiro1/health"
	"net/http"
)

type Checker struct {
	serverName string
}

func NewHealthChecker(serverName string) *Checker {
	return &Checker{
		serverName: serverName,
	}
}

// CheckHandlerCustom The controller for state control is rsaEncrypt.
// @Summary Verificar el estado del servicio
// @Description Devuelve el estado del servicio.
// @Tags Health
// @Produce json
// @Success 200
// @Failure 503
// @Router /health [GET]
func (ch *Checker) CheckHandlerCustom() http.Handler {
	handler := health.NewHandler()
	handler.AddInfo("service", ch.serverName)
	handler.AddInfo("endpoint", "health")
	return handler
}
