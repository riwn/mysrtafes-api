package game

import (
	"encoding/json"
	"mysrtafes-backend/pkg/game"
	"net/http"
	"time"
)

type GameResponse struct {
	ID          game.ID          `json:"id"`
	Name        game.Name        `json:"name"`
	Description game.Description `json:"description"`
	Publisher   game.Publisher   `json:"publisher"`
	Developer   game.Developer   `json:"developer"`
	ReleaseDate game.ReleaseDate `json:"release_date"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
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
			ReleaseDate: game.ReleaseDate,
			CreatedAt:   game.CreatedAt,
			UpdatedAt:   game.UpdatedAt,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(&body)
}

func writeGames(w http.ResponseWriter, statusCode int, msg string, games []*game.Game, option *game.FindOption) error {
	responses := make([]GameResponse, 0, len(games))
	var lastID game.ID
	for _, game := range games {
		responses = append(
			responses,
			GameResponse{
				ID:          game.ID,
				Name:        game.Name,
				Description: game.Description,
				Publisher:   game.Publisher,
				Developer:   game.Developer,
				ReleaseDate: game.ReleaseDate,
				CreatedAt:   game.CreatedAt,
				UpdatedAt:   game.UpdatedAt,
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
