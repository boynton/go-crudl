package main

import(
	"log"
	"net/http"

	"example/crudl"
)


func main() {
	controller := NewController()
	endpoint := "localhost:8000"
	handler := crudl.InitServer(controller, "/")
	handler = crudl.WebLog(handler)
	handler = crudl.AllowCors(handler, "http://localhost:8080")
	log.Printf("Listening on http://%s/\n", endpoint)
   log.Fatal(http.ListenAndServe(endpoint, handler))
}
