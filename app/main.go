package main

import (
	"log"
	"markdown-note/internal/routes"
	"markdown-note/internal/services/database"
	"markdown-note/internal/services/reply"
	"net/http"
)

const (
	PORT    = "3000"         // Server port
	DB_PATH = "data/data.db" // Database path
)

func main() {
	db := database.Connect(DB_PATH)

	routes.RegisterNote(db)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reply.New(w).Error(reply.CodeNotFound, "Path not found").Fail(http.StatusNotFound)
	})

	log.Println("Server listening on port " + PORT)
	http.ListenAndServe(":"+PORT, nil)
}
