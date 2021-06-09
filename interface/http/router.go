package http

//Router is the interface to be implemented by the HTTP routers
type Router interface {
	routesWithMiddleware()
	routesWithOutMiddleware()
	serve(port string)
}
