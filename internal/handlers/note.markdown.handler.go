package handlers

import (
	"markdown-note/internal/services/reply"
	"net/http"
)

func (h *NoteHandler) GetParsed(w http.ResponseWriter, r *http.Request) {
	rp := reply.New(w)
	s := h.s.AttachCtx(r.Context())
	id := r.PathValue("id")

	html, err := s.ParseNoteMarkdown(id)
	if err != nil {
		notFoundNoteOrServerError(rp, err)
		return
	}

	rp.Header().Set("Content-Type", "text/html")
	rp.Success(html).RawReply(http.StatusOK)
}
