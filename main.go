package main

import (
	"fmt"
	"net/http"
	"os"
)

func getPort() string {
	p := os.Getenv("APP_PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		hostname, _ := os.Hostname()
		env := os.Getenv("APP_ENV")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Pod name is %s, in %s", hostname, env)))
	})

	http.ListenAndServe(getPort(), nil)
}
