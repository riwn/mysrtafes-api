package tag

import (
	"log"
	"mysrtafes-backend/handle/http/v1/errors"
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
	case http.MethodDelete:
		h.delete(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *tagHandler) HandleTagForMultiple(w http.ResponseWriter, r *http.Request) {
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

	WriteCreateTag(w, tag)
}

func (h *tagHandler) read(w http.ResponseWriter, r *http.Request) {
	tagID, err := NewTagID(r)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	tag, err := h.server.Read(tagID)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	WriteReadTag(w, tag)
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

	WriteFindTag(w, tags, findOption)
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

	WriteUpdateTag(w, tag)
}

func (h *tagHandler) delete(w http.ResponseWriter, r *http.Request) {
	tagID, err := NewTagID(r)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	err = h.server.Delete(tagID)
	if err != nil {
		log.Println(err)
		errors.WriteError(w, err)
		return
	}

	WriteDeleteTag(w, tagID)
}
