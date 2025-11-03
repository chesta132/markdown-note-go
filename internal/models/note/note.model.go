package note

import (
	gorm_model "markdown-note/internal/models/gorm"
)

type Note struct {
	gorm_model.Model
	Title    string `json:"title"`
	Markdown string `json:"note"`
}
