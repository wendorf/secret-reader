package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if err := http.ListenAndServe("0.0.0.0:10000", handler{}); err != nil {
		panic(err)
	}
}

type handler struct{}

func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("file") {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("no files requested"))
		return
	}

	for _, file := range r.URL.Query()["file"] {
		b, err := os.ReadFile(file)
		if err != nil {
			fileNotFound(w, file)
			continue
		}

		_, _ = w.Write([]byte(fmt.Sprintf("%s:\n%s\n\n", file, b)))
	}
}

func fileNotFound(w http.ResponseWriter, file string) {
	_, _ = w.Write([]byte(fmt.Sprintf("%s: not found\n\n", file)))
}
