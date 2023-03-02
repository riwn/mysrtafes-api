package challenge

import (
	"encoding/json"
	"mysrtafes-backend/pkg/challenge"
	"mysrtafes-backend/pkg/challenge/detail"
	"mysrtafes-backend/pkg/challenge/detail/goal"
	"mysrtafes-backend/pkg/errors"
	"mysrtafes-backend/pkg/game"
	"net/http"
)

type Challenge struct {
	Name        challenge.Name        `json:"name"`
	ReadingName challenge.ReadingName `json:"name_read"`
	Password    challenge.Password    `json:"password"`
	Twitter     challenge.Twitter     `json:"twitter"`
	Discord     challenge.Discord     `json:"discord"`
	IsStream    challenge.IsStream    `json:"is_stream"`
	URL         string                `json:"stream_url"`
	Comment     challenge.Comment     `json:"comment"`
	Details     []*Detail             `json:"challenge_details"`
}

type Detail struct {
	GameID     game.ID           `json:"game_master_id"`
	GameName   game.Name         `json:"game_name"`
	GoalIDs    []goal.ID         `json:"goal_genre_master_ids"`
	GoalDetail detail.GoalDetail `json:"goal_detail"`
	Department detail.Department `json:"department"`
}

func NewChallengeCreate(r *http.Request) (*challenge.Challenge, error) {
	defer r.Body.Close()

	body := Challenge{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Request,
			errors.NewInformation(
				errors.ID_JsonDecodeError,
				err.Error(),
				nil,
			),
			"json decode error. bad format request",
		)
	}

	if len(body.Details) == 0 {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Request,
			errors.NewInformation(
				errors.ID_InvalidParams,
				err.Error(),
				[]errors.InvalidParams{
					errors.NewInvalidParams("detail", body.Details),
				},
			),
			"detail nothing error",
		)
	}

	// Detailの生成
	details := make([]*detail.Detail, 0, len(body.Details))
	for _, bodyDetail := range body.Details {
		details = append(
			details,
			detail.New(
				bodyDetail.GameID,
				bodyDetail.GameName,
				bodyDetail.GoalIDs,
				bodyDetail.GoalDetail,
				bodyDetail.Department,
			),
		)
	}

	url, err := challenge.NewURL(body.URL)
	if err != nil {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Request,
			errors.NewInformation(
				errors.ID_InvalidParams,
				err.Error(),
				[]errors.InvalidParams{
					errors.NewInvalidParams("url", body.URL),
				},
			),
			"url create error",
		)
	}

	return challenge.New(
		body.Name,
		body.ReadingName,
		body.Password,
		body.Twitter,
		body.Discord,
		body.IsStream,
		url,
		body.Comment,
		details,
	), nil
}
