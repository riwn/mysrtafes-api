package challenge

import (
	"encoding/json"
	"mysrtafes-backend/pkg/challenge"
	"mysrtafes-backend/pkg/challenge/detail"
	"mysrtafes-backend/pkg/challenge/detail/goal"
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

	body := struct {
		Challenge
	}{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	// Detailの生成
	var details []*detail.Detail
	for _, bodyDetail := range body.Challenge.Details {
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

	url, err := challenge.NewURL(body.Challenge.URL)
	if err != nil {
		return nil, err
	}

	return challenge.New(
		body.Challenge.Name,
		body.Challenge.ReadingName,
		body.Challenge.Password,
		body.Challenge.Twitter,
		body.Challenge.Discord,
		body.Challenge.IsStream,
		url,
		body.Challenge.Comment,
		details,
	), nil
}
