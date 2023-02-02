package goal

// GoalID
type ID uint64

// 目標名
type Name string

// 目標説明
type Description string

// 目標
type Goal struct {
	ID          ID
	Name        Name
	Description Description
}
