package mysrtafes_backend

import (
	"mysrtafes-backend/pkg/challenge/detail/goal"
	"mysrtafes-backend/pkg/errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GoalGenreMaster interface {
	Create(*gorm.DB) error
	Read(db *gorm.DB) error
	Update(db *gorm.DB) error
	Delete(db *gorm.DB) error
	NewEntity() *goal.Goal
}

type goalGenreMaster struct {
	ID          goal.ID `gorm:"primaryKey;autoIncrement"`
	Name        goal.Name
	Description goal.Description
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (goalGenreMaster) TableName() string {
	return "goal_genre_masters"
}

func NewGoalGenreMaster(goal *goal.Goal) GoalGenreMaster {
	return &goalGenreMaster{
		ID:          goal.ID,
		Name:        goal.Name,
		Description: goal.Description,
	}
}

func NewGoalGenreMasterFromID(goalID goal.ID) GoalGenreMaster {
	return &goalGenreMaster{
		ID: goalID,
	}
}

func NewGoalGenreMasterFromGoals(goals []*goal.Goal) []*goalGenreMaster {
	goalModels := make([]*goalGenreMaster, 0, len(goals))
	for _, goal := range goals {
		goalModels = append(goalModels, &goalGenreMaster{
			ID: goal.ID,
		})
	}
	return goalModels
}

func (g *goalGenreMaster) Create(db *gorm.DB) error {
	result := db.Create(g)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBCreateError,
				result.Error.Error(),
				nil,
			),
			"create goal_genre_masters error",
		)
	}
	return nil
}

func (g *goalGenreMaster) Read(db *gorm.DB) error {
	result := db.Where("id = ?", g.ID).Find(&g)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBReadError,
				result.Error.Error(),
				nil,
			),
			"read goal_genre_masters error",
		)
	}
	return nil
}

func (g *goalGenreMaster) Update(db *gorm.DB) error {
	// TODO: 更新の時だけCreatedAtが入ってこない問題がある。
	result := db.Updates(g)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBUpdateError,
				result.Error.Error(),
				nil,
			),
			"update goal_genre_masters error",
		)
	}
	return nil
}

func (g *goalGenreMaster) Delete(db *gorm.DB) error {
	result := db.Delete(g)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBDeleteError,
				result.Error.Error(),
				nil,
			),
			"delete goal_genre_masters error",
		)
	}
	return nil
}

func (g *goalGenreMaster) NewEntity() *goal.Goal {
	return &goal.Goal{
		ID:          g.ID,
		Name:        g.Name,
		Description: g.Description,
		CreatedAt:   g.CreatedAt,
		UpdatedAt:   g.UpdatedAt,
	}
}

type GoalGenreMasters interface {
	Find(db *gorm.DB, findOption *goal.FindOption) error
	NewEntities() []*goal.Goal
}

type goalGenreMasters []*goalGenreMaster

func NewGoalGenreMasters() GoalGenreMasters {
	return goalGenreMasters{}
}

func (p goalGenreMasters) Find(db *gorm.DB, findOption *goal.FindOption) error {
	// 検索モードで調整
	switch findOption.SearchMode {
	case goal.SearchMode_Pagination:
		db = db.Limit(findOption.Pagination.Limit).Offset(findOption.Pagination.Offset)
	case goal.SearchMode_Seek:
		db = db.Where("id > ?", findOption.Seek.LastID).Limit(findOption.Seek.Count)
	}

	switch findOption.OrderOption.Order {
	case goal.Order_Name:
		db.Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}, Desc: findOption.OrderOption.Desc})
	}

	result := db.Find(&p)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBReadError,
				result.Error.Error(),
				nil,
			),
			"find goal_genre_masters error",
		)
	}
	return nil
}

func (p goalGenreMasters) NewEntities() []*goal.Goal {
	entities := make([]*goal.Goal, 0, len(p))
	for _, model := range p {
		entities = append(entities, model.NewEntity())
	}
	return entities
}
