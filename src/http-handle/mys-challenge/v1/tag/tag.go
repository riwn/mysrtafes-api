package tag

import (
	"log"
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
	default:
		http.NotFound(w, r)
	}
}

func (h *tagHandler) create(w http.ResponseWriter, r *http.Request) {
	tag, err := NewTagCreate(r)
	if err != nil {
		log.Println(err)
		// TODO: errResponse
		return
	}

	tag, err = h.server.Create(tag)
	if err != nil {
		log.Println(err)
		// TODO: errResponse
		return
	}

	if err := WriteCreateTag(w, tag); err != nil {
		log.Println(err)
		// TODO: errResponse
	}
}
func (h *tagHandler) read(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}