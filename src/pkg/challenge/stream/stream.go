package stream

import (
	"net/url"
	"time"
)

// ChallengeID
type ID uint64

// 配信中フラグ
type IsLive bool

// 総配信時間
type TotalLiveTime time.Duration

// 配信タイトル
type Title string

// 配信リンク
type LiveURL url.URL

// サムネイル
type Thumbnail url.URL

// 配信開始時間
type LiveStartTime time.Time

// 配信詳細
type Detail struct {
	LiveStartTime LiveStartTime
	Title         Title
	LiveURL       LiveURL
	Thumbnail     Thumbnail
	TotalLiveTime TotalLiveTime
}

// 最終更新日
type LastUpdate time.Time

// 配信状況
type Status struct {
	ID         ID
	IsLive     IsLive
	Detail     Detail
	LastUpdate LastUpdate
}
