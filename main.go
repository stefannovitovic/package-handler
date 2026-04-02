package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"package-handler/handlers"
	"package-handler/logic"
)

//go:embed static/index.html
var indexHTML []byte

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write(indexHTML)
	})
	mux.Handle("POST /calculate-packs", handlers.CalculatePacksHandler(logic.FindPacks))

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
