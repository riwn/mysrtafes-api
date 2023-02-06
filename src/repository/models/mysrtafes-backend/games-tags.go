package mysrtafes_backend

import (
	"mysrtafes-backend/pkg/errors"
	"mysrtafes-backend/pkg/game"
	"mysrtafes-backend/pkg/game/tag"
	"time"

	"gorm.io/gorm"
)

type gameTagLink struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement"`
	GameMasterID game.ID
	TagMasterID  tag.ID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (gameTagLink) TableName() string {
	return "game_tag_links"
}

func (g *gameTagLink) BeforeCreate(db *gorm.DB) error {
	// tagの存在チェック
	tagModel := NewTagMasterFromID(g.TagMasterID)
	result := db.First(tagModel)
	if result.Error != nil {
		return errors.NewInvalidValidate(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBReadError,
				result.Error.Error(),
				[]errors.InvalidParams{
					errors.NewInvalidParams("tag_ids.id", g.TagMasterID),
				},
			),
			"tag_ids.id model is nothing error",
		)
	}
	return nil
}
