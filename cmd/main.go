package main

import (
	"context"
	"go_todo/config"
	"go_todo/pkg/router"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
)

const (
	configPath = "config/reader"
)

func main() {
	config.GetConfig(configPath)

	routeHandler := new(router.RouteHandler)
	routeHandler.Validate = validator.New()

	//adding route groups to the routes.
	ginRouter := router.GetRoutingEngine("/v1", routeHandler)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(config.Config.RestPort),
		Handler: ginRouter,
	}

	go func() {
		// server connections
		if err := server.ListenAndServe(); err != nil {
			log.WithFields(log.Fields{"error": err.Error()}).Fatalf("server listen failed")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.WithFields(log.Fields{"method": "main"}).Infoln("server shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Fatalf("server shutdown failed")
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.WithFields(log.Fields{"method": "main"}).Infoln("server exited")
}
