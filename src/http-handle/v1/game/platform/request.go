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

func NewPlatformID(r *http.Request) (platform.ID, error) {
	platformIDStr := chi.URLParam(r, "platformID")
	// 空文字の時、0にして複数検索と判断
	if platformIDStr == "" {
		return 0, nil
	}

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

func NewPlatformFindOption(r *http.Request) (*platform.FindOption, error) {
	// デフォルト値生成
	findOption := platform.NewFindOption()
	// NOTE: QueryがNilのときはデフォルトで返却
	q := r.URL.Query()
	if q == nil {
		return findOption, nil
	}

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

func NewPlatformUpdate(r *http.Request) (*platform.Platform, error) {
	defer r.Body.Close()

	body := struct {
		ID          platform.ID          `json:"id"`
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

	return platform.NewWithID(
		body.ID,
		body.Name,
		body.Description,
	), nil
}

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

	// 並び順のチェック
	switch q.Get("order") {
	case "name":
		findOption.SetOrder(platform.Order_Name, desc)
	default:
		findOption.SetOrder(platform.Order_ID, desc)
	}
	return nil
}

func setSearchMode(findOption *platform.FindOption, q url.Values) error {
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
	}
	return nil
}
