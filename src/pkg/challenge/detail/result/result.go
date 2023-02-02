package result

import "net/url"

// ResultID
type ID uint64

// 達成したか
type IsAchievement bool

// 画像URL
type Image url.URL

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
}
