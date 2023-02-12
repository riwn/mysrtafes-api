package game

import (
	"encoding/json"
	"mysrtafes-backend/pkg/game"
	"mysrtafes-backend/pkg/game/platform"
	"mysrtafes-backend/pkg/game/tag"
	"net/http"
	"time"
)

type GameResponse struct {
	ID          game.ID            `json:"id"`
	Name        game.Name          `json:"name"`
	Description game.Description   `json:"description"`
	Publisher   game.Publisher     `json:"publisher"`
	Developer   game.Developer     `json:"developer"`
	Links       []LinkResponse     `json:"links"`
	Tags        []TagResponse      `json:"tags"`
	Platforms   []PlatformResponse `json:"platforms"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	// TODO: LaravelでなぜかReleaseDateを追加するの忘れてた。
	// ReleaseDate string           `json:"release_date"`
}

type PlatformResponse struct {
	ID          platform.ID          `json:"id"`
	Name        platform.Name        `json:"name"`
	Description platform.Description `json:"description"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

type TagResponse struct {
	ID          tag.ID          `json:"id"`
	Name        tag.Name        `json:"name"`
	Description tag.Description `json:"description"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type LinkResponse struct {
	URL         string               `json:"url"`
	Description game.LinkDescription `json:"description"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

func WriteCreateGame(w http.ResponseWriter, game *game.Game) error {
	return writeGame(w, http.StatusCreated, "success create game", game)
}

// write read response for game
func WriteReadGame(w http.ResponseWriter, game *game.Game) error {
	return writeGame(w, http.StatusOK, "success read game", game)
}

// write update response for game
func WriteUpdateGame(w http.ResponseWriter, game *game.Game) error {
	return writeGame(w, http.StatusOK, "success update game", game)
}

// write delete response for game
func WriteDeleteGame(w http.ResponseWriter, gameID game.ID) error {
	body := struct {
		Code    int     `json:"code"`
		Message string  `json:"message"`
		Data    game.ID `json:"deleteID"`
	}{
		Code:    http.StatusOK,
		Message: "success delete game",
		Data:    gameID,
	}
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(&body)
}

// write find response for game
func WriteFindGame(w http.ResponseWriter, games []*game.Game, option *game.FindOption) error {
	return writeGames(w, http.StatusOK, "success find game", games, option)
}

func writeGame(w http.ResponseWriter, statusCode int, msg string, game *game.Game) error {
	links := make([]LinkResponse, 0, len(game.Links))
	for _, link := range game.Links {
		links = append(links, LinkResponse{
			URL:         link.URL.URL().String(),
			Description: link.LinkDescription,
			CreatedAt:   link.CreatedAt,
			UpdatedAt:   link.UpdatedAt,
		})
	}

	tags := make([]TagResponse, 0, len(game.Tags))
	for _, tag := range game.Tags {
		tags = append(tags, TagResponse{
			ID:          tag.ID,
			Name:        tag.Name,
			Description: tag.Description,
			CreatedAt:   tag.CreatedAt,
			UpdatedAt:   tag.UpdatedAt,
		})
	}

	platforms := make([]PlatformResponse, 0, len(game.Platforms))
	for _, platform := range game.Platforms {
		platforms = append(platforms, PlatformResponse{
			ID:          platform.ID,
			Name:        platform.Name,
			Description: platform.Description,
			CreatedAt:   platform.CreatedAt,
			UpdatedAt:   platform.UpdatedAt,
		})
	}

	body := struct {
		Code    int          `json:"code"`
		Message string       `json:"message"`
		Data    GameResponse `json:"data"`
	}{
		Code:    statusCode,
		Message: msg,
		Data: GameResponse{
			ID:          game.ID,
			Name:        game.Name,
			Description: game.Description,
			Publisher:   game.Publisher,
			Developer:   game.Developer,
			Links:       links,
			Tags:        tags,
			Platforms:   platforms,
			CreatedAt:   game.CreatedAt,
			UpdatedAt:   game.UpdatedAt,
			// ReleaseDate: game.ReleaseDate.String(),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(&body)
}

func writeGames(w http.ResponseWriter, statusCode int, msg string, games []*game.Game, option *game.FindOption) error {
	responses := make([]GameResponse, 0, len(games))
	var lastID game.ID
	for _, game := range games {
		links := make([]LinkResponse, 0, len(game.Links))
		for _, link := range game.Links {
			links = append(links, LinkResponse{
				URL:         link.URL.URL().String(),
				Description: link.LinkDescription,
				CreatedAt:   link.CreatedAt,
				UpdatedAt:   link.UpdatedAt,
			})
		}

		tags := make([]TagResponse, 0, len(game.Tags))
		for _, tag := range game.Tags {
			tags = append(tags, TagResponse{
				ID:          tag.ID,
				Name:        tag.Name,
				Description: tag.Description,
				CreatedAt:   tag.CreatedAt,
				UpdatedAt:   tag.UpdatedAt,
			})
		}

		platforms := make([]PlatformResponse, 0, len(game.Platforms))
		for _, platform := range game.Platforms {
			platforms = append(platforms, PlatformResponse{
				ID:          platform.ID,
				Name:        platform.Name,
				Description: platform.Description,
				CreatedAt:   platform.CreatedAt,
				UpdatedAt:   platform.UpdatedAt,
			})
		}
		responses = append(
			responses,
			GameResponse{
				ID:          game.ID,
				Name:        game.Name,
				Description: game.Description,
				Publisher:   game.Publisher,
				Developer:   game.Developer,
				Links:       links,
				Tags:        tags,
				Platforms:   platforms,
				CreatedAt:   game.CreatedAt,
				UpdatedAt:   game.UpdatedAt,
				// ReleaseDate: game.ReleaseDate.String(),
			},
		)
		if lastID < game.ID {
			lastID = game.ID
		}
	}

	switch option.SearchMode {
	case game.SearchMode_Seek:
		type Next struct {
			LastID game.LastID `json:"last_id"`
			Count  game.Count  `json:"count"`
		}
		var next *Next
		if len(responses) == int(option.Seek.Count) {
			next = &Next{
				LastID: lastID,
				Count:  option.Seek.Count,
			}
		}

		body := struct {
			Code    int            `json:"code"`
			Message string         `json:"message"`
			Data    []GameResponse `json:"data"`
			Next    *Next          `json:"next"`
		}{
			Code:    statusCode,
			Message: msg,
			Data:    responses,
			Next:    next,
		}
		w.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(w).Encode(&body)
	case game.SearchMode_Pagination:
		type Next struct {
			Limit  game.Limit  `json:"limit"`
			Offset game.Offset `json:"offset"`
		}
		var next *Next
		if len(responses) == int(option.Pagination.Limit) {
			next = &Next{
				Limit:  option.Pagination.Limit,
				Offset: option.Pagination.Offset + len(responses),
			}
		}

		body := struct {
			Code    int            `json:"code"`
			Message string         `json:"message"`
			Data    []GameResponse `json:"data"`
			Next    *Next          `json:"next"`
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
			Code    int            `json:"code"`
			Message string         `json:"message"`
			Data    []GameResponse `json:"data"`
		}{
			Code:    statusCode,
			Message: msg,
			Data:    responses,
		}
		w.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(w).Encode(&body)
	}
}
