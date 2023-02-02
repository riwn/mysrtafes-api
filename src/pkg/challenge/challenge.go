package challenge

import (
	"net/url"
	"regexp"
)

// ChallengeID
type ID uint64

// 応募者名
type Name string

// よみがな
type ReadingName string

// パスワード
type Password string

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
}

// DiscordID
type Discord string

func (d Discord) Valid() bool {
	r := regexp.MustCompile(`.+#\d{4}`)
	return r.MatchString(string(d))
}

// TwitterID
type Twitter string

func (t Twitter) Valid() bool {
	r := regexp.MustCompile(`@.+`)
	return r.MatchString(string(t))
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
	Name       Name
	Challenger Challenger
	Stream     Stream
	SNS        SNS
	Comment    Comment
}
