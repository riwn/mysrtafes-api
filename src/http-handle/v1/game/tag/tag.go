package tag

import (
	"log"
	"mysrtafes-backend/http-handle/v1/errors"
	"mysrtafes-backend/pkg/game/tag"
	"net/http"
)

type tagHandler struct {
	server tag.Server
}

func NewTagHandler(s tag.Server) *tagHandler {
	return &tagHandler{s}
}

func (h *tagHandler) HandleTag(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.read(w, r)
	case http.MethodPost:
		h.create(w, r)
	case http.MethodPut:
		h.update(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *tagHandler) create(w http.ResponseWriter, r *http.Request) {
	tag, err := NewTagCreate(r)
	if err != nil {
		// TODO: logの改善(トレーサーなど)
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	tag, err = h.server.Create(tag)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	if err := WriteCreateTag(w, tag); err != nil {
		log.Println(err)
		errors.WriteError(w, err)
	}
}

func (h *tagHandler) read(w http.ResponseWriter, r *http.Request) {
	tagID, err := NewTagRead(r)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	// IDがないときは複数検索にする
	if !tagID.Valid() {
		h.find(w, r)
		return
	}

	tag, err := h.server.Read(tagID)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	if err := WriteReadTag(w, tag); err != nil {
		log.Println(err)
		errors.WriteError(w, err)
	}
}

func (h *tagHandler) find(w http.ResponseWriter, r *http.Request) {
	findOption, err := NewTagFindOption(r)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	tags, err := h.server.Find(findOption)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	if err := WriteFindTag(w, tags, findOption); err != nil {
		log.Println(err)
		errors.WriteError(w, err)
	}
}

func (h *tagHandler) update(w http.ResponseWriter, r *http.Request) {
	tag, err := NewTagUpdate(r)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	tag, err = h.server.Update(tag)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	if err := WriteUpdateTag(w, tag); err != nil {
		log.Println(err)
		errors.WriteError(w, err)
	}
}
