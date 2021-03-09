package webhook

import (
	"log"
	"net/http"
	"os"
)

/*
StartServer serves webhook requests
*/
func StartServer() {
	serverLogger := log.New(os.Stdout, "http", -1)
	server := http.NewServeMux()
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	err := http.ListenAndServe("localhost:8000", server)
	if err != nil {
		serverLogger.Fatal(err)
	}
}
