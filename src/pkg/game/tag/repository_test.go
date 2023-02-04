package tag

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

type repository struct {
	tag *Tag
	err error
	// flags
	create, read, update, delete bool
}

func (r repository) TagCreate(*Tag) (*Tag, error) {
	if r.create {
		return r.tag, r.err
	}
	return nil, fmt.Errorf("failed create")
}
func (r repository) TagRead(ID) (*Tag, error) {
	if r.read {
		return r.tag, r.err
	}
	return nil, fmt.Errorf("failed read")
}
func (r repository) TagUpdate(*Tag) (*Tag, error) {
	if r.update {
		return r.tag, r.err
	}
	return nil, fmt.Errorf("failed update")
}
func (r repository) TagDelete(ID) error {
	if r.delete {
		return r.err
	}
	return fmt.Errorf("failed delete")
}

func Test_server_Create(t *testing.T) {
	type fields struct {
		repository Repository
	}
	type args struct {
		t *Tag
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Tag
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				repository: repository{
					tag: &Tag{
						ID:          1,
						Name:        "OK",
						Description: "OKですよ",
					},
					create: true,
				},
			},
			args: args{
				t: &Tag{
					Name:        "OK",
					Description: "OKですよ",
				},
			},
			want: &Tag{
				ID:          1,
				Name:        "OK",
				Description: "OKですよ",
			},
		},
		{
			name: "名前のバリデートエラー",
			fields: fields{
				repository: repository{
					tag:    nil,
					err:    nil,
					create: true,
				},
			},
			args: args{
				t: &Tag{
					Name:        "",
					Description: "NGですよ",
				},
			},
			wantErr: true,
		},
		{
			name: "説明のバリデートエラー",
			fields: fields{
				repository: repository{
					tag:    nil,
					err:    nil,
					create: true,
				},
			},
			args: args{
				t: &Tag{
					Name: "NG",
					Description: func() Description {
						var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
						s := make([]rune, 2049)
						for i := range s {
							s[i] = letters[rand.Intn(len(letters))]
						}
						return Description(s)
					}(),
				},
			},
			wantErr: true,
		},
		{
			name: "repositoryのエラー",
			fields: fields{
				repository: repository{
					tag:    nil,
					err:    fmt.Errorf("create error"),
					create: true,
				},
			},
			args: args{
				t: &Tag{
					Name:        "NG",
					Description: "NGですよ",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				repository: tt.fields.repository,
			}
			got, err := s.Create(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
