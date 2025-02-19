package handlers

import (
	"distress-management/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const (
	uploadDir = "./uploads"
	maxFileSize = 10 << 20 // 10 MB
)

func init() {
	// Create uploads directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		panic(err)
	}
}

func (app *App) UploadDocument(w http.ResponseWriter, r *http.Request) {
	// Parse case ID from URL
	vars := mux.Vars(r)
	caseID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid case ID")
		return
	}

	// Verify case exists
	_, err = models.GetCase(app.DB, caseID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Case not found")
		return
	}

	// Limit file size
	r.Body = http.MaxBytesReader(w, r.Body, maxFileSize)
	if err := r.ParseMultipartForm(maxFileSize); err != nil {
		respondWithError(w, http.StatusBadRequest, "File too large (max 10MB)")
		return
	}

	// Get file from form
	file, header, err := r.FormFile("document")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error retrieving file")
		return
	}
	defer file.Close()

	// Validate file type
	fileType := header.Header.Get("Content-Type")
	if !isAllowedFileType(fileType) {
		respondWithError(w, http.StatusBadRequest, "File type not allowed")
		return
	}

	// Create unique filename
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("case_%d_%d%s", caseID, time.Now().Unix(), ext)
	filePath := filepath.Join(uploadDir, filename)

	// Create file on disk
	dst, err := os.Create(filePath)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error saving file")
		return
	}
	defer dst.Close()

	// Copy file contents
	if _, err := io.Copy(dst, file); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error saving file")
		return
	}

	// Create document record
	doc := &models.Document{
		CaseID:   caseID,
		FileName: header.Filename,
		FilePath: filePath,
		FileType: fileType,
		FileSize: header.Size,
	}

	if err := doc.Create(app.DB); err != nil {
		// Clean up file if database insert fails
		os.Remove(filePath)
		respondWithError(w, http.StatusInternalServerError, "Error saving document record")
		return
	}

	respondWithJSON(w, http.StatusCreated, doc)
}

func (app *App) GetDocuments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	caseID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid case ID")
		return
	}

	documents, err := models.GetDocumentsByCase(app.DB, caseID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error retrieving documents")
		return
	}

	respondWithJSON(w, http.StatusOK, documents)
}

func (app *App) DeleteDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	docID, err := strconv.ParseInt(vars["docId"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid document ID")
		return
	}

	// Get document to get file path
	doc := &models.Document{ID: docID}
	if err := doc.Delete(app.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error deleting document")
		return
	}

	// Delete file from disk
	if err := os.Remove(doc.FilePath); err != nil {
		// Log error but don't fail the request
		fmt.Printf("Error deleting file %s: %v\n", doc.FilePath, err)
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Document deleted"})
}

func isAllowedFileType(fileType string) bool {
	allowedTypes := map[string]bool{
		"application/pdf":                true,
		"application/msword":            true,
		"application/vnd.ms-excel":      true,
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":       true,
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
	}
	return allowedTypes[fileType]
}
