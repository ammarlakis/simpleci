package webhook

import (
	"fmt"
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
		repos, repoProvided := r.URL.Query()["repo"]
		branches, branchProvided := r.URL.Query()["branch"]

		if !repoProvided {
			w.WriteHeader(400)
			w.Write([]byte("Repository url is missing"))
			return
		}

		if !branchProvided {
			w.WriteHeader(400)
			w.Write([]byte("Branch name is missing"))
			return
		}

		repo := repos[0]
		branch := branches[0]

		w.Write([]byte(fmt.Sprintf("got webhook for %s:%s", repo, branch)))
	})
	err := http.ListenAndServe("localhost:8000", server)
	if err != nil {
		serverLogger.Fatal(err)
	}
}
