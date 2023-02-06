package platform

import (
	"math/rand"
	"reflect"
	"testing"
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
			if got := tt.i.Valid(); got != tt.want {
				t.Errorf("ID.Valid() = %v, want %v", got, tt.want)
			}
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
			if got := tt.d.Valid(); got != tt.want {
				t.Errorf("Description.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		name        Name
		description Description
	}
	tests := []struct {
		name string
		args args
		want *Platform
	}{
		{
			name: "OK",
			args: args{
				name:        "Nintendo64",
				description: "もうレトロゲームと化しました",
			},
			want: &Platform{
				Name:        "Nintendo64",
				Description: "もうレトロゲームと化しました",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.name, tt.args.description); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithID(t *testing.T) {
	type args struct {
		id          ID
		name        Name
		description Description
	}
	tests := []struct {
		name string
		args args
		want *Platform
	}{
		{
			name: "OK",
			args: args{
				id:          9853,
				name:        "PS2",
				description: "ギリギリレトロゲームだとは思ってないハードです。",
			},
			want: &Platform{
				ID:          9853,
				Name:        "PS2",
				Description: "ギリギリレトロゲームだとは思ってないハードです。",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWithID(tt.args.id, tt.args.name, tt.args.description); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithID() = %v, want %v", got, tt.want)
			}
		})
	}
}
