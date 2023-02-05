package platform

import "time"

// PlatformID
type ID uint64

// 1 ≦ id
func (i ID) Valid() bool {
	return i > 0
}

// プラットフォーム名
type Name string

// 1 ≦ name.length ≦ 256
func (n Name) Valid() bool {
	return len(n) > 0 && len(n) < 256
}

// プラットフォーム説明
type Description string

// 0 ≦ Description.length ≦ 2048
func (d Description) Valid() bool {
	return len(d) >= 0 && len(d) <= 2048
}

// プラットフォーム
type Platform struct {
	ID          ID
	Name        Name
	Description Description
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func New(name Name, description Description) *Platform {
	return &Platform{
		Name:        name,
		Description: description,
	}
}

func NewWithID(id ID, name Name, description Description) *Platform {
	return &Platform{
		ID:          id,
		Name:        name,
		Description: description,
	}
}
