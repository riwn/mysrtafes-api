package game

import (
	"encoding/json"
	"mysrtafes-backend/pkg/errors"
	"mysrtafes-backend/pkg/game"
	"mysrtafes-backend/pkg/game/platform"
	"mysrtafes-backend/pkg/game/tag"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func NewGameCreate(r *http.Request) (*game.Game, []platform.ID, []tag.ID, error) {
	defer r.Body.Close()

	type Link struct {
		Title           game.Title           `json:"title"`
		URL             string               `json:"url"`
		LinkDescription game.LinkDescription `json:"description"`
	}

	body := struct {
		Name        game.Name        `json:"name"`
		Description game.Description `json:"description"`
		Publisher   game.Publisher   `json:"publisher"`
		Developer   game.Developer   `json:"developer"`
		ReleaseDate string           `json:"release_date"`
		Links       []Link           `json:"links"`
		PlatformIDs []platform.ID    `json:"platform_ids"`
		TagIDs      []tag.ID         `json:"tag_ids"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, nil, nil, errors.NewInvalidRequest(
			errors.Layer_Request,
			errors.NewInformation(
				errors.ID_JsonDecodeError,
				err.Error(),
				nil,
			),
			"json decode error. bad format request.",
		)
	}

	releaseDate, err := game.NewReleaseDate(body.ReleaseDate)
	if err != nil {
		return nil, nil, nil, errors.NewInvalidRequest(
			errors.Layer_Request,
			errors.NewInformation(
				errors.ID_InvalidParams,
				err.Error(),
				[]errors.InvalidParams{
					errors.NewInvalidParams("release_date", body.ReleaseDate),
				},
			),
			"release_date create error",
		)
	}

	Links := make([]*game.Link, 0, len(body.Links))
	for _, link := range body.Links {
		url, err := game.NewURL(link.URL)
		if err != nil {
			return nil, nil, nil, errors.NewInvalidRequest(
				errors.Layer_Request,
				errors.NewInformation(
					errors.ID_InvalidParams,
					err.Error(),
					[]errors.InvalidParams{
						errors.NewInvalidParams("links.url", link.URL),
					},
				),
				"links.url create error",
			)
		}
		Links = append(Links, game.NewLink(link.Title, url, link.LinkDescription))
	}

	return game.New(
		body.Name,
		body.Description,
		body.Publisher,
		body.Developer,
		releaseDate,
		Links,
	), body.PlatformIDs, body.TagIDs, nil
}

func NewGameID(r *http.Request) (game.ID, error) {
	gameIDStr := chi.URLParam(r, "gameID")

	gameID, err := strconv.Atoi(gameIDStr)
	if err != nil {
		return 0, errors.NewInvalidRequest(
			errors.Layer_Request,
			errors.NewInformation(
				errors.ID_InvalidParams,
				err.Error(),
				[]errors.InvalidParams{
					errors.NewInvalidParams("gameID", gameIDStr),
				},
			),
			"gameID convert error",
		)
	}
	return game.ID(gameID), nil
}

func NewGameFindOption(r *http.Request) (*game.FindOption, error) {
	// デフォルト値生成
	findOption := game.NewFindOption()

	q := r.URL.Query()
	// 検索設定
	if err := setSearchMode(findOption, q); err != nil {
		return nil, err
	}

	// 並び替え設定
	if err := setOrder(findOption, q); err != nil {
		return nil, err
	}

	return findOption, nil
}

func NewGameUpdate(r *http.Request) (*game.Game, []platform.ID, []tag.ID, error) {
	defer r.Body.Close()

	gameID, err := NewGameID(r)
	if err != nil {
		return nil, nil, nil, err
	}

	type Link struct {
		Title           game.Title           `json:"title"`
		URL             string               `json:"url"`
		LinkDescription game.LinkDescription `json:"description"`
	}

	body := struct {
		Name        game.Name        `json:"name"`
		Description game.Description `json:"description"`
		Publisher   game.Publisher   `json:"publisher"`
		Developer   game.Developer   `json:"developer"`
		ReleaseDate string           `json:"release_date"`
		Links       []Link           `json:"links"`
		PlatformIDs []platform.ID    `json:"platform_ids"`
		TagIDs      []tag.ID         `json:"tag_ids"`
	}{}
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, nil, nil, errors.NewInvalidRequest(
			errors.Layer_Request,
			errors.NewInformation(
				errors.ID_JsonDecodeError,
				err.Error(),
				nil,
			),
			"json decode error. bad format request.",
		)
	}

	releaseDate, err := game.NewReleaseDate(body.ReleaseDate)
	if err != nil {
		return nil, nil, nil, errors.NewInvalidRequest(
			errors.Layer_Request,
			errors.NewInformation(
				errors.ID_InvalidParams,
				err.Error(),
				[]errors.InvalidParams{
					errors.NewInvalidParams("release_date", body.ReleaseDate),
				},
			),
			"release_date create error",
		)
	}

	Links := make([]*game.Link, 0, len(body.Links))
	for _, link := range body.Links {
		url, err := game.NewURL(link.URL)
		if err != nil {
			return nil, nil, nil, errors.NewInvalidRequest(
				errors.Layer_Request,
				errors.NewInformation(
					errors.ID_InvalidParams,
					err.Error(),
					[]errors.InvalidParams{
						errors.NewInvalidParams("links.url", link.URL),
					},
				),
				"links.url create error",
			)
		}
		Links = append(Links, game.NewLink(link.Title, url, link.LinkDescription))
	}

	return game.NewWithID(
		gameID,
		body.Name,
		body.Description,
		body.Publisher,
		body.Developer,
		releaseDate,
		Links,
	), body.PlatformIDs, body.TagIDs, nil
}

// Find: set order param
func setOrder(findOption *game.FindOption, q url.Values) error {
	// シーク法のときは並び替えするとおかしくなるので禁止
	if findOption.SearchMode == game.SearchMode_Seek {
		return nil
	}
	// Descのチェック
	desc := false
	var err error
	if q.Has("desc") {
		desc, err = strconv.ParseBool(q.Get("desc"))
		if err != nil {
			return errors.NewInvalidRequest(
				errors.Layer_Request,
				errors.NewInformation(
					errors.ID_InvalidParams,
					err.Error(),
					[]errors.InvalidParams{
						errors.NewInvalidParams("desc", q.Get("desc")),
					},
				),
				"desc convert error",
			)
		}
	}

	// 並び順がないときはIDのOrderで返却
	if !q.Has("order") {
		findOption.SetOrder(game.Order_ID, desc)
		return nil
	}

	// 並び順のチェック
	switch q.Get("order") {
	case "name":
		findOption.SetOrder(game.Order_Name, desc)
	case "id":
		findOption.SetOrder(game.Order_ID, desc)
	default:
		return errors.NewInvalidRequest(
			errors.Layer_Request,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"order convert error",
				[]errors.InvalidParams{
					errors.NewInvalidParams("order", q.Get("order")),
				},
			),
			"order convert error",
		)
	}
	return nil
}

// Find: set search mode param
func setSearchMode(findOption *game.FindOption, q url.Values) error {
	// モードがないときは何もせず終了
	if !q.Has("mode") {
		return nil
	}
	// 検索モードのチェック
	var err error
	switch q.Get("mode") {
	case "seek":
		// シーク法
		var lastID, count int = 0, 30
		if q.Has("last_id") {
			lastIDStr := q.Get("last_id")
			lastID, err = strconv.Atoi(lastIDStr)
			if err != nil {
				return errors.NewInvalidRequest(
					errors.Layer_Request,
					errors.NewInformation(
						errors.ID_InvalidParams,
						err.Error(),
						[]errors.InvalidParams{
							errors.NewInvalidParams("last_id", lastIDStr),
						},
					),
					"last_id convert error",
				)
			}
		}
		if q.Has("count") {
			CountStr := q.Get("count")
			count, err = strconv.Atoi(CountStr)
			if err != nil {
				return errors.NewInvalidRequest(
					errors.Layer_Request,
					errors.NewInformation(
						errors.ID_InvalidParams,
						err.Error(),
						[]errors.InvalidParams{
							errors.NewInvalidParams("count", CountStr),
						},
					),
					"count convert error",
				)
			}
		}
		findOption.SetSeek(game.LastID(lastID), game.Count(count))
	case "page":
		// ページネーション法
		var limit, offset int = 30, 0
		var err error
		if q.Has("limit") {
			limitStr := q.Get("limit")
			limit, err = strconv.Atoi(limitStr)
			if err != nil {
				return errors.NewInvalidRequest(
					errors.Layer_Request,
					errors.NewInformation(
						errors.ID_InvalidParams,
						err.Error(),
						[]errors.InvalidParams{
							errors.NewInvalidParams("limit", limitStr),
						},
					),
					"limit convert error",
				)
			}
		}
		if q.Has("offset") {
			OffsetStr := q.Get("offset")
			offset, err = strconv.Atoi(OffsetStr)
			if err != nil {
				return errors.NewInvalidRequest(
					errors.Layer_Request,
					errors.NewInformation(
						errors.ID_InvalidParams,
						err.Error(),
						[]errors.InvalidParams{
							errors.NewInvalidParams("offset", OffsetStr),
						},
					),
					"offset convert error",
				)
			}
		}
		findOption.SetPagination(game.Limit(limit), game.Offset(offset))
	default:
		return errors.NewInvalidRequest(
			errors.Layer_Request,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"nothing mode error",
				[]errors.InvalidParams{
					errors.NewInvalidParams("mode", q.Get("mode")),
				},
			),
			"nothing mode error",
		)
	}
	return nil
}
