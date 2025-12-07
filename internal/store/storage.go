package store

import (
	"context"
	"database/sql"
	"tiagoapi/internal/model"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *model.Post) error
	}
	Users interface {
		Create(context.Context, *model.User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UserStore{db},
	}
}
