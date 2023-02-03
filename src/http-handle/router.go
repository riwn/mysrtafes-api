package handle

import (
	v1 "mysrtafes-backend/http-handle/mys-challenge/v1"
	"mysrtafes-backend/pkg/challenge"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type services struct {
	addr      string
	Challenge challenge.Server
	// TODO: HandleをもつServiceの追加
}

func NewServices(addr string, challenge challenge.Server) services {
	return services{addr, challenge}
}

func (s services) Server() *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/api/v1/mystery-challenge2", s.mysChallengeRouter())
	r.Mount("/game", s.gameRouter())
	return &http.Server{
		Addr:    s.addr,
		Handler: r,
	}
}

func (s services) mysChallengeRouter() http.Handler {
	r := chi.NewRouter()
	challengeHandler := v1.NewChallengeHandler(s.Challenge)
	r.Post("/challenge", challengeHandler.HandleChallenge)
	return r
}

func (s services) gameRouter() http.Handler {
	r := chi.NewRouter()
	// TODO: Routerの追加
	return r
}
