package challenge

import (
	"math/rand"
	"net/url"
	"reflect"
	"testing"
)

func Test挑戦者名バリデート(t *testing.T) {
	tests := []struct {
		name string
		n    Name
		want bool
	}{
		{
			name: "OK",
			n:    "あーる",
			want: true,
		},
		{
			name: "空文字",
			n:    "",
			want: false,
		},
		{
			name: "長すぎる名前",
			n:    "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゐゆゑよらりるれろわをんあいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゐゆゑよらりるれろわをんあいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゐゆゑよらりるれろわをんあいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゐゆゑよらりるれろわをんあいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゐゆゑよらりるれろわをんあいうえおかきくけこさしすせそた",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Valid(); got != tt.want {
				t.Errorf("Name.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Testよみがなバリデート(t *testing.T) {
	tests := []struct {
		name string
		r    ReadingName
		want bool
	}{
		{
			name: "OK",
			r:    "あーる",
			want: true,
		},
		{
			name: "空文字",
			r:    "",
			want: false,
		},
		{
			name: "長すぎる名前",
			r:    "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゐゆゑよらりるれろわをんあいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゐゆゑよらりるれろわをんあいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゐゆゑよらりるれろわをんあいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゐゆゑよらりるれろわをんあいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゐゆゑよらりるれろわをんあいうえおかきくけこさしすせそた",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Valid(); got != tt.want {
				t.Errorf("ReadingName.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Testパスワードバリデート(t *testing.T) {
	tests := []struct {
		name string
		p    Password
		want bool
	}{
		{
			name: "OK",
			p:    "12345678",
			want: true,
		},
		{
			name: "空文字",
			p:    "",
			want: false,
		},
		{
			name: "3文字",
			p:    "123",
			want: false,
		},
		{
			name: "4文字",
			p:    "1234",
			want: true,
		},
		{
			name: "長すぎるパスワード",
			p:    "12345678901234567",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Valid(); got != tt.want {
				t.Errorf("Password.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Testパスワード出力(t *testing.T) {
	tests := []struct {
		name string
		p    Password
		want string
	}{
		{
			name: "マスキング",
			p:    "12345678",
			want: "****",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Password.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURL変換(t *testing.T) {
	tests := []struct {
		name string
		u    URL
		want url.URL
	}{
		{
			name: "変換",
			u: func() URL {
				uri := "http://example.com"
				exampleURL, err := url.Parse(uri)
				if err != nil {
					panic(err)
				}
				return URL(*exampleURL)
			}(),
			want: func() url.URL {
				uri := "http://example.com"
				exampleURL, err := url.Parse(uri)
				if err != nil {
					panic(err)
				}
				return *exampleURL
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.URL(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URL.URL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscordIDバリデート(t *testing.T) {
	tests := []struct {
		name string
		d    Discord
		want bool
	}{
		{
			name: "OK",
			d:    "あーる#4600",
			want: true,
		},
		{
			name: "ギリギリのみじかさ",
			d:    "a#1000",
			want: true,
		},
		{
			name: "ギリギリの長さ",
			d:    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa#1000",
			want: true,
		},
		{
			name: "空文字",
			d:    "",
			want: false,
		},
		{
			name: "フォーマット違い",
			d:    "あーる@example.com",
			want: false,
		},
		{
			name: "短すぎる",
			d:    "#1000",
			want: false,
		},
		{
			name: "長すぎる",
			d:    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa#1000",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Valid(); got != tt.want {
				t.Errorf("Discord.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwitterIDバリデート(t *testing.T) {
	tests := []struct {
		name string
		tr   Twitter
		want bool
	}{
		{
			name: "OK",
			tr:   "@r_sprl",
			want: true,
		},
		{
			name: "ギリギリのみじかさ",
			tr:   "@a",
			want: true,
		},
		{
			name: "ギリギリの長さ",
			tr:   "@aaaaaaaaaaaaaaa",
			want: true,
		},
		{
			name: "空文字",
			tr:   "",
			want: false,
		},
		{
			name: "フォーマット違い",
			tr:   "aaaaa#example",
			want: false,
		},
		{
			name: "短すぎる",
			tr:   "@",
			want: false,
		},
		{
			name: "長すぎる",
			tr:   "@aaaaaaaaaaaaaaaa",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Valid(); got != tt.want {
				t.Errorf("Twitter.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSNSバリデート(t *testing.T) {
	type fields struct {
		Discord Discord
		Twitter Twitter
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "どちらもある",
			fields: fields{
				Discord: "あーる#4600",
				Twitter: "@r_sprl",
			},
			want: true,
		},
		{
			name: "Twitterだけある",
			fields: fields{
				Discord: "",
				Twitter: "@r_sprl",
			},
			want: true,
		},
		{
			name: "Discordだけある",
			fields: fields{
				Discord: "",
				Twitter: "@r_sprl",
			},
			want: true,
		},
		{
			name: "両方ない",
			fields: fields{
				Discord: "",
				Twitter: "",
			},
			want: false,
		},
		{
			name: "Discordが不正な値",
			fields: fields{
				Discord: "aaaa#o1",
				Twitter: "@r_sprl",
			},
			want: false,
		},
		{
			name: "Twitterが不正な値",
			fields: fields{
				Discord: "あーる#4600",
				Twitter: "a#example.com",
			},
			want: false,
		},
		{
			name: "どっちも不正な値",
			fields: fields{
				Discord: "あーる@4600",
				Twitter: "a#example.com",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SNS{
				Discord: tt.fields.Discord,
				Twitter: tt.fields.Twitter,
			}
			if got := s.Valid(); got != tt.want {
				t.Errorf("SNS.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComment_Valid(t *testing.T) {
	tests := []struct {
		name string
		n    Comment
		want bool
	}{
		{
			name: "OK",
			n:    "ものすごく頑張ります",
			want: true,
		},
		{
			name: "空文字",
			n:    "",
			want: false,
		},
		{
			name: "長すぎる文字列",
			n: func() Comment {
				var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

				s := make([]rune, 2049)
				for i := range s {
					s[i] = letters[rand.Intn(len(letters))]
				}
				return Comment(s)
			}(),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Valid(); got != tt.want {
				t.Errorf("Comment.Valid() = %v, want %v", got, tt.want)
			}
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
				us: "テスト",
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
