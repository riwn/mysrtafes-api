package detail

import (
	"mysrtafes-backend/pkg/challenge/detail/goal"
	"mysrtafes-backend/pkg/challenge/detail/result"
)

// DetailID
type ID uint64

// 目標説明
type GoalDetail string

func (n GoalDetail) Valid() bool {
	return len(n) > 0 && len(n) < 2049
}

// 部門
type Department uint8

func (d Department) Valid() bool {
	return d < Department_MAX
}

func (d Department) String() string {
	switch d {
	case Department_BEGINNER:
		return "ちょっと不思議部門"
	case Department_EXPERT:
		return "もっと不思議部門"
	default:
		return "謎の部門"
	}
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
