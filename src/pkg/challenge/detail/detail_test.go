package detail

import (
	"math/rand"
	"testing"
)

func Test部門_バリデート(t *testing.T) {
	tests := []struct {
		name string
		d    Department
		want bool
	}{
		{
			name: "ちょっと不思議部門を指定",
			d:    Department_BEGINNER,
			want: true,
		},
		{
			name: "もっと不思議部門を指定",
			d:    Department_EXPERT,
			want: true,
		},
		{
			name: "存在しない部門を指定",
			d:    3,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Valid(); got != tt.want {
				t.Errorf("Department.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test部門文字列(t *testing.T) {
	tests := []struct {
		name string
		d    Department
		want string
	}{
		{
			name: "ちょっと不思議部門のラベル出力",
			d:    Department_BEGINNER,
			want: "ちょっと不思議部門",
		},
		{
			name: "もっと不思議部門のラベル出力",
			d:    Department_EXPERT,
			want: "もっと不思議部門",
		},
		{
			name: "存在しない部門のラベル出力",
			d:    3,
			want: "謎の部門",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("Department.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoalDetail_Valid(t *testing.T) {
	tests := []struct {
		name string
		n    GoalDetail
		want bool
	}{
		{
			name: "OK",
			n:    "すごく早くクリアを目指します",
			want: true,
		},
		{
			name: "空文字",
			n:    "",
			want: false,
		},
		{
			name: "長すぎる文字列",
			n: func() GoalDetail {
				var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

				s := make([]rune, 2049)
				for i := range s {
					s[i] = letters[rand.Intn(len(letters))]
				}
				return GoalDetail(s)
			}(),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Valid(); got != tt.want {
				t.Errorf("GoalDetail.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
