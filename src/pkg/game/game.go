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

// 1 ≦ name.length ≦ 256
func (n Name) Valid() bool {
	return len(n) > 0 && len(n) < 256
}

// ゲーム説明
type Description string

// 1 ≦ Description.length ≦ 2049
func (d Description) Valid() bool {
	return len(d) > 0 && len(d) < 2049
}

// 企画元
type Publisher string

// 1 ≦ Publisher.length ≦ 256
func (p Publisher) Valid() bool {
	return len(p) > 0 && len(p) < 256
}

// 開発元
type Developer string

// 1 ≦ Developer.length ≦ 256
func (d Developer) Valid() bool {
	return len(d) > 0 && len(d) < 256
}

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
