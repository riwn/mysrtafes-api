package mysrtafes_backend

import (
	"mysrtafes-backend/pkg/game/tag"
	"time"

	"gorm.io/gorm"
)

type TagMaster interface {
	Create(*gorm.DB) error
	NewEntity() *tag.Tag
}

type tagMaster struct {
	ID          tag.ID `gorm:"primaryKey;autoIncrement"`
	Name        tag.Name
	Description tag.Description
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTagMaster(tag *tag.Tag) TagMaster {
	return &tagMaster{
		Name:        tag.Name,
		Description: tag.Description,
	}
}

func (tagMaster) TableName() string {
	return "tag_masters"
}

func (t *tagMaster) Create(db *gorm.DB) error {
	result := db.Create(t)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *tagMaster) NewEntity() *tag.Tag {
	return &tag.Tag{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
