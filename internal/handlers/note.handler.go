package handlers

import (
	"encoding/json"
	"errors"
	"markdown-note/internal/models/note"
	"markdown-note/internal/services"
	"markdown-note/internal/services/reply"
	"markdown-note/internal/services/request"
	"net/http"

	"gorm.io/gorm"
)

type NoteHandler struct {
	s *services.NoteService
}

func Note(service *services.NoteService) *NoteHandler {
	return &NoteHandler{service}
}

func notFoundNoteOrServerError(rp *reply.Reply, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		rp.Error(reply.CodeNotFound, "Note not found").Fail(http.StatusNotFound)
		return
	}
	if err != nil {
		rp.Error(reply.CodeServerError, "Server Error", err.Error()).Fail(http.StatusInternalServerError)
		return
	}
}

func (h *NoteHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	rp := reply.New(w)
	id := r.PathValue("id")
	s := h.s.AttachCtx(r.Context())

	note, err := s.GetById(id)
	if err != nil {
		notFoundNoteOrServerError(rp, err)
		return
	}

	rp.Success(note).Ok()
}

func (h *NoteHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	rp := reply.New(w)
	s := h.s.AttachCtx(r.Context())

	notes, err := s.GetAll()
	if err != nil {
		rp.Error(reply.CodeServerError, "Server Error", err.Error()).Fail(http.StatusInternalServerError)
		return
	}

	rp.Success(notes).Ok()
}

func (h *NoteHandler) CreateOne(w http.ResponseWriter, r *http.Request) {
	rp := reply.New(w)
	s := h.s.AttachCtx(r.Context())
	file, errTF := request.TakeFile(r)
	if errTF != nil {
		rp.Error(errTF.Code, errTF.Message).Fail(errTF.Status)
		return
	}

	note, err := s.CreateOne(file)
	if err != nil {
		rp.Error(reply.CodeServerError, "Server Error", err.Error()).Fail(http.StatusInternalServerError)
		return
	}

	rp.Success(note).Created()
}

func (h *NoteHandler) UpdateOne(w http.ResponseWriter, r *http.Request) {
	rp := reply.New(w)
	s := h.s.AttachCtx(r.Context())
	id := r.PathValue("id")
	body := note.Note{}
	json.NewDecoder(r.Body).Decode(&body)

	note, err := s.UpdateByIdAndReturn(id, note.Note{Title: body.Title, Markdown: body.Markdown})
	if err != nil {
		notFoundNoteOrServerError(rp, err)
		return
	}

	rp.Success(note).Ok()
}

func (h *NoteHandler) DeleteOne(w http.ResponseWriter, r *http.Request) {
	rp := reply.New(w)
	s := h.s.AttachCtx(r.Context())
	id := r.PathValue("id")

	note, err := s.DeleteByIdAndReturn(id)
	if err != nil {
		notFoundNoteOrServerError(rp, err)
		return
	}

	rp.Success(note).Ok()
}
