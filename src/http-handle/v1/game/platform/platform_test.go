package platform

import (
	"context"
	"fmt"
	"io"
	"mysrtafes-backend/pkg/game/platform"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

type server struct {
	Platform  *platform.Platform
	Platforms []*platform.Platform
	err       error
	// flags
	create, read, find, update, delete bool
}

func (s *server) Create(*platform.Platform) (*platform.Platform, error) {
	if s.create {
		return s.Platform, s.err
	}
	return nil, fmt.Errorf("failed create")
}
func (s *server) Read(platform.ID) (*platform.Platform, error) {
	if s.read {
		return s.Platform, s.err
	}
	return nil, fmt.Errorf("failed read")
}
func (s *server) Find(*platform.FindOption) ([]*platform.Platform, error) {
	if s.find {
		return s.Platforms, s.err
	}
	return nil, fmt.Errorf("failed find")
}
func (s *server) Update(*platform.Platform) (*platform.Platform, error) {
	if s.update {
		return s.Platform, s.err
	}
	return nil, fmt.Errorf("failed update")
}
func (s *server) Delete(platform.ID) error {
	if s.delete {
		return s.err
	}
	return fmt.Errorf("failed delete")
}

func TestNewPlatformHandler(t *testing.T) {
	type args struct {
		s platform.Server
	}
	tests := []struct {
		name string
		args args
		want *platformHandler
	}{
		{
			name: "ok",
			args: args{
				s: &server{
					err: fmt.Errorf("failed delete"),
				},
			},
			want: &platformHandler{
				server: &server{
					err: fmt.Errorf("failed delete"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, NewPlatformHandler(tt.args.s), tt.want)
		})
	}
}

func Test_platformHandler_HandlePlatform(t *testing.T) {
	type fields struct {
		server platform.Server
	}
	type args struct {
		w         *httptest.ResponseRecorder
		method    string
		url       string
		body      io.Reader
		pathParam map[string]string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
		wantBody       string
	}{
		{
			name: "Get OK",
			fields: fields{
				server: &server{
					Platform: &platform.Platform{
						ID:          1,
						Name:        "OK",
						Description: "OKです",
					},
					read: true,
				},
			},
			args: args{
				w:      httptest.NewRecorder(),
				method: http.MethodGet,
				url:    "http://example.com/1",
				body:   strings.NewReader(``),
				pathParam: map[string]string{
					"platformID": "1",
				},
			},
			wantStatusCode: http.StatusOK,
			// TODO: ここのBodyの比較が微妙すぎるので要検討
			wantBody: `{"code":200,"message":"success read platform","data":{"id":1,"name":"OK","description":"OKです","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &platformHandler{
				server: tt.fields.server,
			}
			ctx := chi.NewRouteContext()
			for key, val := range tt.args.pathParam {
				ctx.URLParams.Add(key, val)
			}
			r := httptest.NewRequest(tt.args.method, tt.args.url, tt.args.body)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
			h.HandlePlatform(tt.args.w, r)
			if !assert.Equal(t, tt.wantStatusCode, tt.args.w.Code) {
				return
			}
			if !assert.Equal(t, tt.wantBody, tt.args.w.Body.String()) {
				return
			}
		})
	}
}

func Test_platformHandler_HandlePlatformForMultiple(t *testing.T) {
	type fields struct {
		server platform.Server
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &platformHandler{
				server: tt.fields.server,
			}
			h.HandlePlatformForMultiple(tt.args.w, tt.args.r)
		})
	}
}
