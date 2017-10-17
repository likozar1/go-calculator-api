package main

import (
	"flag"
	"github.com/go-calculator-api/pkg/router"
	"log"
	"net/http"
)

var Host string

var RouteCollection = router.Routes{
	router.Route{
		"Calculate",
		"GET",
		"/",
		Calculate,
	},

	router.Route{
		"Calculator",
		"POST",
		"/calculate",
		Calculate,
	},

	router.Route{
		"Index",
		"GET",
		"/index",
		Index,
	},
}

func init() {
	flag.StringVar(&Host, "host", "0.0.0.0:8079", "Api host'")
}

func main() {
	router := router.NewRouter(RouteCollection)
	log.Fatal(http.ListenAndServe(Host, router))
}
