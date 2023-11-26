package entities

import (
	"errors"
)

type Wallet struct {
	ID    int    `json:"id"`
	Owner string `json:"owner"`
	Count int64  `json:"count"`
}

func (w *Wallet) Validate() error {
	if w.Owner == "" || w.Count == 0 {
		return errors.New("invalid wallet structure")
	}

	return nil
}
