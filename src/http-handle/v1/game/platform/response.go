package platform

import (
	"encoding/json"
	"mysrtafes-backend/pkg/game/platform"
	"net/http"
	"time"
)

type PlatformResponse struct {
	ID          platform.ID          `json:"id"`
	Name        platform.Name        `json:"name"`
	Description platform.Description `json:"description"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

// write create response for platform
func WriteCreatePlatform(w http.ResponseWriter, platform *platform.Platform) error {
	return writePlatform(w, http.StatusCreated, "success create platform", platform)
}

// write read response for platform
func WriteReadPlatform(w http.ResponseWriter, platform *platform.Platform) error {
	return writePlatform(w, http.StatusOK, "success read platform", platform)
}

// write update response for platform
func WriteUpdatePlatform(w http.ResponseWriter, platform *platform.Platform) error {
	return writePlatform(w, http.StatusOK, "success update platform", platform)
}

// write delete response for platform
func WriteDeletePlatform(w http.ResponseWriter, platformID platform.ID) error {
	body := struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    platform.ID `json:"deleteID"`
	}{
		Code:    http.StatusOK,
		Message: "success delete platform",
		Data:    platformID,
	}
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(&body)
}

// write find response for platform
func WriteFindPlatform(w http.ResponseWriter, platforms []*platform.Platform, option *platform.FindOption) error {
	return writePlatforms(w, http.StatusOK, "success find platform", platforms, option)
}

func writePlatform(w http.ResponseWriter, statusCode int, msg string, platform *platform.Platform) error {
	body := struct {
		Code    int              `json:"code"`
		Message string           `json:"message"`
		Data    PlatformResponse `json:"data"`
	}{
		Code:    statusCode,
		Message: msg,
		Data: PlatformResponse{
			ID:          platform.ID,
			Name:        platform.Name,
			Description: platform.Description,
			CreatedAt:   platform.CreatedAt,
			UpdatedAt:   platform.UpdatedAt,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(&body)
}

func writePlatforms(w http.ResponseWriter, statusCode int, msg string, platforms []*platform.Platform, option *platform.FindOption) error {
	responses := make([]PlatformResponse, 0, len(platforms))
	var lastID platform.ID
	for _, platform := range platforms {
		responses = append(
			responses,
			PlatformResponse{
				ID:          platform.ID,
				Name:        platform.Name,
				Description: platform.Description,
				CreatedAt:   platform.CreatedAt,
				UpdatedAt:   platform.UpdatedAt,
			},
		)
		if lastID < platform.ID {
			lastID = platform.ID
		}
	}

	switch option.SearchMode {
	case platform.SearchMode_Seek:
		type Next struct {
			LastID platform.LastID `json:"last_id"`
			Count  platform.Count  `json:"count"`
		}
		var next *Next
		if len(responses) == int(option.Seek.Count) {
			next = &Next{
				LastID: lastID,
				Count:  option.Seek.Count,
			}
		}

		body := struct {
			Code    int                `json:"code"`
			Message string             `json:"message"`
			Data    []PlatformResponse `json:"data"`
			Next    *Next              `json:"next"`
		}{
			Code:    statusCode,
			Message: msg,
			Data:    responses,
			Next:    next,
		}
		w.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(w).Encode(&body)
	case platform.SearchMode_Pagination:
		type Next struct {
			Limit  platform.Limit  `json:"limit"`
			Offset platform.Offset `json:"offset"`
		}
		var next *Next
		if len(responses) == int(option.Pagination.Limit) {
			next = &Next{
				Limit:  option.Pagination.Limit,
				Offset: option.Pagination.Offset + len(responses),
			}
		}

		body := struct {
			Code    int                `json:"code"`
			Message string             `json:"message"`
			Data    []PlatformResponse `json:"data"`
			Next    *Next              `json:"next"`
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
			Code    int                `json:"code"`
			Message string             `json:"message"`
			Data    []PlatformResponse `json:"data"`
		}{
			Code:    statusCode,
			Message: msg,
			Data:    responses,
		}
		w.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(w).Encode(&body)
	}
}
