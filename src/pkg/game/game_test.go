package game

import (
	"math/rand"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestID_Valid(t *testing.T) {
	tests := []struct {
		name string
		i    ID
		want bool
	}{
		{
			name: "OK",
			i:    1,
			want: true,
		},
		{
			name: "NG",
			i:    0,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.i.Valid())
		})
	}
}

func TestName_Valid(t *testing.T) {
	tests := []struct {
		name string
		n    Name
		want bool
	}{
		{
			name: "OK",
			n:    "風来のシレン2",
			want: true,
		},
		{
			name: "空文字",
			n:    "",
			want: false,
		},
		{
			name: "長すぎる文字列のギリギリ",
			n: func() Name {
				var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

				s := make([]rune, 255)
				for i := range s {
					s[i] = letters[rand.Intn(len(letters))]
				}
				return Name(s)
			}(),
			want: true,
		},
		{
			name: "長すぎる文字列",
			n: func() Name {
				var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

				s := make([]rune, 256)
				for i := range s {
					s[i] = letters[rand.Intn(len(letters))]
				}
				return Name(s)
			}(),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.n.Valid())
		})
	}
}

func TestDescription_Valid(t *testing.T) {
	tests := []struct {
		name string
		d    Description
		want bool
	}{
		{
			name: "OK",
			d:    "城を建てます",
			want: true,
		},
		{
			name: "空文字",
			d:    "",
			want: true,
		},
		{
			name: "長すぎる文字列のギリギリ",
			d: func() Description {
				var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

				s := make([]rune, 2048)
				for i := range s {
					s[i] = letters[rand.Intn(len(letters))]
				}
				return Description(s)
			}(),
			want: true,
		},
		{
			name: "長すぎる文字列",
			d: func() Description {
				var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

				s := make([]rune, 2049)
				for i := range s {
					s[i] = letters[rand.Intn(len(letters))]
				}
				return Description(s)
			}(),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.d.Valid())
		})
	}
}

func TestPublisher_Valid(t *testing.T) {
	tests := []struct {
		name string
		p    Publisher
		want bool
	}{
		{
			name: "OK",
			p:    "中村光一",
			want: true,
		},
		{
			name: "空文字",
			p:    "",
			want: true,
		},
		{
			name: "長すぎる文字列のギリギリ",
			p: func() Publisher {
				var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

				s := make([]rune, 255)
				for i := range s {
					s[i] = letters[rand.Intn(len(letters))]
				}
				return Publisher(s)
			}(),
			want: true,
		},
		{
			name: "長すぎる文字列",
			p: func() Publisher {
				var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

				s := make([]rune, 256)
				for i := range s {
					s[i] = letters[rand.Intn(len(letters))]
				}
				return Publisher(s)
			}(),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.p.Valid())
		})
	}
}

func TestDeveloper_Valid(t *testing.T) {
	tests := []struct {
		name string
		d    Developer
		want bool
	}{
		{
			name: "OK",
			d:    "大三元ソフトウェア",
			want: true,
		},
		{
			name: "空文字",
			d:    "",
			want: true,
		},
		{
			name: "長すぎる文字列のギリギリ",
			d: func() Developer {
				var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

				s := make([]rune, 255)
				for i := range s {
					s[i] = letters[rand.Intn(len(letters))]
				}
				return Developer(s)
			}(),
			want: true,
		},
		{
			name: "長すぎる文字列",
			d: func() Developer {
				var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

				s := make([]rune, 256)
				for i := range s {
					s[i] = letters[rand.Intn(len(letters))]
				}
				return Developer(s)
			}(),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.d.Valid())
		})
	}
}

func Test発売日の生成(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    ReleaseDate
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				date: "2000-09-27",
			},
			want: func() ReleaseDate {
				date := time.Date(2000, 9, 27, 0, 0, 0, 0, time.UTC)
				return ReleaseDate(date)
			}(),
			wantErr: false,
		},
		{
			name: "空文字",
			args: args{
				date: "",
			},
			wantErr: true,
		},
		{
			name: "指定フォーマットではない",
			args: args{
				date: "2000/09/27",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewReleaseDate(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewReleaseDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReleaseDate_String(t *testing.T) {
	tests := []struct {
		name string
		r    ReleaseDate
		want string
	}{
		{
			name: "OK",
			r: func() ReleaseDate {
				date := time.Date(2000, 9, 27, 0, 0, 0, 0, time.UTC)
				return ReleaseDate(date)
			}(),
			want: "2000-09-27",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.r.String())
		})
	}
}

func Test発売日から時間への型変換(t *testing.T) {
	tests := []struct {
		name string
		r    ReleaseDate
		want time.Time
	}{
		{
			name: "変換",
			r: func() ReleaseDate {
				t := time.Date(2000, 9, 27, 0, 0, 0, 0, time.UTC)
				return ReleaseDate(t)
			}(),
			want: func() time.Time {
				t := time.Date(2000, 9, 27, 0, 0, 0, 0, time.UTC)
				return t
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.r.Time())
		})
	}
}

func TestLinkID_Valid(t *testing.T) {
	tests := []struct {
		name string
		i    LinkID
		want bool
	}{
		{
			name: "OK",
			i:    1,
			want: true,
		},
		{
			name: "NG",
			i:    0,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.i.Valid())
		})
	}
}

func TestTitle_Valid(t *testing.T) {
	tests := []struct {
		name string
		tr   Title
		want bool
	}{
		{
			name: "OK",
			tr:   "wiki",
			want: true,
		},
		{
			name: "空文字",
			tr:   "",
			want: false,
		},
		{
			name: "長すぎる文字列のギリギリ",
			tr: func() Title {
				var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

				s := make([]rune, 255)
				for i := range s {
					s[i] = letters[rand.Intn(len(letters))]
				}
				return Title(s)
			}(),
			want: true,
		},
		{
			name: "長すぎる文字列",
			tr: func() Title {
				var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

				s := make([]rune, 256)
				for i := range s {
					s[i] = letters[rand.Intn(len(letters))]
				}
				return Title(s)
			}(),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.tr.Valid())
		})
	}
}

func TestNewURL(t *testing.T) {
	type args struct {
		us string
	}
	tests := []struct {
		name    string
		args    args
		want    URL
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				us: "https://example.com",
			},
			want: func() URL {
				u, _ := url.Parse("https://example.com")
				return URL(*u)
			}(),
			wantErr: false,
		},
		{
			name: "空文字",
			args: args{
				us: "",
			},
			wantErr: true,
		},
		{
			name: "でたらめな文字列",
			args: args{
				us: "中村光一",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewURL(tt.args.us)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestURL_URL(t *testing.T) {
	tests := []struct {
		name string
		u    URL
		want *url.URL
	}{
		{
			name: "OK",
			u: func() URL {
				u, _ := NewURL("https://example.com")
				return u
			}(),
			want: func() *url.URL {
				u, _ := url.ParseRequestURI("https://example.com")
				return u
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.u.URL())
		})
	}
}

func TestLinkDescription_Valid(t *testing.T) {
	tests := []struct {
		name string
		d    LinkDescription
		want bool
	}{
		{
			name: "OK",
			d:    "シレン2の攻略情報を扱うwikiです。",
			want: true,
		},
		{
			name: "空文字",
			d:    "",
			want: true,
		},
		{
			name: "長すぎる文字列のギリギリ",
			d: func() LinkDescription {
				var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

				s := make([]rune, 2048)
				for i := range s {
					s[i] = letters[rand.Intn(len(letters))]
				}
				return LinkDescription(s)
			}(),
			want: true,
		},
		{
			name: "長すぎる文字列",
			d: func() LinkDescription {
				var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

				s := make([]rune, 2049)
				for i := range s {
					s[i] = letters[rand.Intn(len(letters))]
				}
				return LinkDescription(s)
			}(),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.d.Valid())
		})
	}
}

func TestNewLink(t *testing.T) {
	type args struct {
		title       Title
		url         URL
		description LinkDescription
	}
	tests := []struct {
		name string
		args args
		want *Link
	}{
		{
			name: "OK",
			args: args{
				title: Title("Test"),
				url: func() URL {
					url, _ := NewURL("https://example.com")
					return url
				}(),
				description: LinkDescription("testです"),
			},
			want: func() *Link {
				url, _ := NewURL("https://example.com")
				link := &Link{
					Title:           Title("Test"),
					URL:             url,
					LinkDescription: LinkDescription("testです"),
				}
				return link
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLink(tt.args.title, tt.args.url, tt.args.description)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewLinkWithID(t *testing.T) {
	type args struct {
		id          LinkID
		title       Title
		url         URL
		description LinkDescription
	}
	tests := []struct {
		name string
		args args
		want *Link
	}{
		{
			name: "OK",
			args: args{
				id:    100000,
				title: Title("Test"),
				url: func() URL {
					url, _ := NewURL("https://example.com")
					return url
				}(),
				description: LinkDescription("testです"),
			},
			want: func() *Link {
				url, _ := NewURL("https://example.com")
				link := &Link{
					LinkID:          LinkID(100000),
					Title:           Title("Test"),
					URL:             url,
					LinkDescription: LinkDescription("testです"),
				}
				return link
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLinkWithID(tt.args.id, tt.args.title, tt.args.url, tt.args.description)
			assert.Equal(t, tt.want, got)
		})
	}
}
