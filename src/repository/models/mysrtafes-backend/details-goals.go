package mysrtafes_backend

import (
	"mysrtafes-backend/pkg/challenge/detail"
	"mysrtafes-backend/pkg/challenge/detail/goal"
	"mysrtafes-backend/pkg/errors"
	"time"

	"gorm.io/gorm"
)

type detailGoalLink struct {
	ID                uint64 `gorm:"primaryKey;autoIncrement"`
	ChallengeDetailID detail.ID
	GoalGenreMasterID goal.ID
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (detailGoalLink) TableName() string {
	return "mys_challenge2_goal_genre_detail_links"
}

func (d *detailGoalLink) BeforeCreate(db *gorm.DB) error {
	// goalの存在チェック
	goalModel := NewGoalGenreMasterFromID(d.GoalGenreMasterID)
	result := db.First(goalModel)
	if result.Error != nil {
		return errors.NewInvalidValidate(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBReadError,
				result.Error.Error(),
				[]errors.InvalidParams{
					errors.NewInvalidParams("goal_ids.id", d.GoalGenreMasterID),
				},
			),
			"goal_ids.id model is nothing error",
		)
	}
	return nil
}
