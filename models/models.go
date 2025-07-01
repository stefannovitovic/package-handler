package models

type Request struct {
    Order     int   `json:"order"`
    PackSizes []int `json:"pack_sizes"`
}

type Response struct {
    TotalItems int         `json:"total_items"`
    Packs      map[int]int `json:"packs"`
}