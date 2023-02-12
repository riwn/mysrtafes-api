package game

import (
	"io"
	"mysrtafes-backend/pkg/game"
	"mysrtafes-backend/pkg/game/platform"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestNewGameCreate(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    *game.Game
		want1   []platform.ID
		want2   []game.ID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := NewGameCreate(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGameCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGameCreate() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewGameCreate() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("NewGameCreate() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestNewGameID(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    game.ID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGameID(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGameID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGameID() = %v, want %v", got, tt.want)
			}
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGameFindOption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGameUpdate(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    *game.Game
		want1   []platform.ID
		want2   []game.ID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := NewGameUpdate(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGameUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGameUpdate() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewGameUpdate() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("NewGameUpdate() got2 = %v, want %v", got2, tt.want2)
			}
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
