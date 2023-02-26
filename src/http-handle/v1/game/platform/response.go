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
	body := platFormResponse(http.StatusOK, "success create platform", platform)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(&body)
}

// write read response for platform
func WriteReadPlatform(w http.ResponseWriter, platform *platform.Platform) error {
	body := platFormResponse(http.StatusOK, "success read platform", platform)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(&body)
}

// write update response for platform
func WriteUpdatePlatform(w http.ResponseWriter, platform *platform.Platform) error {
	body := platFormResponse(http.StatusOK, "success update platform", platform)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(&body)
}

// write delete response for platform
func WriteDeletePlatform(w http.ResponseWriter, platformID platform.ID) error {
	body := deletePlatformResponse(platformID)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(&body)
}

// write find response for platform
func WriteFindPlatform(w http.ResponseWriter, platforms []*platform.Platform, option *platform.FindOption) error {
	body := platFormsResponse(http.StatusOK, "success find platform", platforms, option)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(&body)
}

func platFormResponse(statusCode int, msg string, platform *platform.Platform) interface{} {
	return struct {
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
}

func deletePlatformResponse(platformID platform.ID) interface{} {
	return struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    platform.ID `json:"deleteID"`
	}{
		Code:    http.StatusOK,
		Message: "success delete platform",
		Data:    platformID,
	}
}

func platFormsResponse(statusCode int, msg string, platforms []*platform.Platform, option *platform.FindOption) interface{} {
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

		return struct {
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
	case platform.SearchMode_Pagination:
		type Page struct {
			Limit  platform.Limit  `json:"limit"`
			Offset platform.Offset `json:"offset"`
		}
		var page *Page
		if len(responses) == int(option.Pagination.Limit) {
			page = &Page{
				Limit:  option.Pagination.Limit,
				Offset: option.Pagination.Offset + len(responses),
			}
		}

		return struct {
			Code    int                `json:"code"`
			Message string             `json:"message"`
			Data    []PlatformResponse `json:"data"`
			Page    *Page              `json:"page"`
		}{
			Code:    statusCode,
			Message: msg,
			Data:    responses,
			Page:    page,
		}
	default:
		return struct {
			Code    int                `json:"code"`
			Message string             `json:"message"`
			Data    []PlatformResponse `json:"data"`
		}{
			Code:    statusCode,
			Message: msg,
			Data:    responses,
		}
	}
}
