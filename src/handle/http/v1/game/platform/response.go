package platform

import (
	"encoding/json"
	"mysrtafes-backend/pkg/game/platform"
	"net/http"
	"time"
)

type Platform struct {
	ID          platform.ID          `json:"id"`
	Name        platform.Name        `json:"name"`
	Description platform.Description `json:"description"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

type PlatformResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    Platform `json:"data"`
}

type PlatformsResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []Platform `json:"data"`
}

type Page struct {
	Limit  platform.Limit  `json:"limit"`
	Offset platform.Offset `json:"offset"`
}

type PlatformsPageResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []Platform `json:"data"`
	Page    *Page      `json:"page"`
}

type Next struct {
	LastID platform.LastID `json:"last_id"`
	Count  platform.Count  `json:"count"`
}

type PlatformsNextResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []Platform `json:"data"`
	Next    *Next      `json:"next"`
}

// write create response for platform
func WriteCreatePlatform(w http.ResponseWriter, platform *platform.Platform) error {
	body := platformResponse(http.StatusCreated, "success create platform", platform)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(&body)
}

// write read response for platform
func WriteReadPlatform(w http.ResponseWriter, platform *platform.Platform) error {
	body := platformResponse(http.StatusOK, "success read platform", platform)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(&body)
}

// write update response for platform
func WriteUpdatePlatform(w http.ResponseWriter, platform *platform.Platform) error {
	body := platformResponse(http.StatusOK, "success update platform", platform)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(&body)
}

// write delete response for platform
func WriteDeletePlatform(w http.ResponseWriter, platformID platform.ID) error {
	body := deletePlatformResponse(platformID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(&body)
}

// write find response for platform
func WriteFindPlatform(w http.ResponseWriter, platforms []*platform.Platform, option *platform.FindOption) error {
	body := platformsResponse(http.StatusOK, "success find platform", platforms, option)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(&body)
}

func platformResponse(statusCode int, msg string, platform *platform.Platform) interface{} {
	return PlatformResponse{
		Code:    statusCode,
		Message: msg,
		Data: Platform{
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

func platformsResponse(statusCode int, msg string, platforms []*platform.Platform, option *platform.FindOption) interface{} {
	responses := make([]Platform, 0, len(platforms))
	var lastID platform.ID
	for _, platform := range platforms {
		responses = append(
			responses,
			Platform{
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
		var next *Next
		if len(responses) == int(option.Seek.Count) {
			next = &Next{
				LastID: lastID,
				Count:  option.Seek.Count,
			}
		}

		return PlatformsNextResponse{
			Code:    statusCode,
			Message: msg,
			Data:    responses,
			Next:    next,
		}
	case platform.SearchMode_Pagination:
		var page *Page
		if len(responses) == int(option.Pagination.Limit) {
			page = &Page{
				Limit:  option.Pagination.Limit,
				Offset: option.Pagination.Offset + len(responses),
			}
		}

		return PlatformsPageResponse{
			Code:    statusCode,
			Message: msg,
			Data:    responses,
			Page:    page,
		}
	default:
		return PlatformsResponse{
			Code:    statusCode,
			Message: msg,
			Data:    responses,
		}
	}
}
