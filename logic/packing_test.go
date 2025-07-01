package logic

import (
    "reflect"
    "testing"
)

func TestFindPacks(t *testing.T) {
    tests := []struct {
        name      string
        order     int
        packSizes []int
        wantTotal int
        wantPacks map[int]int
    }{
        {
            name:      "Exact match with one pack",
            order:     5000,
            packSizes: []int{250, 500, 1000, 2000, 5000},
            wantTotal: 5000,
            wantPacks: map[int]int{5000: 1},
        },
        {
            name:      "Needs combination, minimal overage",
            order:     4200,
            packSizes: []int{250, 500, 1000, 2000, 5000},
            wantTotal: 4250,
            wantPacks: map[int]int{2000: 2, 250: 1},
        },
        {
            name:      "Impossible to fulfill",
            order:     1,
            packSizes: []int{250, 500},
            wantTotal: 250,
            wantPacks: map[int]int{250: 1},
        },
		{
			name:      "Large order not usuall pack sizes",
			order:     500000,
			packSizes: []int{23, 31, 53},
			wantTotal: 500000,
			wantPacks: map[int]int{53: 9429, 31: 7, 23: 2},
		},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotTotal, gotPacks := FindPacks(tt.order, tt.packSizes)
            if gotTotal != tt.wantTotal {
                t.Errorf("FindPacks() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
            }
            if !reflect.DeepEqual(gotPacks, tt.wantPacks) {
                t.Errorf("FindPacks() gotPacks = %v, want %v", gotPacks, tt.wantPacks)
            }
        })
    }
}