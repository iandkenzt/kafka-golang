package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/iandkenzt/kafka-golang/app/isalive"
	"github.com/iandkenzt/kafka-golang/app/producers"

	"github.com/iandkenzt/kafka-golang/utils"

	c "github.com/iandkenzt/kafka-golang/configuration"
)

func init() {
	// instantiate a logger
	utils.Logger.Out = os.Stdout
	utils.InitLogger()

	// instantiate a kafka configuration
	c.InitProducerConfig()
}

// RestAPIRouter ...
func RestAPIRouter() *mux.Router {

	// setup api prefix & version endpoint
	apiPrefix := c.Conf.APIPrefix
	apiVersion := c.Conf.APIVersion

	// register router path prefix
	router := mux.NewRouter()
	restAPIRouter := router.PathPrefix(apiPrefix + apiVersion).Subrouter()

	// registration blueprint route App
	isalive.BuildIsaliveRoutes(restAPIRouter)
	producers.BuildProducersRoutes(restAPIRouter)

	return router
}

func main() {

	// check HTTP port
	port := "3000"
	if c.Conf.Port != "" {
		port = c.Conf.Port
	}
	utils.Logger.Info("Listening on port:", port)

	// router handler
	w := utils.Logger.Writer()
	restAPIRouter := RestAPIRouter()
	loggedRouter := handlers.LoggingHandler(w, restAPIRouter)

	// setup default configuration http server
	addr := fmt.Sprintf(":%s", port)
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	server := &http.Server{
		Addr:         addr,
		Handler:      loggedRouter,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// check and running http server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		utils.Logger.Fatalf("Could not listen on %s: %v", port, err)
	}
}
