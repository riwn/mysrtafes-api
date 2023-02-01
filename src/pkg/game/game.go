package game

import (
	"fmt"
	"mysrtafes-backend/pkg/game/link"
	"mysrtafes-backend/pkg/game/platform"
	"mysrtafes-backend/pkg/game/tag"
	"time"
)

// TODO: Validationメソッドの追加

// GameID
type ID uint64

// ゲームタイトル
type Name string

// ゲーム説明
type Description string

// 企画元
type Publisher string

// 開発元
type Developer string

// 発売日
type ReleaseDate time.Time

func (r ReleaseDate) Time() time.Time {
	return time.Time(r)
}

// ゲームマスタ
type Game struct {
	ID          ID
	Name        Name
	Description Description
	Publisher   Publisher
	Developer   Developer
	ReleaseDate ReleaseDate
	Links       []*link.Link
	Platforms   []*platform.Platform
	Tags        []*tag.Tag
}

func New(
	id ID,
	name Name,
	description Description,
	publisher Publisher,
	developer Developer,
	releaseDate ReleaseDate,
	links []*link.Link,
	platforms []*platform.Platform,
) *Game {
	return &Game{
		ID:          id,
		Name:        name,
		Description: description,
		Publisher:   publisher,
		Developer:   developer,
		ReleaseDate: releaseDate,
		Links:       links,
		Platforms:   platforms,
	}
}

func (g Game) String() string {
	return fmt.Sprintf("「%s」は企画元 %s, 開発元 %sによって%sにリリースされた%sです", g.Name, g.Publisher, g.Developer, g.ReleaseDate.Time(), g.Description)
}
