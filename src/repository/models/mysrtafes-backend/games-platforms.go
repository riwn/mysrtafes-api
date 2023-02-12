package mysrtafes_backend

import (
	"mysrtafes-backend/pkg/errors"
	"mysrtafes-backend/pkg/game"
	"mysrtafes-backend/pkg/game/platform"
	"time"

	"gorm.io/gorm"
)

type gamePlatformLink struct {
	ID               uint64 `gorm:"primaryKey;autoIncrement"`
	GameMasterID     game.ID
	PlatformMasterID platform.ID
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (gamePlatformLink) TableName() string {
	return "game_platform_links"
}

func (g *gamePlatformLink) BeforeCreate(db *gorm.DB) error {
	// platformの存在チェック
	platformModel := NewPlatformMasterFromID(g.PlatformMasterID)
	result := db.First(platformModel)
	if result.Error != nil {
		return errors.NewInvalidValidate(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBReadError,
				result.Error.Error(),
				[]errors.InvalidParams{
					errors.NewInvalidParams("platform_ids.id", g.PlatformMasterID),
				},
			),
			"platform_ids.id model is nothing error",
		)
	}
	return nil
}
