package link

import (
	"math/rand"
	"testing"
)

func TestTitle_Valid(t *testing.T) {
	tests := []struct {
		name string
		tr   Title
		want bool
	}{
		{
			name: "OK",
			tr:   "不思議RTAフェス",
			want: true,
		},
		{
			name: "空文字",
			tr:   "",
			want: false,
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
			if got := tt.tr.Valid(); got != tt.want {
				t.Errorf("Title.Valid() = %v, want %v", got, tt.want)
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
			d:    "不思議RTAフェスの公式サイトです",
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
