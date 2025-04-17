package encoding

import (
	"github.com/yago-123/stackit/pkg/types"
)

type Encoder interface {
	// todo(): maybe interface instead of user
	EncodeUser(user types.User) ([]byte, error)
	DecodeUser(data []byte) (types.User, error)
}
