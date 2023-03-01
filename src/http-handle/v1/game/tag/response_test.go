package tag

import (
	"mysrtafes-backend/pkg/game/tag"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tagResponse(t *testing.T) {
	type args struct {
		statusCode int
		msg        string
		tag        *tag.Tag
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "ok",
			args: args{
				statusCode: http.StatusOK,
				msg:        "OKです",
				tag: &tag.Tag{
					ID:          100,
					Name:        "tagです",
					Description: "Tagかも",
				},
			},
			want: TagResponse{
				Code:    http.StatusOK,
				Message: "OKです",
				Data: Tag{
					ID:          100,
					Name:        "tagです",
					Description: "Tagかも",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tagResponse(tt.args.statusCode, tt.args.msg, tt.args.tag))
		})
	}
}

func Test_tagsResponse(t *testing.T) {
	type args struct {
		statusCode int
		msg        string
		tags       []*tag.Tag
		option     *tag.FindOption
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "non page",
			args: args{
				statusCode: http.StatusOK,
				msg:        "non page",
				tags: []*tag.Tag{
					{
						ID:          4,
						Name:        "OK",
						Description: "OKです",
					},
				},
				option: &tag.FindOption{
					SearchMode: tag.SearchMode_Pagination,
					Pagination: tag.Pagination{
						Limit:  100,
						Offset: 1000,
					},
					Seek: tag.Seek{
						LastID: 1,
						Count:  100,
					},
				},
			},
			want: TagsPageResponse{
				Code:    http.StatusOK,
				Message: "non page",
				Data: []Tag{
					{
						ID:          4,
						Name:        "OK",
						Description: "OKです",
					},
				},
				Page: nil,
			},
		},
		{
			name: "page",
			args: args{
				statusCode: http.StatusOK,
				msg:        "page",
				tags: []*tag.Tag{
					{
						ID:          4,
						Name:        "OK",
						Description: "OKです",
					},
				},
				option: &tag.FindOption{
					SearchMode: tag.SearchMode_Pagination,
					Pagination: tag.Pagination{
						Limit:  1,
						Offset: 1000,
					},
					Seek: tag.Seek{
						LastID: 1,
						Count:  100,
					},
				},
			},
			want: TagsPageResponse{
				Code:    http.StatusOK,
				Message: "page",
				Data: []Tag{
					{
						ID:          4,
						Name:        "OK",
						Description: "OKです",
					},
				},
				Page: &Page{
					Limit:  1,
					Offset: 1001,
				},
			},
		},
		{
			name: "non seek",
			args: args{
				statusCode: http.StatusOK,
				msg:        "non seek",
				tags: []*tag.Tag{
					{
						ID:          101,
						Name:        "OK",
						Description: "OKです",
					},
				},
				option: &tag.FindOption{
					SearchMode: tag.SearchMode_Seek,
					Pagination: tag.Pagination{
						Limit:  100,
						Offset: 1000,
					},
					Seek: tag.Seek{
						LastID: 100,
						Count:  2,
					},
				},
			},
			want: TagsNextResponse{
				Code:    http.StatusOK,
				Message: "non seek",
				Data: []Tag{
					{
						ID:          101,
						Name:        "OK",
						Description: "OKです",
					},
				},
				Next: nil,
			},
		},
		{
			name: "seek",
			args: args{
				statusCode: http.StatusOK,
				msg:        "seek",
				tags: []*tag.Tag{
					{
						ID:          101,
						Name:        "OK",
						Description: "OKです",
					},
					{
						ID:          301,
						Name:        "OK",
						Description: "OKです",
					},
				},
				option: &tag.FindOption{
					SearchMode: tag.SearchMode_Seek,
					Pagination: tag.Pagination{
						Limit:  1,
						Offset: 1000,
					},
					Seek: tag.Seek{
						LastID: 1,
						Count:  2,
					},
				},
			},
			want: TagsNextResponse{
				Code:    http.StatusOK,
				Message: "seek",
				Data: []Tag{
					{
						ID:          101,
						Name:        "OK",
						Description: "OKです",
					},
					{
						ID:          301,
						Name:        "OK",
						Description: "OKです",
					},
				},
				Next: &Next{
					LastID: 301,
					Count:  2,
				},
			},
		},
		{
			name: "all",
			args: args{
				statusCode: http.StatusOK,
				msg:        "all",
				tags: []*tag.Tag{
					{
						ID:          4,
						Name:        "OK",
						Description: "OKです",
					},
				},
				option: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Pagination: tag.Pagination{
						Limit:  1,
						Offset: 1000,
					},
					Seek: tag.Seek{
						LastID: 1,
						Count:  100,
					},
				},
			},
			want: TagsResponse{
				Code:    http.StatusOK,
				Message: "all",
				Data: []Tag{
					{
						ID:          4,
						Name:        "OK",
						Description: "OKです",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tagsResponse(tt.args.statusCode, tt.args.msg, tt.args.tags, tt.args.option))
		})
	}
}
