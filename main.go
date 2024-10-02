package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello World"))

}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Отображение заметки"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allowed", "POST")

		w.WriteHeader(405)
		w.Write([]byte("GET-Метод запрещен"))
		return
	}
	w.Write([]byte("Создание заметки"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	log.Println("Starting web-server http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
