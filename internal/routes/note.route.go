package routes

import (
	"markdown-note/internal/handlers"
	"markdown-note/internal/repo"
	"markdown-note/internal/services"
	"net/http"

	"gorm.io/gorm"
)

func RegisterNote(db *gorm.DB) {
	nr := repo.Note(db)
	ns := services.Note(nr)
	nh := handlers.Note(ns)

	http.HandleFunc("GET /notes/{id}", nh.GetOne)
	http.HandleFunc("GET /notes", nh.GetAll)
	http.HandleFunc("POST /notes", nh.CreateOne)
	http.HandleFunc("PUT /notes/{id}", nh.UpdateOne)
	http.HandleFunc("PATCH /notes/{id}/fix-grammar", nh.FixGrammar)
	http.HandleFunc("GET /notes/{id}/html", nh.GetParsed)
	http.HandleFunc("DELETE /notes/{id}", nh.DeleteOne)
}
