package tag

import (
	"encoding/json"
	"mysrtafes-backend/pkg/errors"
	"mysrtafes-backend/pkg/game/tag"
	"net/http"
)

type Tag struct {
	Name        tag.Name        `json:"name"`
	Description tag.Description `json:"description"`
}

func NewTagCreate(r *http.Request) (*tag.Tag, error) {
	defer r.Body.Close()

	body := struct {
		Tag
	}{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Request,
			errors.NewInformation(
				errors.ID_JsonDecodeError,
				err.Error(),
				nil,
			),
			"json decode error. bad format request.",
		)
	}

	return tag.New(
		body.Tag.Name,
		body.Tag.Description,
	), nil
}
