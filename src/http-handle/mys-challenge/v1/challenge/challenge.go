package challenge

import (
	"log"
	"mysrtafes-backend/pkg/challenge"
	"net/http"
)

type challengeHandler struct {
	server challenge.Server
}

func NewChallengeHandler(s challenge.Server) *challengeHandler {
	return &challengeHandler{s}
}

func (h *challengeHandler) HandleChallenge(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.read(w, r)
	case http.MethodPost:
		h.create(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *challengeHandler) create(w http.ResponseWriter, r *http.Request) {
	challenge, err := NewChallengeCreate(r)
	if err != nil {
		log.Println(err)
		// TODO: errResponse
		return
	}

	challenge, err = h.server.Create(challenge)
	if err != nil {
		log.Println(err)
		// TODO: errResponse
		return
	}

	if err := WriteCreateChallenge(w, challenge); err != nil {
		log.Println(err)
		// TODO: errResponse
	}
}
func (h *challengeHandler) read(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
