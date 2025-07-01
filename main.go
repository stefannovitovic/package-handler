package main

import (
	"fmt"
	"log"
	"net/http"

	"package-handler/handlers"
	"package-handler/logic"
)
func main() {
    http.HandleFunc("/calculate-packs", handlers.CalculatePacksHandler(logic.FindPacks))
    fmt.Println("Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
