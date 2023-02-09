package mysrtafes_backend

import (
	stdErrors "errors"
	"mysrtafes-backend/pkg/errors"
	"mysrtafes-backend/pkg/game"
	"mysrtafes-backend/pkg/game/platform"
	"mysrtafes-backend/pkg/game/tag"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		// FIXME: ここ、ModelのBeforeCreateのエラーが伝播できてない。
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBCreateError,
				result.Error.Error(),
				nil,
			),
			"create game_masters error",
		)
	}
	return nil
}

func (g *gameMaster) Read(db *gorm.DB) error {
	result := db.
		Preload("GameReferenceURLs").
		Preload("Platforms").
		Preload("Tags").
		Where("id = ?", g.ID).
		Find(&g)

	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBReadError,
				result.Error.Error(),
				nil,
			),
			"read game_masters error",
		)
	}
	return nil
}

func (g *gameMaster) Update(db *gorm.DB) error {
	return nil
}

func (g *gameMaster) Delete(db *gorm.DB) error {
	result := db.Select(
		"GameReferenceURLs",
		"Platforms",
		"Tags",
	).Delete(g)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBDeleteError,
				result.Error.Error(),
				nil,
			),
			"delete game_masters error",
		)
	}
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

	tags := make([]*tag.Tag, 0, len(g.Tags))
	for _, rawTag := range g.Tags {
		tags = append(tags, rawTag.NewEntity())
	}

	platforms := make([]*platform.Platform, 0, len(g.Platforms))
	for _, rawPlatform := range g.Platforms {
		platforms = append(platforms, rawPlatform.NewEntity())
	}

	// TODO: Createの時もここでTagとPlatformsをできれば入れるようにする実装を追加
	return &game.Game{
		ID:          g.ID,
		Name:        g.Name,
		Description: g.Description,
		Publisher:   g.Publisher,
		Developer:   g.Developer,
		Links:       links,
		Tags:        tags,
		Platforms:   platforms,
		CreatedAt:   g.CreatedAt,
		UpdatedAt:   g.UpdatedAt,
	}, nil
}

type gameMasters []*gameMaster

func NewGameMasters() gameMasters {
	return []*gameMaster{}
}

func (g *gameMasters) Find(db *gorm.DB, findOption *game.FindOption) error {
	// 検索モードで調整
	switch findOption.SearchMode {
	case game.SearchMode_Pagination:
		db = db.Limit(findOption.Pagination.Limit).Offset(findOption.Pagination.Offset)
	case game.SearchMode_Seek:
		db = db.Where("id > ?", findOption.Seek.LastID).Limit(findOption.Seek.Count)
	}

	switch findOption.OrderOption.Order {
	case game.Order_Name:
		db.Order(
			clause.OrderByColumn{
				Column: clause.Column{Name: "name"},
				Desc:   findOption.OrderOption.Desc,
			},
		)
	}

	result := db.
		Preload("GameReferenceURLs").
		Preload("Platforms").
		Preload("Tags").
		Find(&g)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBReadError,
				result.Error.Error(),
				nil,
			),
			"find game_masters error",
		)
	}
	return nil
}
