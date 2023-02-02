package game

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

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
			if got := tt.n.Valid(); got != tt.want {
				t.Errorf("Name.Valid() = %v, want %v", got, tt.want)
			}
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
			want: false,
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
			if got := tt.d.Valid(); got != tt.want {
				t.Errorf("Description.Valid() = %v, want %v", got, tt.want)
			}
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
			want: false,
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
			if got := tt.p.Valid(); got != tt.want {
				t.Errorf("Publisher.Valid() = %v, want %v", got, tt.want)
			}
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
			want: false,
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
			if got := tt.d.Valid(); got != tt.want {
				t.Errorf("Developer.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReleaseDate_Time(t *testing.T) {
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
			if got := tt.r.Time(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReleaseDate.Time() = %v, want %v", got, tt.want)
			}
		})
	}
}
