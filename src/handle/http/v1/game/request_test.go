package game

import (
	"context"
	"io"
	"mysrtafes-backend/pkg/game"
	"mysrtafes-backend/pkg/game/platform"
	"mysrtafes-backend/pkg/game/tag"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestNewGameCreate(t *testing.T) {
	type args struct {
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *game.Game
		want1   []platform.ID
		want2   []tag.ID
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				method: http.MethodPost,
				url:    "http://example.com",
				body: strings.NewReader(`{
                    "name": "TestGame",
                    "description": "desc",
                    "publisher":"Nintendo",
                    "developer":"Chu Soft",
                    "release_date": "1997-01-22",
                    "links": [
                        {
                            "title": "wiki",
                            "url": "http://example.com",
                            "description":"色々知れます"
                        },
                        {
                            "title": "wiki2",
                            "url": "http://example.com2",
                            "description":"色々知れます2"
                        }
                    ],
                    "tag_ids": [1,2,3,4],
                    "platform_ids": [2,3,4]
                }`),
			},
			want: &game.Game{
				Name:        "TestGame",
				Description: "desc",
				Publisher:   "Nintendo",
				Developer:   "Chu Soft",
				ReleaseDate: game.ReleaseDate(time.Date(1997, 1, 22, 0, 0, 0, 0, time.UTC)),
				Links: []*game.Link{
					{
						Title: "wiki",
						URL: func() game.URL {
							g, _ := game.NewURL("http://example.com")
							return g
						}(),
						LinkDescription: "色々知れます",
					},
					{
						Title: "wiki2",
						URL: func() game.URL {
							g, _ := game.NewURL("http://example.com2")
							return g
						}(),
						LinkDescription: "色々知れます2",
					},
				},
			},
			want1:   []platform.ID{2, 3, 4},
			want2:   []tag.ID{1, 2, 3, 4},
			wantErr: false,
		},
		{
			name: "release date decode error",
			args: args{
				method: http.MethodPost,
				url:    "http://example.com",
				body: strings.NewReader(`{
                    "name": "TestGame",
                    "description": "desc",
                    "publisher":"Nintendo",
                    "developer":"Chu Soft",
                    "release_date": "1997/01/22",
                    "links": [
                        {
                            "title": "wiki",
                            "url": "http://example.com",
                            "description":"色々知れます"
                        },
                        {
                            "title": "wiki2",
                            "url": "http://example.com2",
                            "description":"色々知れます2"
                        }
                    ],
                    "tag_ids": [1,2,3,4],
                    "platform_ids": [2,3,4]
                }`),
			},
			wantErr: true,
		},
		{
			name: "url decode error",
			args: args{
				method: http.MethodPost,
				url:    "http://example.com",
				body: strings.NewReader(`{
                    "name": "TestGame",
                    "description": "desc",
                    "publisher":"Nintendo",
                    "developer":"Chu Soft",
                    "release_date": "1997-01-22",
                    "links": [
                        {
                            "title": "wiki",
                            "url": "ht",
                            "description":"色々知れます"
                        },
                        {
                            "title": "wiki2",
                            "url": "http://example.com2",
                            "description":"色々知れます2"
                        }
                    ],
                    "tag_ids": [1,2,3,4],
                    "platform_ids": [2,3,4]
                }`),
			},
			wantErr: true,
		},
		{
			name: "decode err",
			args: args{
				method: http.MethodPost,
				url:    "http://example.com",
				body: strings.NewReader(`{
                    "name": "TestGame",
                    "description": "desc",
                    "publisher":"Nintendo",
                    "developer":"Chu Soft",
                    "release_date": "1997-01-22",
                    "links": [
                        {
                            "title": "wiki",
                            "url": "http://example.com",
                            "description":"色々知れます"
                        },
                        {
                            "title": "wiki2",
                            "url": "http://example.com2",
                            "description":"色々知れます2"
                        },
                        "tag_ids": [1,2,3,4],
                        "platform_ids": [2,3,4],
                }`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.args.method, tt.args.url, tt.args.body)
			got, got1, got2, err := NewGameCreate(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGameCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
			assert.Equal(t, tt.want2, got2)
		})
	}
}

func TestNewGameID(t *testing.T) {
	type args struct {
		method    string
		url       string
		body      io.Reader
		pathParam map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    game.ID
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				method: http.MethodGet,
				url:    "http://example.com/1",
				body:   strings.NewReader(``),
				pathParam: map[string]string{
					"gameID": "1",
				},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "blank err",
			args: args{
				method: http.MethodGet,
				url:    "http://example.com",
				body:   strings.NewReader(``),
				pathParam: map[string]string{
					"gameID": "",
				},
			},
			wantErr: true,
		},
		{
			name: "decode err",
			args: args{
				method: http.MethodGet,
				url:    "http://example.com",
				body:   strings.NewReader(``),
				pathParam: map[string]string{
					"gameID": "jh",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := chi.NewRouteContext()
			for key, val := range tt.args.pathParam {
				ctx.URLParams.Add(key, val)
			}
			r := httptest.NewRequest(tt.args.method, tt.args.url, tt.args.body)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

			got, err := NewGameID(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGameID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewGameFindOption(t *testing.T) {
	type args struct {
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *game.FindOption
		wantErr bool
	}{
		{
			name: "no option ok(default)",
			args: args{
				method: http.MethodGet,
				url:    "http://example.com",
				body:   strings.NewReader(``),
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_All,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
		},
		{
			name: "search mode parse error",
			args: args{
				method: http.MethodGet,
				url:    "http://example.com?mode=seek&last_id=aaa",
				body:   strings.NewReader(``),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "order parse error",
			args: args{
				method: http.MethodGet,
				url:    "http://example.com?desc=test",
				body:   strings.NewReader(``),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.args.method, tt.args.url, tt.args.body)
			got, err := NewGameFindOption(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGameFindOption() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewGameUpdate(t *testing.T) {
	type args struct {
		method    string
		url       string
		body      io.Reader
		pathParam map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    *game.Game
		want1   []platform.ID
		want2   []tag.ID
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				method: http.MethodPost,
				url:    "http://example.com",
				body: strings.NewReader(`{
                    "id": 100,
                    "name": "TestGame",
                    "description": "desc",
                    "publisher":"Nintendo",
                    "developer":"Chu Soft",
                    "release_date": "1997-01-22",
                    "links": [
                        {
                            "title": "wiki",
                            "url": "http://example.com",
                            "description":"色々知れます"
                        },
                        {
                            "title": "wiki2",
                            "url": "http://example.com2",
                            "description":"色々知れます2"
                        }
                    ],
                    "tag_ids": [1,2,3,4],
                    "platform_ids": [2,3,4]
                }`),
				pathParam: map[string]string{
					"gameID": "1",
				},
			},
			want: &game.Game{
				ID:          1,
				Name:        "TestGame",
				Description: "desc",
				Publisher:   "Nintendo",
				Developer:   "Chu Soft",
				ReleaseDate: game.ReleaseDate(time.Date(1997, 1, 22, 0, 0, 0, 0, time.UTC)),
				Links: []*game.Link{
					{
						Title: "wiki",
						URL: func() game.URL {
							g, _ := game.NewURL("http://example.com")
							return g
						}(),
						LinkDescription: "色々知れます",
					},
					{
						Title: "wiki2",
						URL: func() game.URL {
							g, _ := game.NewURL("http://example.com2")
							return g
						}(),
						LinkDescription: "色々知れます2",
					},
				},
			},
			want1:   []platform.ID{2, 3, 4},
			want2:   []tag.ID{1, 2, 3, 4},
			wantErr: false,
		},
		{
			name: "bad id error",
			args: args{
				method: http.MethodPost,
				url:    "http://example.com",
				body: strings.NewReader(`{
                    "name": "TestGame",
                    "description": "desc",
                    "publisher":"Nintendo",
                    "developer":"Chu Soft",
                    "release_date": "1997-01-22",
                    "links": [
                        {
                            "title": "wiki",
                            "url": "http://example.com",
                            "description":"色々知れます"
                        },
                        {
                            "title": "wiki2",
                            "url": "http://example.com2",
                            "description":"色々知れます2"
                        }
                    ],
                    "tag_ids": [1,2,3,4],
                    "platform_ids": [2,3,4]
                }`),
				pathParam: map[string]string{
					"gameID": "test",
				},
			},
			wantErr: true,
		},
		{
			name: "release date decode error",
			args: args{
				method: http.MethodPost,
				url:    "http://example.com",
				body: strings.NewReader(`{
                    "name": "TestGame",
                    "description": "desc",
                    "publisher":"Nintendo",
                    "developer":"Chu Soft",
                    "release_date": "1997/01/22",
                    "links": [
                        {
                            "title": "wiki",
                            "url": "http://example.com",
                            "description":"色々知れます"
                        },
                        {
                            "title": "wiki2",
                            "url": "http://example.com2",
                            "description":"色々知れます2"
                        }
                    ],
                    "tag_ids": [1,2,3,4],
                    "platform_ids": [2,3,4]
                }`),
				pathParam: map[string]string{
					"gameID": "1",
				},
			},
			wantErr: true,
		},
		{
			name: "url decode error",
			args: args{
				method: http.MethodPost,
				url:    "http://example.com",
				body: strings.NewReader(`{
                    "name": "TestGame",
                    "description": "desc",
                    "publisher":"Nintendo",
                    "developer":"Chu Soft",
                    "release_date": "1997-01-22",
                    "links": [
                        {
                            "title": "wiki",
                            "url": "ht",
                            "description":"色々知れます"
                        },
                        {
                            "title": "wiki2",
                            "url": "http://example.com2",
                            "description":"色々知れます2"
                        }
                    ],
                    "tag_ids": [1,2,3,4],
                    "platform_ids": [2,3,4]
                }`),
				pathParam: map[string]string{
					"gameID": "1",
				},
			},
			wantErr: true,
		},
		{
			name: "decode err",
			args: args{
				method: http.MethodPost,
				url:    "http://example.com",
				body: strings.NewReader(`{
                    "name": "TestGame",
                    "description": "desc",
                    "publisher":"Nintendo",
                    "developer":"Chu Soft",
                    "release_date": "1997-01-22",
                    "links": [
                        {
                            "title": "wiki",
                            "url": "http://example.com",
                            "description":"色々知れます"
                        },
                        {
                            "title": "wiki2",
                            "url": "http://example.com2",
                            "description":"色々知れます2"
                        },
                        "tag_ids": [1,2,3,4],
                        "platform_ids": [2,3,4],
                }`),
				pathParam: map[string]string{
					"gameID": "1",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := chi.NewRouteContext()
			for key, val := range tt.args.pathParam {
				ctx.URLParams.Add(key, val)
			}
			r := httptest.NewRequest(tt.args.method, tt.args.url, tt.args.body)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
			got, got1, got2, err := NewGameUpdate(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGameUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
			assert.Equal(t, tt.want2, got2)
		})
	}
}

func Test_setOrder(t *testing.T) {
	type args struct {
		findOption *game.FindOption
		q          url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    *game.FindOption
		wantErr bool
	}{
		{
			name: "ok no set",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_All,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok desc true, no order",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc": []string{"true"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_All,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  true,
				},
			},
			wantErr: false,
		},
		{
			name: "ok desc true, name order",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc":  []string{"true"},
					"order": []string{"name"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_All,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_Name,
					Desc:  true,
				},
			},
			wantErr: false,
		},
		{
			name: "ok desc false, name order",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc":  []string{"false"},
					"order": []string{"name"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_All,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_Name,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok no desc, id order",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"order": []string{"id"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_All,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok no desc, name order",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"order": []string{"name"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_All,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_Name,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "when seek mode, default option",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_Seek,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc":  []string{"true"},
					"order": []string{"name"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_Seek,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "bad desc error",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc":  []string{"aaa"},
					"order": []string{"name"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_All,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: true,
		},
		{
			name: "bad order error",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc":  []string{"true"},
					"order": []string{"description"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_All,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := setOrder(tt.args.findOption, tt.args.q); (err != nil) != tt.wantErr {
				t.Errorf("setOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_setSearchMode(t *testing.T) {
	type args struct {
		findOption *game.FindOption
		q          url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    *game.FindOption
		wantErr bool
	}{
		{
			name: "ok no set",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_All,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok seak default",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 100,
						Count:  300,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode": []string{"seek"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_Seek,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok seak with last_id",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  300,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":    []string{"seek"},
					"last_id": []string{"1"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_Seek,
				Seek: game.Seek{
					LastID: 1,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok seak with count",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  300,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":  []string{"seek"},
					"count": []string{"100"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_Seek,
				Seek: game.Seek{
					LastID: 0,
					Count:  100,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok seak with both param",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  300,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":    []string{"seek"},
					"last_id": []string{"100"},
					"count":   []string{"500"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_Seek,
				Seek: game.Seek{
					LastID: 100,
					Count:  500,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "bad last_id error",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":    []string{"seek"},
					"last_id": []string{"aaa"},
					"count":   []string{"500"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_All,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: true,
		},
		{
			name: "bad count error",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":    []string{"seek"},
					"last_id": []string{"1"},
					"count":   []string{"aaa"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_All,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: true,
		},
		{
			name: "ok page default",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  300,
						Offset: 1,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode": []string{"page"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_Pagination,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok page with limit",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  300,
						Offset: 1,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":  []string{"page"},
					"limit": []string{"25"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_Pagination,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  25,
					Offset: 0,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok page with limit",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  300,
						Offset: 1,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":   []string{"page"},
					"offset": []string{"99"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_Pagination,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 99,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok page with both param",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  300,
						Offset: 1,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":   []string{"page"},
					"offset": []string{"999"},
					"limit":  []string{"103"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_Pagination,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  103,
					Offset: 999,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "bad offset error",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 1,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":   []string{"page"},
					"offset": []string{"test"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_All,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 1,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: true,
		},
		{
			name: "bad limit error",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 1,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":  []string{"page"},
					"limit": []string{"test"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_All,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 1,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: true,
		},
		{
			name: "bad mode error",
			args: args{
				findOption: &game.FindOption{
					SearchMode: game.SearchMode_All,
					Seek: game.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: game.Pagination{
						Limit:  30,
						Offset: 1,
					},
					OrderOption: game.OrderOption{
						Order: game.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode": []string{"god"},
				},
			},
			want: &game.FindOption{
				SearchMode: game.SearchMode_All,
				Seek: game.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: game.Pagination{
					Limit:  30,
					Offset: 1,
				},
				OrderOption: game.OrderOption{
					Order: game.Order_ID,
					Desc:  false,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := setSearchMode(tt.args.findOption, tt.args.q); (err != nil) != tt.wantErr {
				t.Errorf("setSearchMode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
