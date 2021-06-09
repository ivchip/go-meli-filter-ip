package http

func StartWebServer(port string) {
	router := newChiRouter()
	router.routesWithMiddleware()
	router.routesWithOutMiddleware()
	router.serve(port)
}
