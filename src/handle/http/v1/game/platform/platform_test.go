package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mysrtafes-backend/pkg/errors"
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
						Name:        "Get OK",
						Description: "Get OKです",
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
			wantBody: func() string {
				body := platformResponse(
					http.StatusOK,
					"success read platform",
					&platform.Platform{
						ID:          1,
						Name:        "Get OK",
						Description: "Get OKです",
					},
				)
				str, _ := json.Marshal(body)
				return string(str)
			}(),
		},
		{
			name: "Post OK",
			fields: fields{
				server: &server{
					Platform: &platform.Platform{
						ID:          2,
						Name:        "Post OK",
						Description: "Post OKです",
					},
					create: true,
				},
			},
			args: args{
				w:         httptest.NewRecorder(),
				method:    http.MethodPost,
				url:       "http://example.com",
				body:      strings.NewReader(`{"name": "Post OK", "description": "Post OKです"}`),
				pathParam: map[string]string{},
			},
			wantStatusCode: http.StatusCreated,
			wantBody: func() string {
				body := platformResponse(
					http.StatusCreated,
					"success create platform",
					&platform.Platform{
						ID:          2,
						Name:        "Post OK",
						Description: "Post OKです",
					},
				)
				str, _ := json.Marshal(body)
				return string(str)
			}(),
		},
		{
			name: "Put OK",
			fields: fields{
				server: &server{
					Platform: &platform.Platform{
						ID:          3,
						Name:        "Put OK",
						Description: "Put OKです",
					},
					update: true,
				},
			},
			args: args{
				w:      httptest.NewRecorder(),
				method: http.MethodPut,
				url:    "http://example.com/3",
				body:   strings.NewReader(`{"name": "Put OK", "description": "Put OKです"}`),
				pathParam: map[string]string{
					"platformID": "3",
				},
			},
			wantStatusCode: http.StatusOK,
			wantBody: func() string {
				body := platformResponse(
					http.StatusOK,
					"success update platform",
					&platform.Platform{
						ID:          3,
						Name:        "Put OK",
						Description: "Put OKです",
					},
				)
				str, _ := json.Marshal(body)
				return string(str)
			}(),
		},
		{
			name: "Delete OK",
			fields: fields{
				server: &server{
					Platform: &platform.Platform{
						ID:          4,
						Name:        "Delete OK",
						Description: "Delete OKです",
					},
					delete: true,
				},
			},
			args: args{
				w:      httptest.NewRecorder(),
				method: http.MethodDelete,
				url:    "http://example.com/4",
				body:   strings.NewReader(`{"name": "Delete OK", "description": "Delete OKです"}`),
				pathParam: map[string]string{
					"platformID": "4",
				},
			},
			wantStatusCode: http.StatusOK,
			wantBody: func() string {
				body := deletePlatformResponse(4)
				str, _ := json.Marshal(body)
				return string(str)
			}(),
		},
		{
			name: "Bad Method NG",
			fields: fields{
				server: &server{
					Platform: &platform.Platform{
						ID:          4,
						Name:        "Delete OK",
						Description: "Delete OKです",
					},
					delete: true,
				},
			},
			args: args{
				w:      httptest.NewRecorder(),
				method: http.MethodPatch,
				url:    "http://example.com/4",
				body:   strings.NewReader(`{"name": "Delete OK", "description": "Delete OKです"}`),
				pathParam: map[string]string{
					"platformID": "4",
				},
			},
			wantStatusCode: http.StatusNotFound,
			wantBody:       "",
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
			// NOTE: BodyのStringは\nが入る仕様らしいので削除
			if tt.wantBody != "" && !assert.Equal(t, tt.wantBody, strings.Replace(tt.args.w.Body.String(), "\n", "", -1)) {
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
			name: "Find OK",
			fields: fields{
				server: &server{
					Platforms: []*platform.Platform{
						{
							ID:          5,
							Name:        "Find1 OK",
							Description: "Find1 OKです",
						},
						{
							ID:          6,
							Name:        "Find2 OK",
							Description: "Find2 OKです",
						},
						{
							ID:          7,
							Name:        "Find3 OK",
							Description: "Find3 OKです",
						},
					},
					find: true,
				},
			},
			args: args{
				w:         httptest.NewRecorder(),
				method:    http.MethodGet,
				url:       "http://example.com",
				body:      strings.NewReader(``),
				pathParam: map[string]string{},
			},
			wantStatusCode: http.StatusOK,
			wantBody: func() string {
				body := platformsResponse(
					http.StatusOK,
					"success find platform",
					[]*platform.Platform{
						{
							ID:          5,
							Name:        "Find1 OK",
							Description: "Find1 OKです",
						},
						{
							ID:          6,
							Name:        "Find2 OK",
							Description: "Find2 OKです",
						},
						{
							ID:          7,
							Name:        "Find3 OK",
							Description: "Find3 OKです",
						},
					},
					platform.NewFindOption(),
				)
				str, _ := json.Marshal(body)
				return string(str)
			}(),
		},
		{
			name: "Bad Method NG",
			fields: fields{
				server: &server{
					Platforms: []*platform.Platform{
						{
							ID:          5,
							Name:        "Find1 OK",
							Description: "Find1 OKです",
						},
						{
							ID:          6,
							Name:        "Find2 OK",
							Description: "Find2 OKです",
						},
						{
							ID:          7,
							Name:        "Find3 OK",
							Description: "Find3 OKです",
						},
					},
					find: true,
				},
			},
			args: args{
				w:         httptest.NewRecorder(),
				method:    http.MethodDelete,
				url:       "http://example.com",
				body:      strings.NewReader(``),
				pathParam: map[string]string{},
			},
			wantStatusCode: http.StatusNotFound,
			wantBody:       "",
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
			h.HandlePlatformForMultiple(tt.args.w, r)
			if !assert.Equal(t, tt.wantStatusCode, tt.args.w.Code) {
				return
			}
			// NOTE: BodyのStringは\nが入る仕様らしいので削除
			if tt.wantBody != "" && !assert.Equal(t, tt.wantBody, strings.Replace(tt.args.w.Body.String(), "\n", "", -1)) {
				return
			}
		})
	}
}

func Test_platformHandler_create(t *testing.T) {
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
			name: "OK",
			fields: fields{
				server: &server{
					Platform: &platform.Platform{
						ID:          1,
						Name:        "OK",
						Description: "OKです",
					},
					create: true,
				},
			},
			args: args{
				w:         httptest.NewRecorder(),
				method:    http.MethodPost,
				url:       "http://example.com",
				body:      strings.NewReader(`{"name": "OK", "description": "OKです"}`),
				pathParam: map[string]string{},
			},
			wantStatusCode: http.StatusCreated,
			wantBody: func() string {
				body := platformResponse(
					http.StatusCreated,
					"success create platform",
					&platform.Platform{
						ID:          1,
						Name:        "OK",
						Description: "OKです",
					},
				)
				str, _ := json.Marshal(body)
				return string(str)
			}(),
		},
		{
			name: "new request NG",
			fields: fields{
				server: &server{
					Platform: &platform.Platform{
						ID:          1,
						Name:        "NG",
						Description: "NGです",
					},
					create: true,
				},
			},
			args: args{
				w:         httptest.NewRecorder(),
				method:    http.MethodPost,
				url:       "http://example.com",
				body:      strings.NewReader(`{"name": "OK", "description": "OKです",}`),
				pathParam: map[string]string{},
			},
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "",
		},
		{
			name: "Create NG",
			fields: fields{
				server: &server{
					Platform: &platform.Platform{
						ID:          1,
						Name:        "NG",
						Description: "NGです",
					},
					err:    errors.NewUnauthorized(errors.Layer_Model, nil, "error"),
					create: true,
				},
			},
			args: args{
				w:         httptest.NewRecorder(),
				method:    http.MethodPost,
				url:       "http://example.com",
				body:      strings.NewReader(`{"name": "OK", "description": "OKです"}`),
				pathParam: map[string]string{},
			},
			wantStatusCode: http.StatusUnauthorized,
			wantBody:       "",
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
			h.create(tt.args.w, r)
			if !assert.Equal(t, tt.wantStatusCode, tt.args.w.Code) {
				return
			}
			// NOTE: BodyのStringは\nが入る仕様らしいので削除
			if tt.wantBody != "" && !assert.Equal(t, tt.wantBody, strings.Replace(tt.args.w.Body.String(), "\n", "", -1)) {
				return
			}
		})
	}
}

func Test_platformHandler_read(t *testing.T) {
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
			name: "OK",
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
			wantBody: func() string {
				body := platformResponse(
					http.StatusOK,
					"success read platform",
					&platform.Platform{
						ID:          1,
						Name:        "OK",
						Description: "OKです",
					},
				)
				str, _ := json.Marshal(body)
				return string(str)
			}(),
		},
		{
			name: "new request NG",
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
				url:    "http://example.com/aa",
				body:   strings.NewReader(``),
				pathParam: map[string]string{
					"platformID": "aa",
				},
			},
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "",
		},
		{
			name: "read NG",
			fields: fields{
				server: &server{
					Platform: &platform.Platform{
						ID:          1,
						Name:        "OK",
						Description: "OKです",
					},
					err:  errors.NewForbidden(errors.Layer_Model, nil, "error"),
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
			wantStatusCode: http.StatusForbidden,
			wantBody:       "",
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
			h.read(tt.args.w, r)
			if !assert.Equal(t, tt.wantStatusCode, tt.args.w.Code) {
				return
			}
			// NOTE: BodyのStringは\nが入る仕様らしいので削除
			if tt.wantBody != "" && !assert.Equal(t, tt.wantBody, strings.Replace(tt.args.w.Body.String(), "\n", "", -1)) {
				return
			}
		})
	}
}

func Test_platformHandler_find(t *testing.T) {
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
			name: "Find OK",
			fields: fields{
				server: &server{
					Platforms: []*platform.Platform{
						{
							ID:          5,
							Name:        "Find1 OK",
							Description: "Find1 OKです",
						},
						{
							ID:          6,
							Name:        "Find2 OK",
							Description: "Find2 OKです",
						},
						{
							ID:          7,
							Name:        "Find3 OK",
							Description: "Find3 OKです",
						},
					},
					find: true,
				},
			},
			args: args{
				w:         httptest.NewRecorder(),
				method:    http.MethodGet,
				url:       "http://example.com",
				body:      strings.NewReader(``),
				pathParam: map[string]string{},
			},
			wantStatusCode: http.StatusOK,
			wantBody: func() string {
				body := platformsResponse(
					http.StatusOK,
					"success find platform",
					[]*platform.Platform{
						{
							ID:          5,
							Name:        "Find1 OK",
							Description: "Find1 OKです",
						},
						{
							ID:          6,
							Name:        "Find2 OK",
							Description: "Find2 OKです",
						},
						{
							ID:          7,
							Name:        "Find3 OK",
							Description: "Find3 OKです",
						},
					},
					platform.NewFindOption(),
				)
				str, _ := json.Marshal(body)
				return string(str)
			}(),
		},
		{
			name: "new request NG",
			fields: fields{
				server: &server{
					Platforms: []*platform.Platform{
						{
							ID:          5,
							Name:        "Find1 OK",
							Description: "Find1 OKです",
						},
						{
							ID:          6,
							Name:        "Find2 OK",
							Description: "Find2 OKです",
						},
						{
							ID:          7,
							Name:        "Find3 OK",
							Description: "Find3 OKです",
						},
					},
					find: true,
				},
			},
			args: args{
				w:      httptest.NewRecorder(),
				method: http.MethodGet,
				url:    "http://example.com?mode=aaaa",
				body:   strings.NewReader(``),
				pathParam: map[string]string{
					"mode": "aaaa",
				},
			},
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "",
		},
		{
			name: "find NG",
			fields: fields{
				server: &server{
					Platforms: []*platform.Platform{
						{
							ID:          5,
							Name:        "Find1 OK",
							Description: "Find1 OKです",
						},
						{
							ID:          6,
							Name:        "Find2 OK",
							Description: "Find2 OKです",
						},
						{
							ID:          7,
							Name:        "Find3 OK",
							Description: "Find3 OKです",
						},
					},
					err:  errors.NewInternalServerError(errors.Layer_Model, nil, "error"),
					find: true,
				},
			},
			args: args{
				w:         httptest.NewRecorder(),
				method:    http.MethodGet,
				url:       "http://example.com",
				body:      strings.NewReader(``),
				pathParam: map[string]string{},
			},
			wantStatusCode: http.StatusInternalServerError,
			wantBody:       "",
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
			h.find(tt.args.w, r)
			if !assert.Equal(t, tt.wantStatusCode, tt.args.w.Code) {
				return
			}
			// NOTE: BodyのStringは\nが入る仕様らしいので削除
			if tt.wantBody != "" && !assert.Equal(t, tt.wantBody, strings.Replace(tt.args.w.Body.String(), "\n", "", -1)) {
				return
			}
		})
	}
}

func Test_platformHandler_update(t *testing.T) {
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
			name: "OK",
			fields: fields{
				server: &server{
					Platform: &platform.Platform{
						ID:          3,
						Name:        "OK",
						Description: "OKです",
					},
					update: true,
				},
			},
			args: args{
				w:      httptest.NewRecorder(),
				method: http.MethodPut,
				url:    "http://example.com/3",
				body:   strings.NewReader(`{"name": "OK", "description": "OKです"}`),
				pathParam: map[string]string{
					"platformID": "3",
				},
			},
			wantStatusCode: http.StatusOK,
			wantBody: func() string {
				body := platformResponse(
					http.StatusOK,
					"success update platform",
					&platform.Platform{
						ID:          3,
						Name:        "OK",
						Description: "OKです",
					},
				)
				str, _ := json.Marshal(body)
				return string(str)
			}(),
		},
		{
			name: "new platform NG",
			fields: fields{
				server: &server{
					Platform: &platform.Platform{
						ID:          3,
						Name:        "OK",
						Description: "OKです",
					},
					update: true,
				},
			},
			args: args{
				w:      httptest.NewRecorder(),
				method: http.MethodPut,
				url:    "http://example.com/3",
				body:   strings.NewReader(`{"name": "OK", "description": "OKです",}`),
				pathParam: map[string]string{
					"platformID": "3",
				},
			},
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "",
		},
		{
			name: "update err",
			fields: fields{
				server: &server{
					Platform: &platform.Platform{
						ID:          3,
						Name:        "OK",
						Description: "OKです",
					},
					err:    errors.NewUnsupportedMediaType(errors.Layer_Model, nil, "error"),
					update: true,
				},
			},
			args: args{
				w:      httptest.NewRecorder(),
				method: http.MethodPut,
				url:    "http://example.com/3",
				body:   strings.NewReader(`{"name": "OK", "description": "OKです"}`),
				pathParam: map[string]string{
					"platformID": "3",
				},
			},
			wantStatusCode: http.StatusUnsupportedMediaType,
			wantBody:       "",
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
			h.update(tt.args.w, r)
			if !assert.Equal(t, tt.wantStatusCode, tt.args.w.Code) {
				return
			}
			// NOTE: BodyのStringは\nが入る仕様らしいので削除
			if tt.wantBody != "" && !assert.Equal(t, tt.wantBody, strings.Replace(tt.args.w.Body.String(), "\n", "", -1)) {
				return
			}
		})
	}
}

func Test_platformHandler_delete(t *testing.T) {
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
			name: "OK",
			fields: fields{
				server: &server{
					Platform: &platform.Platform{
						ID:          4,
						Name:        "OK",
						Description: "OKです",
					},
					delete: true,
				},
			},
			args: args{
				w:      httptest.NewRecorder(),
				method: http.MethodDelete,
				url:    "http://example.com/4",
				body:   strings.NewReader(`{"name": "OK", "description": "OKです"}`),
				pathParam: map[string]string{
					"platformID": "4",
				},
			},
			wantStatusCode: http.StatusOK,
			wantBody: func() string {
				body := deletePlatformResponse(4)
				str, _ := json.Marshal(body)
				return string(str)
			}(),
		},
		{
			name: "new id err",
			fields: fields{
				server: &server{
					Platform: &platform.Platform{
						ID:          4,
						Name:        "OK",
						Description: "OKです",
					},
					delete: true,
				},
			},
			args: args{
				w:      httptest.NewRecorder(),
				method: http.MethodDelete,
				url:    "http://example.com/aaaa",
				body:   strings.NewReader(`{"name": "OK", "description": "OKです"}`),
				pathParam: map[string]string{
					"platformID": "aaaa",
				},
			},
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "",
		},
		{
			name: "delete err",
			fields: fields{
				server: &server{
					Platform: &platform.Platform{
						ID:          1,
						Name:        "OK",
						Description: "OKです",
					},
					err:    errors.NewUnauthorized(errors.Layer_Model, nil, "error"),
					delete: true,
				},
			},
			args: args{
				w:      httptest.NewRecorder(),
				method: http.MethodDelete,
				url:    "http://example.com/1",
				body:   strings.NewReader(`{"name": "OK", "description": "OKです"}`),
				pathParam: map[string]string{
					"platformID": "1",
				},
			},
			wantStatusCode: http.StatusUnauthorized,
			wantBody:       "",
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
			h.delete(tt.args.w, r)
			if !assert.Equal(t, tt.wantStatusCode, tt.args.w.Code) {
				return
			}
			// NOTE: BodyのStringは\nが入る仕様らしいので削除
			if tt.wantBody != "" && !assert.Equal(t, tt.wantBody, strings.Replace(tt.args.w.Body.String(), "\n", "", -1)) {
				return
			}
		})
	}
}
