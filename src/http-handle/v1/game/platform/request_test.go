package platform

import (
	"context"
	"io"
	"mysrtafes-backend/pkg/game/platform"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestNewPlatformCreate(t *testing.T) {
	type args struct {
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *platform.Platform
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				method: http.MethodPost,
				url:    "http://example.com",
				body:   strings.NewReader(`{"name": "platform", "description": "desc"}`),
			},
			want: &platform.Platform{
				Name:        "platform",
				Description: "desc",
			},
			wantErr: false,
		},
		{
			name: "decode err",
			args: args{
				method: http.MethodPost,
				url:    "http://example.com",
				body:   strings.NewReader(`{"name": "platform", "description": "desc",}`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.args.method, tt.args.url, tt.args.body)
			got, err := NewPlatformCreate(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlatformCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlatformCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPlatformID(t *testing.T) {
	type args struct {
		method    string
		url       string
		body      io.Reader
		pathParam map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    platform.ID
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				method: http.MethodGet,
				url:    "http://example.com/1",
				body:   strings.NewReader(``),
				pathParam: map[string]string{
					"platformID": "1",
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
					"platformID": "",
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
					"platformID": "jh",
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

			got, err := NewPlatformID(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlatformID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlatformID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPlatformFindOption(t *testing.T) {
	type args struct {
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *platform.FindOption
		wantErr bool
	}{
		{
			name: "no option ok(default)",
			args: args{
				method: http.MethodGet,
				url:    "http://example.com",
				body:   strings.NewReader(``),
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_All,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
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
			got, err := NewPlatformFindOption(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlatformFindOption() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlatformFindOption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPlatformUpdate(t *testing.T) {
	type args struct {
		method    string
		url       string
		body      io.Reader
		pathParam map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    *platform.Platform
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				method: http.MethodPut,
				url:    "http://example.com",
				body:   strings.NewReader(`{"name": "platform", "description": "desc"}`),
				pathParam: map[string]string{
					"platformID": "1",
				},
			},
			want: &platform.Platform{
				ID:          1,
				Name:        "platform",
				Description: "desc",
			},
			wantErr: false,
		},
		{
			name: "OK(through id test)",
			args: args{
				method: http.MethodPut,
				url:    "http://example.com",
				body:   strings.NewReader(`{"id": 3, "name": "platform", "description": "desc"}`),
				pathParam: map[string]string{
					"platformID": "1",
				},
			},
			want: &platform.Platform{
				ID:          1,
				Name:        "platform",
				Description: "desc",
			},
			wantErr: false,
		},
		{
			name: "bad id err",
			args: args{
				method: http.MethodPut,
				url:    "http://example.com",
				body:   strings.NewReader(`{"name": "platform", "description": "desc",}`),
				pathParam: map[string]string{
					"platformID": "a",
				},
			},
			wantErr: true,
		},
		{
			name: "decode err",
			args: args{
				method: http.MethodPut,
				url:    "http://example.com",
				body:   strings.NewReader(`{"name": "platform", "description": "desc",}`),
				pathParam: map[string]string{
					"platformID": "1",
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
			got, err := NewPlatformUpdate(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlatformUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlatformUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setOrder(t *testing.T) {
	type args struct {
		findOption *platform.FindOption
		q          url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    *platform.FindOption
		wantErr bool
	}{
		{
			name: "ok no set",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_All,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok desc true, no order",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc": []string{"true"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_All,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  true,
				},
			},
			wantErr: false,
		},
		{
			name: "ok desc true, name order",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc":  []string{"true"},
					"order": []string{"name"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_All,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_Name,
					Desc:  true,
				},
			},
			wantErr: false,
		},
		{
			name: "ok desc false, name order",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc":  []string{"false"},
					"order": []string{"name"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_All,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_Name,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok no desc, id order",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"order": []string{"id"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_All,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok no desc, name order",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"order": []string{"name"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_All,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_Name,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "when seek mode, default option",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_Seek,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc":  []string{"true"},
					"order": []string{"name"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_Seek,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "bad desc error",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc":  []string{"aaa"},
					"order": []string{"name"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_All,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  false,
				},
			},
			wantErr: true,
		},
		{
			name: "bad order error",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"desc":  []string{"true"},
					"order": []string{"description"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_All,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
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
				return
			}

			if !reflect.DeepEqual(tt.args.findOption, tt.want) {
				t.Errorf("setOrder() = %v, want %v", tt.args.findOption, tt.want)
			}
		})
	}
}

func Test_setSearchMode(t *testing.T) {
	type args struct {
		findOption *platform.FindOption
		q          url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    *platform.FindOption
		wantErr bool
	}{
		{
			name: "ok no set",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_All,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok seak default",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 100,
						Count:  300,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode": []string{"seek"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_Seek,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok seak with last_id",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  300,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":    []string{"seek"},
					"last_id": []string{"1"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_Seek,
				Seek: platform.Seek{
					LastID: 1,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok seak with count",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  300,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":  []string{"seek"},
					"count": []string{"100"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_Seek,
				Seek: platform.Seek{
					LastID: 0,
					Count:  100,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok seak with both param",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  300,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":    []string{"seek"},
					"last_id": []string{"100"},
					"count":   []string{"500"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_Seek,
				Seek: platform.Seek{
					LastID: 100,
					Count:  500,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "bad last_id error",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":    []string{"seek"},
					"last_id": []string{"aaa"},
					"count":   []string{"500"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_All,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  false,
				},
			},
			wantErr: true,
		},
		{
			name: "bad count error",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 0,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":    []string{"seek"},
					"last_id": []string{"1"},
					"count":   []string{"aaa"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_All,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  false,
				},
			},
			wantErr: true,
		},
		{
			name: "ok page default",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  300,
						Offset: 1,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode": []string{"page"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_Pagination,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok page with limit",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  300,
						Offset: 1,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":  []string{"page"},
					"limit": []string{"25"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_Pagination,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  25,
					Offset: 0,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok page with limit",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  300,
						Offset: 1,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":   []string{"page"},
					"offset": []string{"99"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_Pagination,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 99,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "ok page with both param",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  300,
						Offset: 1,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":   []string{"page"},
					"offset": []string{"999"},
					"limit":  []string{"103"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_Pagination,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  103,
					Offset: 999,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "bad offset error",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 1,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":   []string{"page"},
					"offset": []string{"test"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_All,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 1,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  false,
				},
			},
			wantErr: true,
		},
		{
			name: "bad limit error",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 1,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode":  []string{"page"},
					"limit": []string{"test"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_All,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 1,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
					Desc:  false,
				},
			},
			wantErr: true,
		},
		{
			name: "bad mode error",
			args: args{
				findOption: &platform.FindOption{
					SearchMode: platform.SearchMode_All,
					Seek: platform.Seek{
						LastID: 0,
						Count:  30,
					},
					Pagination: platform.Pagination{
						Limit:  30,
						Offset: 1,
					},
					OrderOption: platform.OrderOption{
						Order: platform.Order_ID,
						Desc:  false,
					},
				},
				q: url.Values{
					"mode": []string{"god"},
				},
			},
			want: &platform.FindOption{
				SearchMode: platform.SearchMode_All,
				Seek: platform.Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: platform.Pagination{
					Limit:  30,
					Offset: 1,
				},
				OrderOption: platform.OrderOption{
					Order: platform.Order_ID,
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

			if !reflect.DeepEqual(tt.args.findOption, tt.want) {
				t.Errorf("setSearchMode() = %v, want %v", tt.args.findOption, tt.want)
			}
		})
	}
}
