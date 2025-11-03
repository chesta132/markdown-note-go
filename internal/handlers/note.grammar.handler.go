package handlers

import (
	"fmt"
	"markdown-note/internal/services/grammar"
	"markdown-note/internal/services/reply"
	"net/http"
)

func (h *NoteHandler) FixGrammar(w http.ResponseWriter, r *http.Request) {
	rp := reply.New(w)
	s := h.s.AttachCtx(r.Context())
	id := r.PathValue("id")
	lang := r.URL.Query().Get("lang")

	if !grammar.IsLanguageAllowed(lang) {
		rp.
			Error(reply.CodeBadRequest, fmt.Sprintf("Language '%s' not allowed", lang), fmt.Sprintf("Supported languages: %v", grammar.AllowedLanguages)).
			Fail(http.StatusBadRequest)
		return
	}

	note, err := s.GetById(id)
	if err != nil {
		notFoundNoteOrServerError(rp, err)
		return
	}

	s.FixGrammarAndUpdate(&note, rp, lang, id)

	rp.Success(note).Ok()
}
