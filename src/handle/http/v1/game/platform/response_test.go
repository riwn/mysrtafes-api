package platform

import (
	"mysrtafes-backend/pkg/game/platform"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_platformsResponse(t *testing.T) {
	type args struct {
		statusCode int
		msg        string
		platforms  []*platform.Platform
		option     *platform.FindOption
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
				platforms: []*platform.Platform{
					{
						ID:          4,
						Name:        "OK",
						Description: "OKです",
					},
				},
				option: &platform.FindOption{
					SearchMode: platform.SearchMode_Pagination,
					Pagination: platform.Pagination{
						Limit:  100,
						Offset: 1000,
					},
					Seek: platform.Seek{
						LastID: 1,
						Count:  100,
					},
				},
			},
			want: PlatformsPageResponse{
				Code:    http.StatusOK,
				Message: "non page",
				Data: []Platform{
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
				platforms: []*platform.Platform{
					{
						ID:          4,
						Name:        "OK",
						Description: "OKです",
					},
				},
				option: &platform.FindOption{
					SearchMode: platform.SearchMode_Pagination,
					Pagination: platform.Pagination{
						Limit:  1,
						Offset: 1000,
					},
					Seek: platform.Seek{
						LastID: 1,
						Count:  100,
					},
				},
			},
			want: PlatformsPageResponse{
				Code:    http.StatusOK,
				Message: "page",
				Data: []Platform{
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
				platforms: []*platform.Platform{
					{
						ID:          101,
						Name:        "OK",
						Description: "OKです",
					},
				},
				option: &platform.FindOption{
					SearchMode: platform.SearchMode_Seek,
					Pagination: platform.Pagination{
						Limit:  100,
						Offset: 1000,
					},
					Seek: platform.Seek{
						LastID: 100,
						Count:  2,
					},
				},
			},
			want: PlatformsNextResponse{
				Code:    http.StatusOK,
				Message: "non seek",
				Data: []Platform{
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
				platforms: []*platform.Platform{
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
				option: &platform.FindOption{
					SearchMode: platform.SearchMode_Seek,
					Pagination: platform.Pagination{
						Limit:  1,
						Offset: 1000,
					},
					Seek: platform.Seek{
						LastID: 1,
						Count:  2,
					},
				},
			},
			want: PlatformsNextResponse{
				Code:    http.StatusOK,
				Message: "seek",
				Data: []Platform{
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
				platforms: []*platform.Platform{
					{
						ID:          4,
						Name:        "OK",
						Description: "OKです",
					},
				},
				option: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Pagination: platform.Pagination{
						Limit:  1,
						Offset: 1000,
					},
					Seek: platform.Seek{
						LastID: 1,
						Count:  100,
					},
				},
			},
			want: PlatformsResponse{
				Code:    http.StatusOK,
				Message: "all",
				Data: []Platform{
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
			got := platformsResponse(tt.args.statusCode, tt.args.msg, tt.args.platforms, tt.args.option)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}
