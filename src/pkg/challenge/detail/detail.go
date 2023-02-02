package detail

import (
	"mysrtafes-backend/pkg/challenge/detail/goal"
	"mysrtafes-backend/pkg/challenge/detail/result"
)

// DetailID
type ID uint64

// 目標説明
type GoalDetail string

// 部門
type Department uint8

func (d Department) Valid() bool {
	return d < Department_MAX
}

const (
	Department_BEGINNER Department = iota
	Department_EXPERT
	Department_MAX
)

type Detail struct {
	ID         ID
	Goal       goal.Goal
	GoalDetail GoalDetail
	Department Department
	Result     *result.Result
}
