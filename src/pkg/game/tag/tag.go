package tag

import (
	"time"
)

// TODO: Validationメソッドの追加

// TagID
type ID uint64

// 1 ≦ id
func (i ID) Valid() bool {
	return i > 0
}

// タグ名
type Name string

// 1 ≦ name.length ≦ 256
func (n Name) Valid() bool {
	return len(n) > 0 && len(n) < 256
}

// タグ説明
type Description string

// 0 ≦ Description.length ≦ 2049
func (d Description) Valid() bool {
	// NOTE: 説明は空でもOK
	return len(d) >= 0 && len(d) < 2049
}

// タグ
type Tag struct {
	ID          ID
	Name        Name
	Description Description
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func New(name Name, description Description) *Tag {
	return &Tag{
		Name:        name,
		Description: description,
	}
}

func NewWithID(id ID, name Name, description Description) *Tag {
	return &Tag{
		ID:          id,
		Name:        name,
		Description: description,
	}
}
