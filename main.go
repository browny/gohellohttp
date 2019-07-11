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
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))
		name, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "host: %v", name)
	})
	log.Println("listening to port *:8080. press ctrl + c to cancel.")
	log.Fatal(http.ListenAndServe(":8080", r))
}
