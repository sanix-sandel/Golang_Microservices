package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Go_Microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gb := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gb)

	http.ListenAndServe(":8000", sm)
}
