package handler

import (
	_ "embed"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yehezkiel1086/devops-with-kubernetes/todo-app/internal/adapter/config"
)

//go:embed index.html
var indexHTML []byte

type Router struct {
	r *chi.Mux
}

func NewRouter() *Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", handleHome)

	return &Router{r}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(indexHTML)
}

func (r *Router) Mux() *chi.Mux {
	return r.r
}

func (r *Router) Run(conf *config.HTTP) error {
	uri := conf.Host + ":" + conf.Port

	fmt.Printf("Server started in port %s\n", conf.Port)
	return http.ListenAndServe(uri, r.r)
}
