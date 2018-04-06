package https

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"hello world with http!")
}

func startHttpServer() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
