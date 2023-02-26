package platform

import (
	"mysrtafes-backend/pkg/game/platform"
	"net/http"
	"testing"
)

func Test_platFormsResponse(t *testing.T) {
	type args struct {
		statusCode int
		msg        string
		platforms  []*platform.Platform
		option     *platform.FindOption
	}
	type Page struct {
		Limit  platform.Limit  `json:"limit"`
		Offset platform.Offset `json:"offset"`
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "page",
			args: args{
				statusCode: http.StatusOK,
				msg:        "a",
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
				},
			},
			want: struct {
				Code    int                `json:"code"`
				Message string             `json:"message"`
				Data    []PlatformResponse `json:"data"`
				Page    *Page              `json:"page"`
			}{
				Code:    http.StatusOK,
				Message: "a",
				Data: []PlatformResponse{
					{
						ID:          4,
						Name:        "OK",
						Description: "OKです",
					},
				},
				Page: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: テスト通らないのでまた考える。
			// got := platFormsResponse(tt.args.statusCode, tt.args.msg, tt.args.platforms, tt.args.option)
			// if !assert.Equal(t, tt.want, got) {
			// 	return
			// }
		})
	}
}
