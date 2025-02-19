package models

import (
	"database/sql"
	"time"
)

type Case struct {
	ID                   int64     `json:"id"`
	ReferenceNumber      string    `json:"reference_number"`
	SenderName          string    `json:"sender_name"`
	ReceivingDate       time.Time `json:"receiving_date"`
	Subject             string    `json:"subject"`
	CountryOfOrigin     string    `json:"country_of_origin"`
	DistressedPersonName string    `json:"distressed_person_name"`
	NatureOfCase        string    `json:"nature_of_case"`
	CaseDetails         string    `json:"case_details"`
	Status              string    `json:"status"`
	AssignedOfficerID   int64     `json:"assigned_officer_id"`
	Stage               string    `json:"stage"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func (c *Case) Create(db *sql.DB) error {
	query := `INSERT INTO cases 
		(reference_number, sender_name, receiving_date, subject, country_of_origin, 
		distressed_person_name, nature_of_case, case_details, status, stage, 
		created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := db.Exec(query,
		c.ReferenceNumber,
		c.SenderName,
		c.ReceivingDate,
		c.Subject,
		c.CountryOfOrigin,
		c.DistressedPersonName,
		c.NatureOfCase,
		c.CaseDetails,
		c.Status,
		c.Stage,
		c.CreatedAt,
		c.UpdatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	c.ID = id
	return nil
}

func GetCase(db *sql.DB, id int64) (*Case, error) {
	c := &Case{}
	query := `SELECT * FROM cases WHERE id = ?`
	err := db.QueryRow(query, id).Scan(
		&c.ID,
		&c.ReferenceNumber,
		&c.SenderName,
		&c.ReceivingDate,
		&c.Subject,
		&c.CountryOfOrigin,
		&c.DistressedPersonName,
		&c.NatureOfCase,
		&c.CaseDetails,
		&c.Status,
		&c.AssignedOfficerID,
		&c.Stage,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func GetCases(db *sql.DB, page, limit int) ([]Case, error) {
	offset := (page - 1) * limit
	query := `
		SELECT 
			id, reference_number, sender_name, receiving_date, subject,
			country_of_origin, distressed_person_name, nature_of_case,
			case_details, status, assigned_officer_id, stage,
			created_at, updated_at
		FROM cases 
		ORDER BY created_at DESC 
		LIMIT ? OFFSET ?
	`
	
	rows, err := db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cases []Case
	for rows.Next() {
		var c Case
		err := rows.Scan(
			&c.ID,
			&c.ReferenceNumber,
			&c.SenderName,
			&c.ReceivingDate,
			&c.Subject,
			&c.CountryOfOrigin,
			&c.DistressedPersonName,
			&c.NatureOfCase,
			&c.CaseDetails,
			&c.Status,
			&c.AssignedOfficerID,
			&c.Stage,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Convert timestamps to Kenyan time
		kenyaLocation, _ := time.LoadLocation("Africa/Nairobi")
		c.CreatedAt = c.CreatedAt.In(kenyaLocation)
		c.UpdatedAt = c.UpdatedAt.In(kenyaLocation)
		c.ReceivingDate = c.ReceivingDate.In(kenyaLocation)

		cases = append(cases, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cases, nil
}

func (c *Case) Update(db *sql.DB) error {
	query := `UPDATE cases SET 
		sender_name = ?,
		subject = ?,
		country_of_origin = ?,
		distressed_person_name = ?,
		nature_of_case = ?,
		case_details = ?,
		status = ?,
		assigned_officer_id = ?,
		stage = ?,
		updated_at = NOW()
		WHERE id = ?`

	_, err := db.Exec(query,
		c.SenderName,
		c.Subject,
		c.CountryOfOrigin,
		c.DistressedPersonName,
		c.NatureOfCase,
		c.CaseDetails,
		c.Status,
		c.AssignedOfficerID,
		c.Stage,
		c.ID,
	)
	return err
}
