package game

import (
	"log"
	"mysrtafes-backend/handle/http/v1/errors"
	"mysrtafes-backend/pkg/game"
	"net/http"
)

type gameHandler struct {
	server game.Server
}

func NewGameHandler(s game.Server) *gameHandler {
	return &gameHandler{s}
}

func (h *gameHandler) HandleGame(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.read(w, r)
	case http.MethodPost:
		h.create(w, r)
	case http.MethodPut:
		h.update(w, r)
	case http.MethodDelete:
		h.delete(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *gameHandler) HandleGameForMultiple(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.find(w, r)
	// TODO: 必要であれば複数登録や更新削除を作る
	// case http.MethodPost:
	// case http.MethodPut:
	// case http.MethodDelete:
	default:
		http.NotFound(w, r)
	}
}

func (h *gameHandler) create(w http.ResponseWriter, r *http.Request) {
	game, platformIDs, tagIDs, err := NewGameCreate(r)
	if err != nil {
		// TODO: logの改善(トレーサーなど)
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	game, err = h.server.Create(game, platformIDs, tagIDs)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	WriteCreateGame(w, game)
}

func (h *gameHandler) read(w http.ResponseWriter, r *http.Request) {
	gameID, err := NewGameID(r)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	game, err := h.server.Read(gameID)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	WriteReadGame(w, game)
}

func (h *gameHandler) find(w http.ResponseWriter, r *http.Request) {
	findOption, err := NewGameFindOption(r)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	games, err := h.server.Find(findOption)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	WriteFindGame(w, games, findOption)
}

func (h *gameHandler) update(w http.ResponseWriter, r *http.Request) {
	game, platformIDs, tagIDs, err := NewGameUpdate(r)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	game, err = h.server.Update(game, platformIDs, tagIDs)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	WriteUpdateGame(w, game)
}

func (h *gameHandler) delete(w http.ResponseWriter, r *http.Request) {
	gameID, err := NewGameID(r)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	err = h.server.Delete(gameID)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	WriteDeleteGame(w, gameID)
}
