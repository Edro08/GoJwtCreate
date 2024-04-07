package bootstrap

import (
	"GoJwtCreate/kit/logger"
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func ServerTurnOn(router *mux.Router, serverName string, port string, log logger.ILogger) {

	// Canal para captar señales del sistema operativo y servidor de errores.
	signals := make(chan os.Signal)
	errors := make(chan error)

	// Imprimir routes de mux
	printRoutes(router)

	// Configuracion del servidor
	server := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: 300 * time.Second,
		ReadTimeout:  300 * time.Second,
	}

	// logger de inicio de server
	log.Info(TitleBootstrap, "Server Name", serverName, "listening on port:", port)

	// Iniciando HTTP server
	go func() {
		errors <- server.ListenAndServe()
	}()

	// Registrar señales del sistema operativo para capturarlas
	signal.Notify(signals, syscall.SIGINT)
	signal.Notify(signals, syscall.SIGTERM)
	signal.Notify(signals, os.Interrupt)

	select {
	case <-signals:
		// if this line is reached, a signal has been caught. A logger could
		// be added to indicate the signal type.
		break
	case <-errors:
		// if this line is reached, an error has occurred. A logger could
		// be added to determine the error details.
		break
	}

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "el tiempo durante el cual el servidor espera con gracia a que finalicen las conexiones existentes, 15s o 1 m")
	flag.Parse()

	// crea un deadline a esperar.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// No bloquea si no hay conexiones, pero en caso contrario esperará hasta que se complete el tiempo de espera.
	// Opcionalmente, puedes ejecutar server.Shutdown en una rutina goroutine y bloquear
	// <-ctx.Done() si tu aplicación debe esperar a otros servicios
	// para finalizar según la cancelación del contexto.
	_ = server.Shutdown(ctx)

	log.Info(TitleBootstrap, "status", "shutting down")

	// Cerrar el archivo Json Logs
	log.Close()

	// Terminar
	os.Exit(0)
}

// Funcion que imprime en consola todas las rutas disponibles en el Router Mux
func printRoutes(router *mux.Router) {
	_ = router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, errPath := route.GetPathTemplate()
		methods, errMethod := route.GetMethods()
		if errPath == nil && errMethod == nil {
			fmt.Println(strings.Join(methods, ","), pathTemplate)
		}
		return nil
	})
}
