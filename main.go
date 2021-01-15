package main

import (
	"encoding/json"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type healthCheck struct {
	Status string `json:"status"`
}

func index(ctx *fasthttp.RequestCtx) {
	log.Println("GET /")

	response := healthCheck{
		Status: "ok!",
	}

	ctx.Response.Header.SetContentType("application/json")

	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		log.Fatal("JSON parser error.")
	} else {
		log.Printf("RESPONSE GET / %+v\n", response)
	}
}

func main() {
	route := router.New()

	route.GET("/", index)

	log.Println("Server start at http://localhost:8000")
	log.Fatal(fasthttp.ListenAndServe(":8000", route.Handler))
}
