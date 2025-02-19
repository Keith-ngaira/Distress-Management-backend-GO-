package models

import (
	"database/sql"
	"time"
)

type Document struct {
	ID         int64     `json:"id"`
	CaseID     int64     `json:"case_id"`
	FileName   string    `json:"file_name"`
	FilePath   string    `json:"file_path"`
	FileType   string    `json:"file_type"`
	FileSize   int64     `json:"file_size"`
	UploadedAt time.Time `json:"uploaded_at"`
}

func (d *Document) Create(db *sql.DB) error {
	query := `INSERT INTO documents (case_id, file_name, file_path, file_type, file_size) 
             VALUES (?, ?, ?, ?, ?)`
	
	result, err := db.Exec(query, d.CaseID, d.FileName, d.FilePath, d.FileType, d.FileSize)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	d.ID = id
	return nil
}

func GetDocumentsByCase(db *sql.DB, caseID int64) ([]Document, error) {
	query := `SELECT id, case_id, file_name, file_path, file_type, file_size, uploaded_at 
             FROM documents WHERE case_id = ?`
	
	rows, err := db.Query(query, caseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []Document
	for rows.Next() {
		var doc Document
		err := rows.Scan(
			&doc.ID,
			&doc.CaseID,
			&doc.FileName,
			&doc.FilePath,
			&doc.FileType,
			&doc.FileSize,
			&doc.UploadedAt,
		)
		if err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}

	return documents, nil
}

func (d *Document) Delete(db *sql.DB) error {
	query := `DELETE FROM documents WHERE id = ?`
	_, err := db.Exec(query, d.ID)
	return err
}
