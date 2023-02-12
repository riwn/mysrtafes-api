package platform

import (
	"encoding/json"
	"mysrtafes-backend/pkg/errors"
	"mysrtafes-backend/pkg/game/platform"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Post: NewPlatformEntity for request
func NewPlatformCreate(r *http.Request) (*platform.Platform, error) {
	defer r.Body.Close()

	body := struct {
		Name        platform.Name        `json:"name"`
		Description platform.Description `json:"description"`
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

	return platform.New(
		body.Name,
		body.Description,
	), nil
}

// Delete: NewPlatformID for request
func NewPlatformID(r *http.Request) (platform.ID, error) {
	platformIDStr := chi.URLParam(r, "platformID")

	platformID, err := strconv.Atoi(platformIDStr)
	if err != nil {
		return 0, errors.NewInvalidRequest(
			errors.Layer_Request,
			errors.NewInformation(
				errors.ID_InvalidParams,
				err.Error(),
				[]errors.InvalidParams{
					errors.NewInvalidParams("platformID", platformIDStr),
				},
			),
			"platformID convert error",
		)
	}
	return platform.ID(platformID), nil
}

// Find: NewFindOptionEntity for request
func NewPlatformFindOption(r *http.Request) (*platform.FindOption, error) {
	// デフォルト値生成
	findOption := platform.NewFindOption()

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

// Update: NewPlatformEntity for request
func NewPlatformUpdate(r *http.Request) (*platform.Platform, error) {
	defer r.Body.Close()

	platformID, err := NewPlatformID(r)
	if err != nil {
		return nil, err
	}

	body := struct {
		Name        platform.Name        `json:"name"`
		Description platform.Description `json:"description"`
	}{}
	err = json.NewDecoder(r.Body).Decode(&body)
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

	return platform.NewWithID(
		platformID,
		body.Name,
		body.Description,
	), nil
}

// Find: set order param
func setOrder(findOption *platform.FindOption, q url.Values) error {
	// シーク法のときは並び替えするとおかしくなるので禁止
	if findOption.SearchMode == platform.SearchMode_Seek {
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
		findOption.SetOrder(platform.Order_ID, desc)
		return nil
	}

	// 並び順のチェック
	switch q.Get("order") {
	case "name":
		findOption.SetOrder(platform.Order_Name, desc)
	case "id":
		findOption.SetOrder(platform.Order_ID, desc)
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
func setSearchMode(findOption *platform.FindOption, q url.Values) error {
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
		findOption.SetSeek(platform.LastID(lastID), platform.Count(count))
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
		findOption.SetPagination(platform.Limit(limit), platform.Offset(offset))
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
