package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"distress-management/models"
	"github.com/gorilla/mux"
)

func (app *App) AddProgressNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	caseID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid case ID")
		return
	}

	var note models.ProgressNote
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	note.CaseID = caseID
	if err := note.Create(app.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, note)
}

func (app *App) GetProgressNotes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	caseID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid case ID")
		return
	}

	notes, err := models.GetProgressNotes(app.DB, caseID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, notes)
}
