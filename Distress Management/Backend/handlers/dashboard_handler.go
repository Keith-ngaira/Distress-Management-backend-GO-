package handlers

import (
	"net/http"
)

type DashboardStats struct {
	TotalCases            int            `json:"totalCases"`
	CasesByStatus        map[string]int `json:"casesByStatus"`
	CasesByNature        map[string]int `json:"casesByNature"`
	RecentCases          []RecentCase   `json:"recentCases"`
	CasesByCountryOrigin map[string]int `json:"casesByCountryOrigin"`
}

type RecentCase struct {
	ID              int64  `json:"id"`
	ReferenceNumber string `json:"referenceNumber"`
	Subject         string `json:"subject"`
	Status          string `json:"status"`
	Nature          string `json:"natureOfCase"`
}

func (app *App) GetDashboardStats(w http.ResponseWriter, r *http.Request) {
	stats := DashboardStats{
		CasesByStatus:        make(map[string]int),
		CasesByNature:        make(map[string]int),
		CasesByCountryOrigin: make(map[string]int),
	}

	// Get total cases
	err := app.DB.QueryRow("SELECT COUNT(*) FROM cases").Scan(&stats.TotalCases)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error getting total cases: "+err.Error())
		return
	}

	// Get cases by status
	rows, err := app.DB.Query(`
		SELECT status, COUNT(*) as count 
		FROM cases 
		GROUP BY status
	`)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error getting cases by status: "+err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error scanning status counts: "+err.Error())
			return
		}
		stats.CasesByStatus[status] = count
	}

	// Get cases by nature
	rows, err = app.DB.Query(`
		SELECT nature_of_case, COUNT(*) as count 
		FROM cases 
		GROUP BY nature_of_case
	`)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error getting cases by nature: "+err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var nature string
		var count int
		if err := rows.Scan(&nature, &count); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error scanning nature counts: "+err.Error())
			return
		}
		stats.CasesByNature[nature] = count
	}

	// Get cases by country of origin
	rows, err = app.DB.Query(`
		SELECT country_of_origin, COUNT(*) as count 
		FROM cases 
		GROUP BY country_of_origin
	`)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error getting cases by country: "+err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var country string
		var count int
		if err := rows.Scan(&country, &count); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error scanning country counts: "+err.Error())
			return
		}
		stats.CasesByCountryOrigin[country] = count
	}

	// Get recent cases
	rows, err = app.DB.Query(`
		SELECT id, reference_number, subject, status, nature_of_case 
		FROM cases 
		ORDER BY created_at DESC 
		LIMIT 5
	`)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error getting recent cases: "+err.Error())
		return
	}
	defer rows.Close()

	var recentCases []RecentCase
	for rows.Next() {
		var c RecentCase
		if err := rows.Scan(&c.ID, &c.ReferenceNumber, &c.Subject, &c.Status, &c.Nature); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error scanning recent case: "+err.Error())
			return
		}
		recentCases = append(recentCases, c)
	}
	stats.RecentCases = recentCases

	respondWithJSON(w, http.StatusOK, stats)
}
