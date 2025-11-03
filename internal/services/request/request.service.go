package request

import (
	"encoding/json"
	"io"
	"markdown-note/internal/models/note"
	"markdown-note/internal/services/reply"
	"net/http"
	"strings"
)

func Body[T any](w *http.Request) *T {
	var b *T
	json.NewDecoder(w.Body).Decode(b)
	return b
}

type takeFileError struct {
	Message string
	Code    string
	Status  int
}

func TakeFile(r *http.Request) (*note.Note, *takeFileError) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, &takeFileError{"Please request a form as body", reply.CodeBadRequest, http.StatusBadRequest}
	}

	title := r.FormValue("title")
	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, &takeFileError{"Please add file field in form", reply.CodeBadRequest, http.StatusBadRequest}
	}
	defer file.Close()

	if !strings.HasSuffix(header.Filename, ".md") {
		return nil, &takeFileError{"Please send a .md extension as uploaded file in form", reply.CodeBadRequest, http.StatusBadRequest}
	}

	if title == "" {
		title = header.Filename
	}

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, &takeFileError{"Error reading file", reply.CodeServerError, http.StatusInternalServerError}
	}

	if len(content) > 5<<20 {
		return nil, &takeFileError{"File too large (max 5MB)", reply.CodeBadRequest, http.StatusBadRequest}
	}

	return &note.Note{Title: title, Markdown: string(content)}, nil
}
