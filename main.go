package main

import (
	"github.com/shshang/gorilla-mux-cassandra/app"
	"github.com/shshang/gorilla-mux-cassandra/logger"
)

func main() {
	logger.Info("Starting application")
	app.Start()
}
