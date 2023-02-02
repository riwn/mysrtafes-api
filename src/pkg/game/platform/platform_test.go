package platform

import (
	"math/rand"
	"testing"
)

func TestName_Valid(t *testing.T) {
	tests := []struct {
		name string
		n    Name
		want bool
	}{
		{
			name: "OK",
			n:    "Nintendo64",
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
			d:    "古めのレトロハードです。",
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
