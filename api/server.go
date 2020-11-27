package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/nitinda/microservice-change-log/api/auto"
	"github.com/nitinda/microservice-change-log/api/router"
	"github.com/nitinda/microservice-change-log/config"
	"github.com/nitinda/microservice-change-log/logger"
)

// Run logger
func Run() {
	// Initial Logger
	logger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	config.LoadEnv()
	auto.LoadData()
	// l := log.New(os.Stdout, "change-log-api ", log.LstdFlags)

	r := router.NewRouter()

	// Create new Server
	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.API_BIND_ADDRESS, config.API_PORT), // Condifgure the bind address
		Handler:      r,                                                              // Set the Default Handler
		ErrorLog:     logger.Trace,                                                   // ErrorLog specifies an optional logger for errors accepting
		ReadTimeout:  5 * time.Second,                                                // ReadTimeout is the maximum duration for reading the entire request
		WriteTimeout: 10 * time.Second,                                               // WriteTimeout is the maximum duration before timing out
		IdleTimeout:  120 * time.Second,                                              // IdleTimeout is the maximum amount of time to wait for the
	}

	// Start the Server
	go func() {
		logger.Info.Println("Starting server on port ", config.API_PORT)

		err := s.ListenAndServe()
		if err != nil {
			logger.Error.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until a signal is received.
	sig := <-sigChan
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	tc, _ := context.WithTimeout(context.Background(), 5*time.Second)
	s.Shutdown(tc)

}
