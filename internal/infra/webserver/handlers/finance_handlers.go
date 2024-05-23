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

type OutputGetFinancesFailure struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// Get user finances
//
//	@Summary		Get finances
//	@Description	Get finances
//	@Tags			finances
//	@Produce		json
//	@Success		200	{object}	[]OutputGetFinances
//	@Failure		500	{object}	OutputGetFinancesFailure
//	@Failure		400	{object}	string
//	@Security		ApiKeyAuth
//	@Router			/finances [get]
func GetFinances(w http.ResponseWriter, r *http.Request) {
	output := []OutputGetFinances{
		{
			Amount: 200.02,
			Date:   "2024-01-01",
			Status: "Approved",
		},
	}
	json.NewEncoder(w).Encode(&output)
}
