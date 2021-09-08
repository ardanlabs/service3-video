// Package user contains user related CRUD functionality.
package user

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Store manages the set of API's for user access.
type Store struct {
	log *zap.SugaredLogger
	db  *sqlx.DB
}

// NewStore constructs a user store for api access.
func NewStore(log *zap.SugaredLogger, db *sqlx.DB) Store {
	return Store{
		log: log,
		db:  db,
	}
}
