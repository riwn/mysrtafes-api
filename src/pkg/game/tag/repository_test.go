package tag

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

type repository struct {
	tag  *Tag
	tags []*Tag
	err  error
	// flags
	create, read, find, update, delete bool
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

func (r repository) TagFind(*FindOption) ([]*Tag, error) {
	if r.find {
		return r.tags, r.err
	}
	return nil, fmt.Errorf("failed find")
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

func TestNewServer(t *testing.T) {
	type args struct {
		repo Repository
	}
	tests := []struct {
		name string
		args args
		want Server
	}{
		{
			name: "new",
			args: args{
				repo: repository{},
			},
			want: &server{
				repository: repository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
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

func Test_server_Read(t *testing.T) {
	type fields struct {
		repository Repository
	}
	type args struct {
		id ID
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
					read: true,
				},
			},
			args: args{
				id: 1,
			},
			want: &Tag{
				ID:          1,
				Name:        "OK",
				Description: "OKですよ",
			},
		},
		{
			name: "idのバリデートエラー",
			fields: fields{
				repository: repository{
					tag: &Tag{
						ID:          0,
						Name:        "OK",
						Description: "OKですよ",
					},
					read: true,
				},
			},
			args: args{
				id: 0,
			},
			wantErr: true,
		},
		{
			name: "repositoryのエラー",
			fields: fields{
				repository: repository{
					tag:  nil,
					err:  fmt.Errorf("read error"),
					read: true,
				},
			},
			args: args{
				id: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				repository: tt.fields.repository,
			}
			got, err := s.Read(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_Find(t *testing.T) {
	type fields struct {
		repository Repository
	}
	type args struct {
		f *FindOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Tag
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				repository: repository{
					tags: []*Tag{
						{
							ID:          1,
							Name:        "OK",
							Description: "OKですよ",
						},
						{
							ID:          2,
							Name:        "OK2",
							Description: "OK2ですよ",
						},
					},
					find: true,
				},
			},
			args: args{
				f: nil,
			},
			want: []*Tag{
				{
					ID:          1,
					Name:        "OK",
					Description: "OKですよ",
				},
				{
					ID:          2,
					Name:        "OK2",
					Description: "OK2ですよ",
				},
			},
		},
		{
			name: "repositoryのエラー",
			fields: fields{
				repository: repository{
					tag:  nil,
					err:  fmt.Errorf("find error"),
					find: true,
				},
			},
			args: args{
				f: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				repository: tt.fields.repository,
			}
			got, err := s.Find(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_Update(t *testing.T) {
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
					update: true,
				},
			},
			args: args{
				t: &Tag{
					ID:          1,
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
			name: "idのバリデートエラー",
			fields: fields{
				repository: repository{
					tag: &Tag{
						ID:          0,
						Name:        "OK",
						Description: "OKですよ",
					},
					update: true,
				},
			},
			args: args{
				t: &Tag{
					ID:          0,
					Name:        "OK",
					Description: "OKですよ",
				},
			},
			wantErr: true,
		},
		{
			name: "Nameのバリデートエラー",
			fields: fields{
				repository: repository{
					tag: &Tag{
						ID:          1,
						Name:        "NG",
						Description: "NGですよ",
					},
					update: true,
				},
			},
			args: args{
				t: &Tag{
					ID:          1,
					Name:        "",
					Description: "NGですよ",
				},
			},
			wantErr: true,
		},
		{
			name: "Descriptionのバリデートエラー",
			fields: fields{
				repository: repository{
					tag: &Tag{
						ID:          1,
						Name:        "NG",
						Description: "NGですよ",
					},
					update: true,
				},
			},
			args: args{
				t: &Tag{
					ID:   1,
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
					err:    fmt.Errorf("read error"),
					update: true,
				},
			},
			args: args{
				t: &Tag{
					Name:        "OK",
					Description: "OKですよ",
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
			got, err := s.Update(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_Delete(t *testing.T) {
	type fields struct {
		repository Repository
	}
	type args struct {
		id ID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
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
					delete: true,
				},
			},
			args: args{
				id: 1,
			},
		},
		{
			name: "idのバリデートエラー",
			fields: fields{
				repository: repository{
					tag: &Tag{
						ID:          0,
						Name:        "OK",
						Description: "OKですよ",
					},
					delete: true,
				},
			},
			args: args{
				id: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				repository: tt.fields.repository,
			}
			if err := s.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("server.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
