package tag

import (
	"encoding/json"
	"mysrtafes-backend/pkg/game/tag"
	"net/http"
	"time"
)

type TagResponse struct {
	ID          tag.ID          `json:"id"`
	Name        tag.Name        `json:"name"`
	Description tag.Description `json:"description"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// write create response for tag
func WriteCreateTag(w http.ResponseWriter, tag *tag.Tag) error {
	return writeTag(w, http.StatusCreated, "success create tag", tag)
}

// write read response for tag
func WriteReadTag(w http.ResponseWriter, tag *tag.Tag) error {
	return writeTag(w, http.StatusOK, "success read tag", tag)
}

// write update response for tag
func WriteUpdateTag(w http.ResponseWriter, tag *tag.Tag) error {
	return writeTag(w, http.StatusOK, "success update tag", tag)
}

// write find response for tag
func WriteFindTag(w http.ResponseWriter, tags []*tag.Tag, option *tag.FindOption) error {
	return writeTags(w, http.StatusOK, "success find tag", tags, option)
}

func writeTag(w http.ResponseWriter, statusCode int, msg string, tag *tag.Tag) error {
	body := struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    TagResponse `json:"data"`
	}{
		Code:    statusCode,
		Message: msg,
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

func writeTags(w http.ResponseWriter, statusCode int, msg string, tags []*tag.Tag, option *tag.FindOption) error {
	responses := make([]TagResponse, 0, len(tags))
	var lastID tag.ID
	for _, tag := range tags {
		responses = append(
			responses,
			TagResponse{
				ID:          tag.ID,
				Name:        tag.Name,
				Description: tag.Description,
				CreatedAt:   tag.CreatedAt,
				UpdatedAt:   tag.UpdatedAt,
			},
		)
		if lastID < tag.ID {
			lastID = tag.ID
		}
	}

	switch option.SearchMode {
	case tag.SearchMode_Seek:
		type Next struct {
			LastID tag.LastID `json:"last_id"`
			Count  tag.Count  `json:"count"`
		}
		var next *Next
		if len(responses) == int(option.Seek.Count) {
			next = &Next{
				LastID: lastID,
				Count:  option.Seek.Count,
			}
		}

		body := struct {
			Code    int           `json:"code"`
			Message string        `json:"message"`
			Data    []TagResponse `json:"data"`
			Next    *Next         `json:"next"`
		}{
			Code:    statusCode,
			Message: msg,
			Data:    responses,
			Next:    next,
		}
		w.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(w).Encode(&body)
	case tag.SearchMode_Pagination:
		type Next struct {
			Limit  tag.Limit  `json:"limit"`
			Offset tag.Offset `json:"offset"`
		}
		var next *Next
		if len(responses) == int(option.Pagination.Limit) {
			next = &Next{
				Limit:  option.Pagination.Limit,
				Offset: option.Pagination.Offset + len(responses),
			}
		}

		body := struct {
			Code    int           `json:"code"`
			Message string        `json:"message"`
			Data    []TagResponse `json:"data"`
			Next    *Next         `json:"next"`
		}{
			Code:    statusCode,
			Message: msg,
			Data:    responses,
			Next:    next,
		}
		w.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(w).Encode(&body)
	default:
		body := struct {
			Code    int           `json:"code"`
			Message string        `json:"message"`
			Data    []TagResponse `json:"data"`
		}{
			Code:    statusCode,
			Message: msg,
			Data:    responses,
		}
		w.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(w).Encode(&body)
	}
}
