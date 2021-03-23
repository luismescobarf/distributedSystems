package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	var cantidadRequests int64 = 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Saludos, se ha hecho el request para: %s\n", r.URL.Path)
		time.Sleep(3 * time.Second)
		fmt.Fprintf(w, "Sisas %s\n", r.URL.Path)
		cantidadRequests++
		fmt.Fprintf(w, "<b>Número de requests %d </b>", cantidadRequests)

	})

	http.ListenAndServe(":80", nil)

}
