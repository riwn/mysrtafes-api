package handle

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type services struct {
	addr string
}

func NewServices(addr string) services {
	return services{
		addr: addr,
	}
}

func (s services) Server() *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Mount("/mys-challenge", s.mysChallengeRouter())
	r.Mount("/game", s.gameRouter())
	return &http.Server{
		Addr:    s.addr,
		Handler: r,
	}
}

func (s services) mysChallengeRouter() http.Handler {
	r := chi.NewRouter()
	return r
}

func (s services) gameRouter() http.Handler {
	r := chi.NewRouter()
	return r
}
