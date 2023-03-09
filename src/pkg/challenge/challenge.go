package challenge

import (
	"mysrtafes-backend/pkg/challenge/detail"
	"mysrtafes-backend/pkg/challenge/stream"
	"mysrtafes-backend/pkg/errors"
	"net/url"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// ChallengeID
type ID uint64

// id > 0
func (i ID) Valid() bool {
	return i > 0
}

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

const PasswordCryptCost = 10

// パスワード
type Password string

// TODO: 半角英数記号みなくてよい？
func (p Password) Valid() bool {
	return len(p) > 3 && len(p) < 17
}

func (p Password) String() string {
	return "****"
}

func (p Password) Hash() (Password, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(p), PasswordCryptCost)
	if err != nil {
		return "", err
	}
	return Password(hashPassword), nil
}

func (p Password) Check(password Password) error {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(password))
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

func (u URL) StreamSite() StreamSite {
	switch {
	case strings.Contains(u.URL().String(), "nicovideo"):
		return StreamStatus_Niconico
	case strings.Contains(u.URL().String(), "youtube"):
		return StreamStatus_YouTube
	case strings.Contains(u.URL().String(), "twitch"):
		return StreamStatus_Twitch
	default:
		return StreamStatus_Unknown
	}
}

// 配信データ
type Stream struct {
	IsStream IsStream
	URL      URL
	Status   *stream.Status
}

// 配信サイト
type StreamSite uint16

const (
	_ StreamSite = iota
	StreamStatus_Niconico
	StreamStatus_YouTube
	StreamStatus_Twitch
	StreamStatus_Unknown
)

func (s StreamSite) String() string {
	switch s {
	case StreamStatus_Niconico:
		return "ニコニコ生放送"
	case StreamStatus_YouTube:
		return "YouTube"
	case StreamStatus_Twitch:
		return "Twitch"
	default:
		return ""
	}
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
	Details    []*detail.Detail
	Stream     Stream
	SNS        SNS
	Comment    Comment
	CreatedAt  time.Time
	UpdatedAt  time.Time
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
		Details: details,
	}
}

func (c *Challenge) ValidCreate() error {
	// NameのValidate
	if !c.Challenger.Name.Valid() {
		return errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("name", c.Challenger.Name),
				},
			),
			"Name Valid error",
		)
	}
	// ReadingNameのValidate
	if !c.Challenger.ReadingName.Valid() {
		return errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("name_read", c.Challenger.ReadingName),
				},
			),
			"ReadingName Valid error",
		)
	}

	// PasswordのValidate
	if !c.Challenger.Password.Valid() {
		return errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("password", c.Challenger.Password.String()),
				},
			),
			"Password Valid error",
		)
	}

	// SNSのValidate
	if !c.SNS.Valid() {
		return errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("discord", c.SNS.Discord),
					errors.NewInvalidParams("twitter", c.SNS.Twitter),
				},
			),
			"SNS Valid error",
		)
	}

	// CommentのValidate
	if !c.Comment.Valid() {
		return errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("comment", c.Comment),
				},
			),
			"Comment Valid error",
		)
	}

	// detailがないときはエラー
	if len(c.Details) == 0 {
		return errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("details count", len(c.Details)),
				},
			),
			"Nothing Detail Error",
		)
	}

	// detailsのValid
	for _, detail := range c.Details {
		if err := detail.ValidCreate(); err != nil {
			return err
		}
	}
	return nil
}

func (c *Challenge) ValidUpdate() error {
	// IDのValidate
	if !c.ID.Valid() {
		return errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("id", c.ID),
				},
			),
			"ID Valid error",
		)
	}

	// Createと同じValidateを流す
	if err := c.ValidCreate(); err != nil {
		return err
	}

	return nil
}
