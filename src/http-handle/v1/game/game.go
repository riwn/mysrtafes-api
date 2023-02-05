package game

import (
	"log"
	"mysrtafes-backend/http-handle/v1/errors"
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

	if err := WriteCreateGame(w, game); err != nil {
		log.Println(err)
		errors.WriteError(w, err)
	}
}

func (h *gameHandler) read(w http.ResponseWriter, r *http.Request) {
	gameID, err := NewGameID(r)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	// IDがないときは複数検索にする
	if !gameID.Valid() {
		h.find(w, r)
		return
	}

	game, err := h.server.Read(gameID)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	if err := WriteReadGame(w, game); err != nil {
		log.Println(err)
		errors.WriteError(w, err)
	}
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

	if err := WriteFindGame(w, games, findOption); err != nil {
		log.Println(err)
		errors.WriteError(w, err)
	}
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

	if err := WriteUpdateGame(w, game); err != nil {
		log.Println(err)
		errors.WriteError(w, err)
	}
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

	if err := WriteDeleteGame(w, gameID); err != nil {
		log.Println(err)
		errors.WriteError(w, err)
	}
}
