package goal

// GoalID
type ID uint64

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
}
