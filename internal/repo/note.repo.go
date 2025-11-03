package repo

import (
	"context"
	"markdown-note/internal/models/note"

	"gorm.io/gorm"
)

type NoteRepo struct {
	Db *gorm.DB
}

func Note(db *gorm.DB) *NoteRepo {
	return &NoteRepo{db}
}

func (r *NoteRepo) GetFirst(ctx context.Context, where []Where) (note.Note, error) {
	w := BuildWhere(where)
	return gorm.G[note.Note](r.Db).Where(w.Q, w.V...).First(ctx)
}

func (r *NoteRepo) GetMany(ctx context.Context, where []Where) ([]note.Note, error) {
	w := BuildWhere(where)
	return gorm.G[note.Note](r.Db).Where(w.Q, w.V...).Find(ctx)
}

func (r *NoteRepo) Updates(ctx context.Context, where []Where, update note.Note) (rowsAffected int, err error) {
	w := BuildWhere(where)
	return gorm.G[note.Note](r.Db).Where(w.Q, w.V...).Updates(ctx, update)
}

func (r *NoteRepo) CreateOne(ctx context.Context, data *note.Note) error {
	return gorm.G[note.Note](r.Db).Create(ctx, data)
}

func (r *NoteRepo) DeleteOne(ctx context.Context, where []Where) (rowsAffected int, err error) {
	w := BuildWhere(where)
	return gorm.G[note.Note](r.Db).Where(w.Q, w.V...).Delete(ctx)
}
