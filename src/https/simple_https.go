package https

import (
	"io"
	"log"
	"net/http"
)

func startHttpsServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello world with https!")
	})
	if e := http.ListenAndServeTLS(":8080", "./server.crt", "./server.key", nil); e != nil {
		log.Fatal("ListenAndServe: ", e)
	}
}
