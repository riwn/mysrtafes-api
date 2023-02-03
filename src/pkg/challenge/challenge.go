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

// TODO: 半角英数記号みなくてよい？
func (p Password) Valid() bool {
	return len(p) > 3 && len(p) < 17
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

func NewURL(us string) (URL, error) {
	u, err := url.Parse(us)
	if err != nil {
		return URL{}, err
	}
	return URL(*u), nil
}

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

// 0≦len≦16 かつxxxx#xxx形式
func (d Discord) Valid() bool {
	r := regexp.MustCompile(`.+#\d{4}`)
	return len(d) > 5 && len(d) < 37 && r.MatchString(string(d))
}

func (d Discord) Has() bool {
	return len(d) != 0
}

// TwitterID
type Twitter string

// 0≦len≦16 かつ@xxx形式
func (t Twitter) Valid() bool {
	r := regexp.MustCompile(`@[0-9a-zA-Z_]{1,15}`)
	return len(t) > 1 && len(t) < 17 && r.MatchString(string(t))
}

func (t Twitter) Has() bool {
	return len(t) != 0
}

type SNS struct {
	Discord Discord
	Twitter Twitter
}

// SNSはどちらか必須
func (s *SNS) Valid() bool {
	// 空文字ではないのにバリデートに引っかかってる時は不正な文字列と判定
	if s.Discord.Has() && !s.Discord.Valid() {
		return false
	}
	if s.Twitter.Has() && !s.Twitter.Valid() {
		return false
	}
	return s.Discord.Has() || s.Twitter.Has()
}

// 挑戦コメント
type Comment string

func (n Comment) Valid() bool {
	return len(n) > 0 && len(n) < 2049
}

// 挑戦
type Challenge struct {
	ID         ID
	Challenger Challenger
	Detail     []*detail.Detail
	Stream     Stream
	SNS        SNS
	Comment    Comment
}

func New(name Name, readingName ReadingName, password Password, twitter Twitter, discord Discord, isStream IsStream, url URL, comment Comment, details []*detail.Detail) *Challenge {
	return &Challenge{
		Challenger: Challenger{
			Name:        name,
			ReadingName: readingName,
			Password:    password,
		},
		Stream: Stream{
			IsStream: isStream,
			URL:      url,
		},
		SNS: SNS{
			Discord: discord,
			Twitter: twitter,
		},
		Comment: comment,
		Detail:  details,
	}
}
