package mysrtafes_backend

import (
	stdErrors "errors"
	"mysrtafes-backend/pkg/errors"
	"mysrtafes-backend/pkg/game"
	"time"

	"gorm.io/gorm"
)

type GameMaster interface {
	Create(*gorm.DB) error
	Read(db *gorm.DB) error
	Update(db *gorm.DB) error
	Delete(db *gorm.DB) error
	NewEntity() (*game.Game, error)
}

type gameMaster struct {
	ID                game.ID `gorm:"primaryKey;autoIncrement"`
	Name              game.Name
	Description       game.Description
	Publisher         game.Publisher
	Developer         game.Developer
	CreatedAt         time.Time
	UpdatedAt         time.Time
	GameReferenceURLs []gameReferenceURLs
	Platforms         []*platformMaster `gorm:"many2many:game_platform_links;"`
	Tags              []*tagMaster      `gorm:"many2many:game_tag_links;"`
	// TODO: LaravelでなぜかReleaseDateを追加するの忘れてたのでどこかのタイミングでmigrateと一緒に追加する。
}

func NewGameMaster(game *game.Game, platforms []*platformMaster, tags []*tagMaster) GameMaster {
	links := make([]gameReferenceURLs, 0, len(game.Links))
	for _, link := range game.Links {
		links = append(links, NewGameReferenceURLs(link))
	}
	return &gameMaster{
		Name:              game.Name,
		Description:       game.Description,
		Publisher:         game.Publisher,
		Developer:         game.Developer,
		GameReferenceURLs: links,
		Platforms:         platforms,
		Tags:              tags,
	}
}

func NewGameMasterFromID(gameID game.ID) GameMaster {
	return &gameMaster{
		ID: gameID,
	}
}

func (gameMaster) TableName() string {
	return "game_masters"
}

func (g *gameMaster) Create(db *gorm.DB) error {
	// 中間テーブルの設定
	err := stdErrors.Join(
		db.SetupJoinTable(&gameMaster{}, "Tags", &gameTagLink{}),
		db.SetupJoinTable(&gameMaster{}, "Platforms", &gamePlatformLink{}),
	)
	if err != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBCreateError,
				err.Error(),
				nil,
			),
			"create game_masters error",
		)
	}
	// NOTE: 中間テーブルのみ作成するためのOmit
	result := db.Omit("Tags.*", "Platforms.*").Create(g)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (g *gameMaster) Read(db *gorm.DB) error {
	return nil
}

func (g *gameMaster) Update(db *gorm.DB) error {
	return nil
}

func (g *gameMaster) Delete(db *gorm.DB) error {
	return nil
}

func (g *gameMaster) NewEntity() (*game.Game, error) {
	links := make([]*game.Link, 0, len(g.GameReferenceURLs))
	for _, rawLink := range g.GameReferenceURLs {
		link, err := rawLink.NewEntity()
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	// TODO: Createの時もここでTagとPlatformsをできれば入れるようにする実装を追加
	return &game.Game{
		ID:          g.ID,
		Name:        g.Name,
		Description: g.Description,
		Publisher:   g.Publisher,
		Developer:   g.Developer,
		Links:       links,
		CreatedAt:   g.CreatedAt,
		UpdatedAt:   g.UpdatedAt,
	}, nil
}
