package main

import (
	"log/slog"
	"os"

	"github.com/yehezkiel1086/devops-with-kubernetes/todo-app/internal/adapter/config"
	"github.com/yehezkiel1086/devops-with-kubernetes/todo-app/internal/adapter/handler"
)

func handleError(msg string, err error) {
	if err != nil {
		slog.Error(msg, "error", err)
		os.Exit(1)
	}
}

func main() {
	// load .env configs
	conf, err := config.New()
	handleError("failed to load .env configs", err)
	slog.Info(".env configs loaded successfully")

	// init router
	r := handler.NewRouter()

	// run api
	err = r.Run(conf.HTTP)
	handleError("failed to run backend api", err)
}
