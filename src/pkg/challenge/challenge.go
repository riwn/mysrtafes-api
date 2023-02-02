package challenge

import (
	"mysrtafes-backend/pkg/challenge/detail"
	"mysrtafes-backend/pkg/challenge/stream"
	"net/url"
	"regexp"
)

// ChallengeID
type ID uint64

// 応募者名
type Name string

func (n Name) Valid() bool {
	return len(n) > 0 && len(n) < 256
}

// よみがな
type ReadingName string

func (r ReadingName) Valid() bool {
	return len(r) > 0 && len(r) < 256
}

// パスワード
type Password string

func (p Password) Valid() bool {
	return len(p) > 0 && len(p) < 16
}

func (p Password) String() string {
	return "****"
}

// 応募者情報
type Challenger struct {
	Name        Name
	ReadingName ReadingName
	Password    Password
}

// 配信するかどうか
type IsStream bool

// 配信サイト
type URL url.URL

func (u URL) URL() url.URL {
	return url.URL(u)
}

// 配信データ
type Stream struct {
	IsStream IsStream
	URL      URL
	Status   *stream.Status
}

// DiscordID
type Discord string

// 0≦len≦16 かつ@xxx形式
func (d Discord) Valid() bool {
	r := regexp.MustCompile(`.+#\d{4}`)
	return len(d) > 5 && len(d) < 37 && r.MatchString(string(d))
}

// TwitterID
type Twitter string

// 0≦len≦16 かつ@xxx形式
func (t Twitter) Valid() bool {
	r := regexp.MustCompile(`@.+`)
	return len(t) > 1 && len(t) < 16 && r.MatchString(string(t))
}

type SNS struct {
	Discord Discord
	Twitter Twitter
}

// SNSはどちらか必須
func (s *SNS) Valid() bool {
	return s.Discord.Valid() || s.Twitter.Valid()
}

// 挑戦コメント
type Comment string

// 挑戦
type Challenge struct {
	ID         ID
	Challenger Challenger
	Detail     []*detail.Detail
	Stream     Stream
	SNS        SNS
	Comment    Comment
}
