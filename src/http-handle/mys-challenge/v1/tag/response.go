package tag

import (
	"encoding/json"
	"mysrtafes-backend/pkg/game/tag"
	"net/http"
	"time"
)

// TODO: Laravelに合わせてるのであまりにも過剰(単純な移行なため仕様合わせでこうなってる。)
// TODO: あと、呼び出し側のことを考えるとjsonはキャメルケースのほうがよさそう。
type TagResponse struct {
	ID          tag.ID          `json:"id"`
	Name        tag.Name        `json:"name"`
	Description tag.Description `json:"description"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

func WriteCreateTag(w http.ResponseWriter, tag *tag.Tag) error {
	body := struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    TagResponse `json:"data"`
	}{
		Code:    http.StatusCreated,
		Message: "success create tag",
		Data: TagResponse{
			ID:          tag.ID,
			Name:        tag.Name,
			Description: tag.Description,
			CreatedAt:   tag.CreatedAt,
			UpdatedAt:   tag.UpdatedAt,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(&body)
}
