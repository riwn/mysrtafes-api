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

// 1 ≦ name.length ≦ 255
func (n Name) Valid() bool {
	return len(n) > 0 && len(n) < 256
}

// ゲーム説明
type Description string

// 0 ≦ Description.length ≦ 2048
func (d Description) Valid() bool {
	// NOTE: 必須情報ではない
	return len(d) >= 0 && len(d) <= 2048
}

// 企画元
type Publisher string

// 0 ≦ Publisher.length ≦ 255
func (p Publisher) Valid() bool {
	// NOTE: 必須情報ではない
	return len(p) >= 0 && len(p) < 256
}

// 開発元
type Developer string

// 0 ≦ Developer.length ≦ 255
func (d Developer) Valid() bool {
	// NOTE: 必須情報ではない
	return len(d) >= 0 && len(d) < 256
}

// 発売日
type ReleaseDate time.Time

func NewReleaseDate(date string) (ReleaseDate, error) {
	// NOTE: 空文字判定のValidationめんどくさいので必須にしたい。
	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ReleaseDate{}, err
	}
	return ReleaseDate(dateTime), nil
}

func (r ReleaseDate) Time() time.Time {
	return time.Time(r)
}

func (r ReleaseDate) String() string {
	return r.Time().Format("2006-01-02")
}

// LinkID
type LinkID uint64

// 1 ≦ id
func (i LinkID) Valid() bool {
	return i > 0
}

// サイトタイトル
type Title string

// 1 ≦ title.length ≦ 255
func (t Title) Valid() bool {
	// NOTE: サイト追加するのであればタイトルは必須にする
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

func (u URL) URL() *url.URL {
	url := url.URL(u)
	return &url
}

// サイト先の説明
type LinkDescription string

// 0 ≦ description.length ≦ 2049
func (d LinkDescription) Valid() bool {
	// NOTE: 必須情報ではない
	return len(d) >= 0 && len(d) <= 2048
}

// リンク
type Link struct {
	LinkID          LinkID
	Title           Title
	URL             URL
	LinkDescription LinkDescription
	CreatedAt       time.Time
	UpdatedAt       time.Time
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
