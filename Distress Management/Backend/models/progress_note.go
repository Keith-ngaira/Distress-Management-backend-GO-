package models

import (
	"database/sql"
	"time"
)

type ProgressNote struct {
	ID        int64     `json:"id"`
	CaseID    int64     `json:"case_id"`
	UserID    int64     `json:"user_id"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateProgressNote adds a new progress note to the database
func (p *ProgressNote) Create(db *sql.DB) error {
	query := `
		INSERT INTO progress_notes (case_id, user_id, note, created_at, updated_at)
		VALUES (?, ?, ?, NOW(), NOW())
	`
	result, err := db.Exec(query, p.CaseID, p.UserID, p.Note)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = id
	return nil
}

// GetProgressNotes retrieves all progress notes for a specific case
func GetProgressNotes(db *sql.DB, caseID int64) ([]ProgressNote, error) {
	query := `
		SELECT id, case_id, user_id, note, created_at, updated_at
		FROM progress_notes
		WHERE case_id = ?
		ORDER BY created_at DESC
	`
	rows, err := db.Query(query, caseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []ProgressNote
	for rows.Next() {
		var note ProgressNote
		err := rows.Scan(&note.ID, &note.CaseID, &note.UserID, &note.Note, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}
