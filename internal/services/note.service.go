package services

import (
	"context"
	"markdown-note/internal/models/note"
	"markdown-note/internal/repo"
	"markdown-note/internal/services/grammar"
	"markdown-note/internal/services/markdown"
	"markdown-note/internal/services/reply"
	"net/http"
)

type NoteService struct {
	r   *repo.NoteRepo
	ctx context.Context
}

func (s *NoteService) AttachCtx(ctx context.Context) *NoteService {
	s.ctx = ctx
	return s
}

func Note(repo *repo.NoteRepo) *NoteService {
	return &NoteService{r: repo, ctx: context.Background()}
}

// Query services

func (s *NoteService) GetById(id string) (note.Note, error) {
	return s.r.GetFirst(s.ctx, []repo.Where{{N: "id", V: id}})
}

func (s *NoteService) GetAll() ([]note.Note, error) {
	return s.r.GetMany(s.ctx, []repo.Where{})
}

func (s *NoteService) CreateOne(note *note.Note) (*note.Note, error) {
	return note, s.r.CreateOne(s.ctx, note)
}

func (s *NoteService) UpdateById(id string, update note.Note) (rowsAffected int, err error) {
	return s.r.Updates(s.ctx, []repo.Where{{N: "id", V: id}}, update)
}

func (s *NoteService) UpdateByIdAndReturn(id string, update note.Note) (note.Note, error) {
	if _, err := s.UpdateById(id, update); err != nil {
		return note.Note{}, err
	}
	return s.r.GetFirst(s.ctx, []repo.Where{{N: "id", V: id}})
}

func (s *NoteService) DeleteById(id string) (rowsAffected int, err error) {
	return s.r.DeleteOne(s.ctx, []repo.Where{{N: "id", V: id}})
}

func (s *NoteService) DeleteByIdAndReturn(id string) (note.Note, error) {
	n, err := s.GetById(id)
	if err != nil {
		return note.Note{}, err
	}
	_, err = s.r.DeleteOne(s.ctx, []repo.Where{{N: "id", V: id}})
	return n, err
}

// Non query services

func (s *NoteService) FixGrammarAndUpdate(note *note.Note, rp *reply.Reply, lang, id string) {
	fixed, err := grammar.FixGrammar(note.Markdown, lang)
	if err != nil {
		rp.Error(reply.CodeBadGateWay, err.Error()).Fail(http.StatusBadGateway)
		return
	}
	note.Markdown = fixed

	_, err = s.UpdateById(id, *note)
	if err != nil {
		rp.Error(reply.CodeServerError, "Server Error", err.Error()).Fail(http.StatusInternalServerError)
		return
	}
}

func (s *NoteService) ParseNoteMarkdown(id string) (string, error) {
	note, err := s.GetById(id)
	if err != nil {
		return "", err
	}
	html := markdown.ParseMarkdown(note.Markdown)

	return html, nil
}
