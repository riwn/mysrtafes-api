package detail

import (
	"mysrtafes-backend/pkg/challenge/detail/goal"
	"mysrtafes-backend/pkg/challenge/detail/result"
	"mysrtafes-backend/pkg/game"
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
	Game       game.Game
	Goals      []*goal.Goal
	GoalDetail GoalDetail
	Department Department
	Result     *result.Result
}

func New(gameID game.ID, gameName game.Name, goalIDs []goal.ID, goalDetail GoalDetail, department Department) *Detail {
	var goals []*goal.Goal
	for _, goalID := range goalIDs {
		goals = append(goals, &goal.Goal{
			ID: goalID,
		})
	}
	return &Detail{
		Game: game.Game{
			ID:   gameID,
			Name: gameName,
		},
		Goals:      goals,
		GoalDetail: goalDetail,
		Department: department,
	}
}
