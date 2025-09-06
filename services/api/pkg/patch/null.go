package patch

import (
	"encoding/json"
	"fmt"

	"github.com/atareversei/quardian/services/api/pkg/richerror"
)

type Null[T any] struct {
	IsSet bool
	Value *T
}

func (n *Null[T]) UnmarshalJSON(data []byte) error {
	const op = "patch.UnmarshalJSON"

	n.IsSet = true

	if string(data) == "null" {
		n.Value = nil
		return nil
	}

	var temp T
	if err := json.Unmarshal(data, &temp); err != nil {
		return richerror.New(op).
			WithMessage(fmt.Sprintf("failed to unmarshal into Null[%T]", temp)).
			WithError(err)
	}
	n.Value = &temp
	return nil
}

func (n Null[T]) Get() (T, bool) {
	if n.Value != nil {
		return *n.Value, true
	}
	var zero T
	return zero, false
}

func (n Null[T]) HasValue() bool {
	return n.IsSet && n.Value != nil
}
