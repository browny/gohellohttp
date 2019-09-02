package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	r.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// dump request body
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(requestDump))

		// print host name
		name, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "host: %v", name)
	})

	r.GET("/kservice", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		kService := os.Getenv("K_SERVICE")
		fmt.Fprintf(w, "K_SERVICE: %s", kService)
	})

	r.GET("/health", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		isHealth := os.Getenv("HEALTH")
		fmt.Printf("isHealth: %s\n", isHealth)

		if isHealth == "false" {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	log.Println("listening to port *:8080. press ctrl + c to cancel.")
	log.Fatal(http.ListenAndServe(":8080", r))
}
