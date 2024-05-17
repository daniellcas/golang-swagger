package handlers

import (
	"encoding/json"
	"net/http"
)

type OutputGetFinances struct {
	Amount float64 `json:"amount"`
	Date   string  `json:"date"`
	Status string  `json:"status"`
}

func GetFinances(w http.ResponseWriter, r *http.Request) {
	output := OutputGetFinances{
		Amount: 200.02,
		Date:   "2024-01-01",
		Status: "Approved",
	}
	json.NewEncoder(w).Encode(&output)
}
