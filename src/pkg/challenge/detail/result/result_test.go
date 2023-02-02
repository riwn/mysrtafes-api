package result

import (
	"math/rand"
	"testing"
)

func TestComment_Valid(t *testing.T) {
	tests := []struct {
		name string
		n    Comment
		want bool
	}{
		{
			name: "OK",
			n:    "早くクリアできました。",
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
