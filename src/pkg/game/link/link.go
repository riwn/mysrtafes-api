package link

import (
	"net/url"
)

// LinkID
type ID uint64

// サイトタイトル
type Title string

// 1 ≦ title.length ≦ 256
func (t Title) Valid() bool {
	return len(t) > 0 && len(t) < 256
}

// サイトURL
type URL url.URL

// サイト先の説明
type Description string

// 1 ≦ description.length ≦ 2049
func (d Description) Valid() bool {
	return len(d) > 0 && len(d) < 2049
}

// リンク
type Link struct {
	ID          ID
	Title       Title
	URL         URL
	Description Description
}

func New(id ID, title Title, url URL, description Description) *Link {
	return &Link{
		ID:          id,
		Title:       title,
		URL:         url,
		Description: description,
	}
}
