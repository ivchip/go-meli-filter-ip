package http

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	middleware2 "github.com/ivchip/go-meli-filter-ip/interface/http/middleware"
	"log"
	"net/http"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

//NewChiRouter is a Chi HTTP router constructor
func newChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) routesWithMiddleware() {
	chiDispatcher.Group(func(r chi.Router) {
		middleware := middleware2.InitMiddleware()
		r.Use(middleware.LimitRate)
		for _, route := range routesPrivate {
			r.Method(route.Method, route.Pattern, route.HandlerFunc)
		}
	})
}

func (r *chiRouter) routesWithOutMiddleware() {
	chiDispatcher.Group(func(r chi.Router) {
		for _, route := range routesPublic {
			r.Method(route.Method, route.Pattern, route.HandlerFunc)
		}
	})
}

func (*chiRouter) serve(port string) {
	fmt.Printf("Chi HTTP Server running on port: %v", port)
	err := http.ListenAndServe(":"+port, chiDispatcher)
	if err != nil {
		log.Println("An error occurred starting HTTP server at port " + port)
		log.Fatal("Error: " + err.Error())
	}
}
