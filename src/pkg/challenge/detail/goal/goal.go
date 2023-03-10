package goal

import "time"

// GoalID
type ID uint64

// 1 ≦ id
func (i ID) Valid() bool {
	return i > 0
}

// 目標名
type Name string

func (n Name) Valid() bool {
	return len(n) > 0 && len(n) < 256
}

// 目標説明
type Description string

func (n Description) Valid() bool {
	return len(n) > 0 && len(n) < 2049
}

// 目標
type Goal struct {
	ID          ID
	Name        Name
	Description Description
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func New(name Name, description Description) *Goal {
	return &Goal{
		Name:        name,
		Description: description,
	}
}

func NewWithID(id ID, name Name, description Description) *Goal {
	return &Goal{
		ID:          id,
		Name:        name,
		Description: description,
	}
}
