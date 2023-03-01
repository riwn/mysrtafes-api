package tag

import (
	"encoding/json"
	"mysrtafes-backend/pkg/game/tag"
	"net/http"
	"time"
)

type Tag struct {
	ID          tag.ID          `json:"id"`
	Name        tag.Name        `json:"name"`
	Description tag.Description `json:"description"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type TagResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Tag    `json:"data"`
}

type TagsResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []Tag  `json:"data"`
}

type Next struct {
	LastID tag.LastID `json:"last_id"`
	Count  tag.Count  `json:"count"`
}

type TagsNextResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []Tag  `json:"data"`
	Next    *Next  `json:"next"`
}

type Page struct {
	Limit  tag.Limit  `json:"limit"`
	Offset tag.Offset `json:"offset"`
}

type TagsPageResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []Tag  `json:"data"`
	Page    *Page  `json:"page"`
}

// write create response for tag
func WriteCreateTag(w http.ResponseWriter, tag *tag.Tag) error {
	body := tagResponse(http.StatusCreated, "success create tag", tag)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(&body)
}

// write read response for tag
func WriteReadTag(w http.ResponseWriter, tag *tag.Tag) error {
	body := tagResponse(http.StatusOK, "success read tag", tag)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(&body)
}

// write update response for tag
func WriteUpdateTag(w http.ResponseWriter, tag *tag.Tag) error {
	body := tagResponse(http.StatusOK, "success update tag", tag)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(&body)
}

// write delete response for tag
func WriteDeleteTag(w http.ResponseWriter, tagID tag.ID) error {
	body := deleteTagResponse(tagID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(&body)
}

// write find response for tag
func WriteFindTag(w http.ResponseWriter, tags []*tag.Tag, option *tag.FindOption) error {
	body := tagsResponse(http.StatusOK, "success find tag", tags, option)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(&body)
}

func tagResponse(statusCode int, msg string, tag *tag.Tag) interface{} {
	return TagResponse{
		Code:    statusCode,
		Message: msg,
		Data: Tag{
			ID:          tag.ID,
			Name:        tag.Name,
			Description: tag.Description,
			CreatedAt:   tag.CreatedAt,
			UpdatedAt:   tag.UpdatedAt,
		},
	}
}

func deleteTagResponse(tagID tag.ID) interface{} {
	return struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    tag.ID `json:"deleteID"`
	}{
		Code:    http.StatusOK,
		Message: "success delete tag",
		Data:    tagID,
	}
}

func tagsResponse(statusCode int, msg string, tags []*tag.Tag, option *tag.FindOption) interface{} {
	responses := make([]Tag, 0, len(tags))
	var lastID tag.ID
	for _, tag := range tags {
		responses = append(
			responses,
			Tag{
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
		var next *Next
		if len(responses) == int(option.Seek.Count) {
			next = &Next{
				LastID: lastID,
				Count:  option.Seek.Count,
			}
		}

		return TagsNextResponse{
			Code:    statusCode,
			Message: msg,
			Data:    responses,
			Next:    next,
		}
	case tag.SearchMode_Pagination:
		var page *Page
		if len(responses) == int(option.Pagination.Limit) {
			page = &Page{
				Limit:  option.Pagination.Limit,
				Offset: option.Pagination.Offset + len(responses),
			}
		}

		return TagsPageResponse{
			Code:    statusCode,
			Message: msg,
			Data:    responses,
			Page:    page,
		}
	default:
		return TagsResponse{
			Code:    statusCode,
			Message: msg,
			Data:    responses,
		}
	}
}
