package tag

import (
	"context"
	"io"
	"mysrtafes-backend/pkg/game/tag"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestNewTagCreate(t *testing.T) {
	type args struct {
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *tag.Tag
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				method: http.MethodPost,
				url:    "http://example.com",
				body:   strings.NewReader(`{"name": "tag", "description": "desc"}`),
			},
			want: &tag.Tag{
				Name:        "tag",
				Description: "desc",
			},
			wantErr: false,
		},
		{
			name: "decode err",
			args: args{
				method: http.MethodPost,
				url:    "http://example.com",
				body:   strings.NewReader(`{"name": "tag", "description": "desc",}`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.args.method, tt.args.url, tt.args.body)
			got, err := NewTagCreate(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTagCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTagCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTagID(t *testing.T) {
	type args struct {
		method    string
		url       string
		body      io.Reader
		pathParam map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    tag.ID
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				method: http.MethodPost,
				url:    "http://example.com/1",
				body:   strings.NewReader(``),
				pathParam: map[string]string{
					"tagID": "1",
				},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "blank err",
			args: args{
				method: http.MethodPost,
				url:    "http://example.com/1",
				body:   strings.NewReader(``),
				pathParam: map[string]string{
					"tagID": "",
				},
			},
			wantErr: true,
		},
		{
			name: "decode err",
			args: args{
				method: http.MethodPost,
				url:    "http://example.com",
				body:   strings.NewReader(``),
				pathParam: map[string]string{
					"tagID": "jh",
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
			got, err := NewTagID(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTagID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTagID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTagFindOption(t *testing.T) {
	type args struct {
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *tag.FindOption
		wantErr bool
	}{
		{
			name: "no option ok(default)",
			args: args{
				method: http.MethodGet,
				url:    "http://example.com",
				body:   strings.NewReader(``),
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_All,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
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
			got, err := NewTagFindOption(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTagFindOption() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTagFindOption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTagUpdate(t *testing.T) {
	type args struct {
		method    string
		url       string
		body      io.Reader
		pathParam map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    *tag.Tag
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				method: http.MethodPut,
				url:    "http://example.com",
				body:   strings.NewReader(`{"name": "tag", "description": "desc"}`),
				pathParam: map[string]string{
					"tagID": "1",
				},
			},
			want: &tag.Tag{
				ID:          1,
				Name:        "tag",
				Description: "desc",
			},
			wantErr: false,
		},
		{
			name: "OK(through id test)",
			args: args{
				method: http.MethodPut,
				url:    "http://example.com",
				body:   strings.NewReader(`{"id": 3, "name": "tag", "description": "desc"}`),
				pathParam: map[string]string{
					"tagID": "1",
				},
			},
			want: &tag.Tag{
				ID:          1,
				Name:        "tag",
				Description: "desc",
			},
			wantErr: false,
		},
		{
			name: "bad id err",
			args: args{
				method: http.MethodPut,
				url:    "http://example.com",
				body:   strings.NewReader(`{"name": "tag", "description": "desc",}`),
				pathParam: map[string]string{
					"tagID": "a",
				},
			},
			wantErr: true,
		},
		{
			name: "decode err",
			args: args{
				method: http.MethodPut,
				url:    "http://example.com",
				body:   strings.NewReader(`{"name": "tag", "description": "desc",}`),
				pathParam: map[string]string{
					"tagID": "1",
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
			got, err := NewTagUpdate(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTagUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTagUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setOrder(t *testing.T) {
	type args struct {
		findOption *tag.FindOption
		q          url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    *tag.FindOption
		wantErr bool
	}{
		{
			name: "ok no set",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_All,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok desc true, no order",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc": []string{"true"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_All,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  true,
				},
			},
			wantErr: false,
		},
		{
			name: "ok desc true, name order",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc":  []string{"true"},
					"order": []string{"name"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_All,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_Name,
					Desc:  true,
				},
			},
			wantErr: false,
		},
		{
			name: "ok desc false, name order",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc":  []string{"false"},
					"order": []string{"name"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_All,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_Name,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok no desc, id order",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"order": []string{"id"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_All,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok no desc, name order",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"order": []string{"name"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_All,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_Name,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "when seek mode, default option",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_Seek,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc":  []string{"true"},
					"order": []string{"name"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_Seek,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "bad desc error",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc":  []string{"aaa"},
					"order": []string{"name"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_All,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  false,
				},
			},
			wantErr: true,
		},
		{
			name: "bad order error",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc":  []string{"true"},
					"order": []string{"description"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_All,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
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
		findOption *tag.FindOption
		q          url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    *tag.FindOption
		wantErr bool
	}{
		{
			name: "ok no set",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_All,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok seak default",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 100,
						Count:  300,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode": []string{"seek"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_Seek,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok seak with last_id",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  300,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":    []string{"seek"},
					"last_id": []string{"1"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_Seek,
				Seek: tag.Seek{
					LastID: 1,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok seak with count",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  300,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":  []string{"seek"},
					"count": []string{"100"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_Seek,
				Seek: tag.Seek{
					LastID: 0,
					Count:  100,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok seak with both param",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  300,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":    []string{"seek"},
					"last_id": []string{"100"},
					"count":   []string{"500"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_Seek,
				Seek: tag.Seek{
					LastID: 100,
					Count:  500,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "bad last_id error",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":    []string{"seek"},
					"last_id": []string{"aaa"},
					"count":   []string{"500"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_All,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  false,
				},
			},
			wantErr: true,
		},
		{
			name: "bad count error",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":    []string{"seek"},
					"last_id": []string{"1"},
					"count":   []string{"aaa"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_All,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  false,
				},
			},
			wantErr: true,
		},
		{
			name: "ok page default",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  300,
						Offset: 1,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode": []string{"page"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_Pagination,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok page with limit",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  300,
						Offset: 1,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":  []string{"page"},
					"limit": []string{"25"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_Pagination,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  25,
					Offset: 0,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok page with limit",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  300,
						Offset: 1,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":   []string{"page"},
					"offset": []string{"99"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_Pagination,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 99,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok page with both param",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  300,
						Offset: 1,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":   []string{"page"},
					"offset": []string{"999"},
					"limit":  []string{"103"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_Pagination,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  103,
					Offset: 999,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "bad offset error",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 1,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":   []string{"page"},
					"offset": []string{"test"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_All,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 1,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  false,
				},
			},
			wantErr: true,
		},
		{
			name: "bad limit error",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 1,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":  []string{"page"},
					"limit": []string{"test"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_All,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 1,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
					Desc:  false,
				},
			},
			wantErr: true,
		},
		{
			name: "bad mode error",
			args: args{
				findOption: &tag.FindOption{
					SearchMode: tag.SearchMode_All,
					Seek: tag.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: tag.Pagination{
						Limit:  30,
						Offset: 1,
					},
					OrderOption: tag.OrderOption{
						Order: tag.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode": []string{"god"},
				},
			},
			want: &tag.FindOption{
				SearchMode: tag.SearchMode_All,
				Seek: tag.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: tag.Pagination{
					Limit:  30,
					Offset: 1,
				},
				OrderOption: tag.OrderOption{
					Order: tag.Order_ID,
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
