package encoding

import (
	"github.com/yago-123/stackit/pkg/types"
)

type JSONEncoder struct {
}

func NewJson() *JSONEncoder {
	return &JSONEncoder{}
}

func (json *JSONEncoder) EncodeUser(_ types.User) ([]byte, error) {
	return []byte{}, nil
}

func (json *JSONEncoder) DecodeUser(_ []byte) (types.User, error) {
	return types.User{}, nil
}
