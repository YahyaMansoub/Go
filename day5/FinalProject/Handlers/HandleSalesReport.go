package Handlers

import (
	"FinalProject/functions"
	"encoding/json"
	"net/http"
	"time"
)

// HandleListSalesReports handles the API endpoint for listing sales reports within a specified date range
func HandleListSalesReports(w http.ResponseWriter, r *http.Request) {
	startDateStr := r.URL.Query().Get("start")
	endDateStr := r.URL.Query().Get("end")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		http.Error(w, "Invalid start date format. Use YYYY-MM-DD.", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		http.Error(w, "Invalid end date format. Use YYYY-MM-DD.", http.StatusBadRequest)
		return
	}

	reports, err := functions.GetSalesReportsInRange(startDate, endDate)
	if err != nil {
		http.Error(w, "Error fetching sales reports.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}
