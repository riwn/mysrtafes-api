package result

import (
	"net/url"
	"time"
)

// ResultID
type ID uint64

// id > 0
func (i ID) Valid() bool {
	return i > 0
}

// 達成したか
type IsAchievement bool

// 画像URL
type Image url.URL

func NewImage(us string) (Image, error) {
	u, err := url.ParseRequestURI(us)
	if err != nil {
		return Image{}, err
	}
	return Image(*u), nil
}

func (i Image) URL() *url.URL {
	url := url.URL(i)
	return &url
}

// コメント
type Comment string

func (n Comment) Valid() bool {
	return len(n) > 0 && len(n) < 2049
}

// 結果
type Result struct {
	ID            ID
	IsAchievement IsAchievement
	Image         Image
	Comment       Comment
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
