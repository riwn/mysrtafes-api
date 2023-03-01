package handle

import (
	v1Game "mysrtafes-backend/handle/http/v1/game"
	v1Platform "mysrtafes-backend/handle/http/v1/game/platform"
	v1Tag "mysrtafes-backend/handle/http/v1/game/tag"
	v1Challenge "mysrtafes-backend/handle/http/v1/mystery-challenge2/challenge"
	"mysrtafes-backend/pkg/challenge"
	"mysrtafes-backend/pkg/game"
	"mysrtafes-backend/pkg/game/platform"
	"mysrtafes-backend/pkg/game/tag"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type services struct {
	addr      string
	Game      game.Server
	Challenge challenge.Server
	Tag       tag.Server
	Platform  platform.Server
	// TODO: HandleをもつServiceの追加
}

func NewServices(addr string, game game.Server, challenge challenge.Server, tag tag.Server, platform platform.Server) services {
	return services{addr, game, challenge, tag, platform}
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
	// /api/v1/mystery-challenge2
	r.Mount("/mystery-challenge2", s.mysChallengeRouter())
	// /api/v1/games
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
	// /api/v1/games/tags
	r.Mount("/tags", s.tagRouter())
	// /api/v1/games/platforms
	r.Mount("/platforms", s.platformRouter())

	gameHandler := v1Game.NewGameHandler(s.Game)
	// 複数操作
	r.Get("/", gameHandler.HandleGameForMultiple)
	// 単体操作
	r.Get("/{gameID}", gameHandler.HandleGame)
	r.Post("/", gameHandler.HandleGame)
	r.Put("/{gameID}", gameHandler.HandleGame)
	r.Delete("/{gameID}", gameHandler.HandleGame)
	return r
}

func (s services) tagRouter() http.Handler {
	r := chi.NewRouter()
	tagHandler := v1Tag.NewTagHandler(s.Tag)
	// 複数操作
	r.Get("/", tagHandler.HandleTagForMultiple)
	// 単体操作
	r.Get("/{tagID}", tagHandler.HandleTag)
	r.Post("/", tagHandler.HandleTag)
	r.Put("/{tagID}", tagHandler.HandleTag)
	r.Delete("/{tagID}", tagHandler.HandleTag)
	return r
}

func (s services) platformRouter() http.Handler {
	r := chi.NewRouter()
	platformHandler := v1Platform.NewPlatformHandler(s.Platform)
	// 複数操作
	r.Get("/", platformHandler.HandlePlatformForMultiple)
	// 単体操作
	r.Get("/{platformID}", platformHandler.HandlePlatform)
	r.Post("/", platformHandler.HandlePlatform)
	r.Put("/{platformID}", platformHandler.HandlePlatform)
	r.Delete("/{platformID}", platformHandler.HandlePlatform)
	return r
}
