package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: подходящей записпи не найдено")

type Snippet struct {
	Id             int
	Title, Content string
}
