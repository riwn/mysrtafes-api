package mysrtafes_backend

import (
	"mysrtafes-backend/pkg/errors"
	"mysrtafes-backend/pkg/game"
	"time"
)

type gameReferenceURLs struct {
	ID           game.LinkID `gorm:"primaryKey;autoIncrement;"`
	GameMasterID game.ID
	Title        game.Title
	URL          string
	Description  game.LinkDescription
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (gameReferenceURLs) TableName() string {
	return "game_reference_urls"
}

func NewGameReferenceURLs(link *game.Link) gameReferenceURLs {
	return gameReferenceURLs{
		ID:          link.LinkID,
		Title:       link.Title,
		URL:         link.URL.URL().String(),
		Description: link.LinkDescription,
	}
}

func (g *gameReferenceURLs) NewEntity() (*game.Link, error) {
	url, err := game.NewURL(g.URL)
	if err != nil {
		return nil, errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBCreateError,
				err.Error(),
				nil,
			),
			"game.links.url DB Data convert error",
		)
	}
	return &game.Link{
		Title:     g.Title,
		URL:       url,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}, nil
}
