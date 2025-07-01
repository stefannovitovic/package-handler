package handlers

import (
    "encoding/json"
    "net/http"

	"package-handler/models"
)


func CalculatePacksHandler(findPacks func(int, []int) (int, map[int]int)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req models.Request
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Order <= 0 || len(req.PackSizes) == 0 {
            http.Error(w, "Invalid input", http.StatusBadRequest)
            return
        }

        total, packBreakdown := findPacks(req.Order, req.PackSizes)
        if packBreakdown == nil {
            http.Error(w, "Order could not be fullfiled with following packs", http.StatusUnprocessableEntity)
            return
        }

        resp := models.Response{
            TotalItems: total,
            Packs:      packBreakdown,
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
    }
}