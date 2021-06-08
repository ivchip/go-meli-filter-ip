package http

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

//NewChiRouter is a Chi HTTP router constructor
func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) ROUTES() {
	for _, route := range routes {
		chiDispatcher.Method(route.Method, route.Pattern, route.HandlerFunc)
	}
}

func (r *chiRouter) MIDDLEWARES() {
	middleware := InitMiddleware()
	chiDispatcher.Use(middleware.LimitRate)
}

func (*chiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP Server running on port: %v", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("An error occurred starting HTTP server at port " + port)
		log.Fatal("Error: " + err.Error())
	}
}
