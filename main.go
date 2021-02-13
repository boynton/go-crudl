package main

import(
	"log"
	"net/http"

	"example/crudl"
)


func main() {
	controller := NewController()
	endpoint := "localhost:8080"
	handler := crudl.InitServer(controller, "/")
	handler = crudl.WebLog(handler)
	log.Printf("Listening on http://%s/\n", endpoint)
	http.ListenAndServe(endpoint, handler)

}
