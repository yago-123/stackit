package storage

import (
	"github.com/yago-123/stackit/pkg/encoding"
	"github.com/yago-123/stackit/pkg/types"
)

type Memory struct {
	// todo(): add read/write mutex for handling race conditions
	users   map[string]types.User
	encoder encoding.Encoder
}

func NewMemory(encoder encoding.Encoder) *Memory {
	return &Memory{
		users:   make(map[string]types.User),
		encoder: encoder,
	}
}

func (m *Memory) PersistUser(user types.User) error {
	m.users[user.Email] = user
	return nil
}

func (m *Memory) RetrieveUsers() ([]types.User, error) {
	// Pre-allocate all users into array
	users := make([]types.User, 0, len(m.users))

	for _, user := range m.users {
		users = append(users, user)
	}

	return users, nil
}
