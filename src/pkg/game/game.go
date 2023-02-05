package game

import (
	"mysrtafes-backend/pkg/game/platform"
	"mysrtafes-backend/pkg/game/tag"
	"net/url"
	"time"
)

// GameID
type ID uint64

// 1 ≦ id
func (i ID) Valid() bool {
	return i > 0
}

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

func NewReleaseDate(date string) (ReleaseDate, error) {
	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ReleaseDate{}, err
	}
	return ReleaseDate(dateTime), nil
}

func (r ReleaseDate) Time() time.Time {
	return time.Time(r)
}

// LinkID
type LinkID uint64

// 1 ≦ id
func (i LinkID) Valid() bool {
	return i > 0
}

// サイトタイトル
type Title string

// 1 ≦ title.length ≦ 256
func (t Title) Valid() bool {
	return len(t) > 0 && len(t) < 256
}

// サイトURL
type URL url.URL

func NewURL(us string) (URL, error) {
	u, err := url.ParseRequestURI(us)
	if err != nil {
		return URL{}, err
	}
	return URL(*u), nil
}

func (u URL) URL() url.URL {
	return url.URL(u)
}

// サイト先の説明
type LinkDescription string

// 0 ≦ description.length ≦ 2049
func (d LinkDescription) Valid() bool {
	return len(d) >= 0 && len(d) <= 2049
}

// リンク
type Link struct {
	LinkID          LinkID
	Title           Title
	URL             URL
	LinkDescription LinkDescription
}

func NewLink(title Title, url URL, description LinkDescription) *Link {
	return &Link{
		Title:           title,
		URL:             url,
		LinkDescription: description,
	}
}

func NewLinkWithID(id LinkID, title Title, url URL, description LinkDescription) *Link {
	return &Link{
		LinkID:          id,
		Title:           title,
		URL:             url,
		LinkDescription: description,
	}
}

// ゲームマスタ
type Game struct {
	ID          ID
	Name        Name
	Description Description
	Publisher   Publisher
	Developer   Developer
	ReleaseDate ReleaseDate
	Links       []*Link
	Platforms   []*platform.Platform
	Tags        []*tag.Tag
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func New(
	name Name,
	description Description,
	publisher Publisher,
	developer Developer,
	releaseDate ReleaseDate,
	links []*Link,
) *Game {
	return &Game{
		Name:        name,
		Description: description,
		Publisher:   publisher,
		Developer:   developer,
		ReleaseDate: releaseDate,
		Links:       links,
	}
}

func NewWithID(
	id ID,
	name Name,
	description Description,
	publisher Publisher,
	developer Developer,
	releaseDate ReleaseDate,
	links []*Link,
) *Game {
	return &Game{
		ID:          id,
		Name:        name,
		Description: description,
		Publisher:   publisher,
		Developer:   developer,
		ReleaseDate: releaseDate,
		Links:       links,
	}
}
