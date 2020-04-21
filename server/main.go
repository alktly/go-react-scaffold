package main

import (
	"fmt"
	"net/http"
)

const (
	host      = "http://localhost"
	port      = ":8080"
	staticDir = "../client/build"
)

func main() {

	mux := http.NewServeMux()

	for _, route := range getRoutes() {
		mux.HandleFunc(
			route.path,
			methodHandler(
				route.method,
				route.handler,
			),
		)
	}

	fmt.Printf("Listening on %v%v\n", host, port)

	if err := http.ListenAndServe(port, mux); err != nil {
		panic(err)
	}
}

func getRoutes() []route {
	return []route{
		{
			method:  http.MethodGet,
			path:    "/",
			handler: http.FileServer(http.Dir(staticDir)).ServeHTTP,
		},
	}
}
