package storage

import (
	"github.com/yago-123/stackit/pkg/types"
)

type Store interface {
	// todo(): rename from persist to other more convenient name
	PersistUser(user types.User) error
	RetrieveUsers() ([]types.User, error)
}
