package handle

import (
	v1Tag "mysrtafes-backend/http-handle/v1/game/tag"
	v1Challenge "mysrtafes-backend/http-handle/v1/mystery-challenge2/challenge"
	"mysrtafes-backend/pkg/challenge"
	"mysrtafes-backend/pkg/game/tag"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type services struct {
	addr      string
	Challenge challenge.Server
	Tag       tag.Server
	// TODO: HandleをもつServiceの追加
}

func NewServices(addr string, challenge challenge.Server, tag tag.Server) services {
	return services{addr, challenge, tag}
}

func (s services) Server() *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/api/v1", s.apiV1Router())
	return &http.Server{
		Addr:    s.addr,
		Handler: r,
	}
}

func (s services) apiV1Router() http.Handler {
	r := chi.NewRouter()
	r.Mount("/mystery-challenge2", s.mysChallengeRouter())
	r.Mount("/games", s.gameRouter())
	return r
}

func (s services) mysChallengeRouter() http.Handler {
	r := chi.NewRouter()
	challengeHandler := v1Challenge.NewChallengeHandler(s.Challenge)
	r.Post("/challenges", challengeHandler.HandleChallenge)
	return r
}

func (s services) gameRouter() http.Handler {
	r := chi.NewRouter()
	r.Mount("/tags", s.tagRouter())
	// TODO: Routerの追加
	return r
}

func (s services) tagRouter() http.Handler {
	r := chi.NewRouter()
	tagHandler := v1Tag.NewTagHandler(s.Tag)
	r.Post("/", tagHandler.HandleTag)
	return r
}
