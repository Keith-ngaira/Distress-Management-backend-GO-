package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetCases returns a list of all cases
func (app *App) GetCases(w http.ResponseWriter, r *http.Request) {
	page := 1
	limit := 10

	// Parse page and limit from query parameters
	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	fmt.Printf("Fetching cases with page: %d, limit: %d\n", page, limit)
	rows, err := app.DB.Query(`
		SELECT id, reference_number, sender_name, receiving_date, subject, 
		country_of_origin, distressed_person_name, nature_of_case, case_details, 
		status, stage, created_at, updated_at
		FROM cases
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`, limit, (page-1)*limit)
	if err != nil {
		fmt.Printf("Error fetching cases: %v\n", err)
		http.Error(w, "Error retrieving cases: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cases []map[string]interface{}
	for rows.Next() {
		var c struct {
			ID                   int64   `json:"id"`
			ReferenceNumber     string  `json:"referenceNumber"`
			SenderName          string  `json:"senderName"`
			ReceivingDate       string  `json:"receivingDate"`
			Subject             string  `json:"subject"`
			CountryOfOrigin     string  `json:"countryOfOrigin"`
			DistressedPersonName string `json:"distressedPersonName"`
			NatureOfCase        string  `json:"natureOfCase"`
			CaseDetails         string  `json:"caseDetails"`
			Status              string  `json:"status"`
			Stage               string  `json:"stage"`
			CreatedAt           string  `json:"createdAt"`
			UpdatedAt           string  `json:"updatedAt"`
		}

		err := rows.Scan(
			&c.ID, &c.ReferenceNumber, &c.SenderName, &c.ReceivingDate,
			&c.Subject, &c.CountryOfOrigin, &c.DistressedPersonName,
			&c.NatureOfCase, &c.CaseDetails, &c.Status, &c.Stage,
			&c.CreatedAt, &c.UpdatedAt,
		)
		if err != nil {
			http.Error(w, "Error scanning case: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert struct to map
		b, _ := json.Marshal(c)
		var m map[string]interface{}
		json.Unmarshal(b, &m)
		cases = append(cases, m)
	}

	if cases == nil {
		cases = []map[string]interface{}{} // Return empty array instead of null
	}

	fmt.Printf("Successfully fetched %d cases\n", len(cases))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cases)
}

// GetCase returns a single case by ID
func (app *App) GetCase(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid case ID", http.StatusBadRequest)
		return
	}

	var c struct {
		ID                   int64   `json:"id"`
		ReferenceNumber     string  `json:"referenceNumber"`
		SenderName          string  `json:"senderName"`
		ReceivingDate       string  `json:"receivingDate"`
		Subject             string  `json:"subject"`
		CountryOfOrigin     string  `json:"countryOfOrigin"`
		DistressedPersonName string `json:"distressedPersonName"`
		NatureOfCase        string  `json:"natureOfCase"`
		CaseDetails         string  `json:"caseDetails"`
		Status              string  `json:"status"`
		Stage               string  `json:"stage"`
		CreatedAt           string  `json:"createdAt"`
		UpdatedAt           string  `json:"updatedAt"`
	}

	err = app.DB.QueryRow(`
		SELECT id, reference_number, sender_name, receiving_date, subject,
		country_of_origin, distressed_person_name, nature_of_case, case_details,
		status, stage, created_at, updated_at
		FROM cases WHERE id = ?
	`, id).Scan(
		&c.ID, &c.ReferenceNumber, &c.SenderName, &c.ReceivingDate,
		&c.Subject, &c.CountryOfOrigin, &c.DistressedPersonName,
		&c.NatureOfCase, &c.CaseDetails, &c.Status, &c.Stage,
		&c.CreatedAt, &c.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Case not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Convert struct to map
	b, _ := json.Marshal(c)
	var m map[string]interface{}
	json.Unmarshal(b, &m)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

// CreateCase creates a new case
func (app *App) CreateCase(w http.ResponseWriter, r *http.Request) {
	var input struct {
		SenderName           string `json:"senderName"`
		Subject             string `json:"subject"`
		CountryOfOrigin     string `json:"countryOfOrigin"`
		DistressedPersonName string `json:"distressedPersonName"`
		NatureOfCase        string `json:"natureOfCase"`
		CaseDetails         string `json:"caseDetails"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Generate reference number (you might want to make this more sophisticated)
	var lastID int
	err := app.DB.QueryRow("SELECT COALESCE(MAX(id), 0) FROM cases").Scan(&lastID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	referenceNumber := fmt.Sprintf("REF%05d", lastID+1)

	result, err := app.DB.Exec(`
		INSERT INTO cases (
			reference_number, sender_name, receiving_date, subject,
			country_of_origin, distressed_person_name, nature_of_case,
			case_details, status, stage
		) VALUES (?, ?, NOW(), ?, ?, ?, ?, ?, 'Pending', 'Front Office Receipt')
	`,
		referenceNumber, input.SenderName, input.Subject,
		input.CountryOfOrigin, input.DistressedPersonName,
		input.NatureOfCase, input.CaseDetails,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":              id,
		"referenceNumber": referenceNumber,
		"message":         "Case created successfully",
	})
}

// UpdateCase updates an existing case
func (app *App) UpdateCase(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid case ID", http.StatusBadRequest)
		return
	}

	var input struct {
		SenderName           string `json:"senderName"`
		Subject             string `json:"subject"`
		CountryOfOrigin     string `json:"countryOfOrigin"`
		DistressedPersonName string `json:"distressedPersonName"`
		NatureOfCase        string `json:"natureOfCase"`
		CaseDetails         string `json:"caseDetails"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	_, err = app.DB.Exec(`
		UPDATE cases
		SET sender_name = ?, subject = ?, country_of_origin = ?,
			distressed_person_name = ?, nature_of_case = ?, case_details = ?,
			updated_at = NOW()
		WHERE id = ?
	`,
		input.SenderName, input.Subject, input.CountryOfOrigin,
		input.DistressedPersonName, input.NatureOfCase, input.CaseDetails,
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Case updated successfully",
	})
}

// UpdateCaseStatus updates the status and stage of a case
func (app *App) UpdateCaseStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid case ID", http.StatusBadRequest)
		return
	}

	var input struct {
		Status string `json:"status"`
		Stage  string `json:"stage"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	_, err = app.DB.Exec(`
		UPDATE cases
		SET status = ?, stage = ?, updated_at = NOW()
		WHERE id = ?
	`, input.Status, input.Stage, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Case status updated successfully",
	})
}
