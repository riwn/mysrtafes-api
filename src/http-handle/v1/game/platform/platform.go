package platform

import (
	"log"
	"mysrtafes-backend/http-handle/v1/errors"
	"mysrtafes-backend/pkg/game/platform"
	"net/http"
)

type platformHandler struct {
	server platform.Server
}

func NewPlatformHandler(s platform.Server) *platformHandler {
	return &platformHandler{s}
}

func (h *platformHandler) HandlePlatform(w http.ResponseWriter, r *http.Request) {
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

func (h *platformHandler) create(w http.ResponseWriter, r *http.Request) {
	platform, err := NewPlatformCreate(r)
	if err != nil {
		// TODO: logの改善(トレーサーなど)
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	platform, err = h.server.Create(platform)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	if err := WriteCreatePlatform(w, platform); err != nil {
		log.Println(err)
		errors.WriteError(w, err)
	}
}

func (h *platformHandler) read(w http.ResponseWriter, r *http.Request) {
	platformID, err := NewPlatformID(r)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	// IDがないときは複数検索にする
	if !platformID.Valid() {
		h.find(w, r)
		return
	}

	platform, err := h.server.Read(platformID)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	if err := WriteReadPlatform(w, platform); err != nil {
		log.Println(err)
		errors.WriteError(w, err)
	}
}

func (h *platformHandler) find(w http.ResponseWriter, r *http.Request) {
	findOption, err := NewPlatformFindOption(r)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	platforms, err := h.server.Find(findOption)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	if err := WriteFindPlatform(w, platforms, findOption); err != nil {
		log.Println(err)
		errors.WriteError(w, err)
	}
}

func (h *platformHandler) update(w http.ResponseWriter, r *http.Request) {
	platform, err := NewPlatformUpdate(r)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	platform, err = h.server.Update(platform)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	if err := WriteUpdatePlatform(w, platform); err != nil {
		log.Println(err)
		errors.WriteError(w, err)
	}
}

func (h *platformHandler) delete(w http.ResponseWriter, r *http.Request) {
	platformID, err := NewPlatformID(r)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	err = h.server.Delete(platformID)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	if err := WriteDeletePlatform(w, platformID); err != nil {
		log.Println(err)
		errors.WriteError(w, err)
	}
}