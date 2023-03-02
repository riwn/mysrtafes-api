package challenge

import (
	"encoding/json"
	challenges "mysrtafes-backend/pkg/challenge"
	"mysrtafes-backend/pkg/challenge/detail"
	"mysrtafes-backend/pkg/challenge/detail/goal"
	"mysrtafes-backend/pkg/challenge/stream"
	"mysrtafes-backend/pkg/game"
	"mysrtafes-backend/pkg/game/platform"
	"mysrtafes-backend/pkg/game/tag"
	"net/http"
	"time"
)

// TODO: Laravelに合わせてるのであまりにも過剰(単純な移行なため仕様合わせでこうなってる。)
// TODO: あと、呼び出し側のことを考えるとjsonはキャメルケースのほうがよさそう。
type TagResponse struct {
	ID          tag.ID          `json:"id"`
	Name        tag.Name        `json:"name"`
	Description tag.Description `json:"description"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}
type PlatformResponse struct {
	ID          platform.ID          `json:"id"`
	Name        platform.Name        `json:"name"`
	Description platform.Description `json:"description"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}
type GameResponse struct {
	ID          game.ID            `json:"id"`
	Name        game.Name          `json:"name"`
	Description game.Description   `json:"description"`
	Publisher   game.Publisher     `json:"publisher"`
	Developer   game.Developer     `json:"developer"`
	Platforms   []PlatformResponse `json:"platforms"`
	Tags        []TagResponse      `json:"tags"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}
type GoalResponse struct {
	ID          goal.ID          `json:"id"`
	Name        goal.Name        `json:"name"`
	Description goal.Description `json:"description"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}
type DetailResponse struct {
	ID           detail.ID         `json:"id"`
	GameMasterID game.ID           `json:"game_master_id"`
	GameName     game.Name         `json:"game_name"`
	Game         GameResponse      `json:"game"`
	Goals        []GoalResponse    `json:"goal_genres"`
	GoalDetail   detail.GoalDetail `json:"goal_detail"`
	Department   detail.Department `json:"department"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}
type StreamStatusResponse struct {
	ID            stream.ID            `json:"id"`
	ChallengeID   challenges.ID        `json:"challenge_id"`
	IsLive        stream.IsLive        `json:"is_live"`
	TotalLiveTime stream.TotalLiveTime `json:"total_live_time"`
	Title         stream.Title         `json:"title"`
	StreamURL     stream.LiveURL       `json:"stream_url"`
	Thumbnail     stream.Thumbnail     `json:"thumbnail"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
}
type ChallengeResponse struct {
	ID               challenges.ID          `json:"id"`
	Name             challenges.Name        `json:"name"`
	NameRead         challenges.ReadingName `json:"name_read"`
	Twitter          challenges.Twitter     `json:"twitter"`
	Discord          challenges.Discord     `json:"discord"`
	IsStream         challenges.IsStream    `json:"is_stream"`
	StreamURL        string                 `json:"stream_url"`
	Comment          challenges.Comment     `json:"comment"`
	StreamStatus     *StreamStatusResponse  `json:"stream_status"`
	StreamSite       string                 `json:"stream_site"`
	ChallengeDetails []DetailResponse       `json:"challenge_details"`
	CreatedAt        time.Time              `json:"created_at"`
	UpdatedAt        time.Time              `json:"updated_at"`
}

func WriteCreateChallenge(w http.ResponseWriter, challenge *challenges.Challenge) error {
	body := struct {
		Code    int               `json:"code"`
		Message string            `json:"message"`
		Data    ChallengeResponse `json:"data"`
	}{
		Code:    http.StatusCreated,
		Message: "success create challenge",
		Data:    createChallengeResponse(challenge),
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(&body)
}

func createChallengeResponse(challenge *challenges.Challenge) ChallengeResponse {
	details := make([]DetailResponse, 0, len(challenge.Detail))
	for _, detailData := range challenge.Detail {
		detail := DetailResponse{
			ID:           detailData.ID,
			GameMasterID: detailData.Game.ID,
			GameName:     detailData.Game.Name,
			GoalDetail:   detailData.GoalDetail,
			Department:   detailData.Department,
		}
		for _, goalData := range detailData.Goals {
			detail.Goals = append(
				detail.Goals,
				GoalResponse{
					ID:          goalData.ID,
					Name:        goalData.Name,
					Description: goalData.Description,
				},
			)
		}
		details = append(details, detail)
	}

	var streamStatus *StreamStatusResponse
	if challenge.Stream.Status != nil {
		streamStatus = &StreamStatusResponse{
			ID:            challenge.Stream.Status.ID,
			ChallengeID:   challenge.ID,
			IsLive:        challenge.Stream.Status.IsLive,
			TotalLiveTime: challenge.Stream.Status.Detail.TotalLiveTime,
			Title:         challenge.Stream.Status.Detail.Title,
			StreamURL:     challenge.Stream.Status.Detail.LiveURL,
			Thumbnail:     challenge.Stream.Status.Detail.Thumbnail,
		}
	}
	return ChallengeResponse{
		ID:               challenge.ID,
		Name:             challenge.Challenger.Name,
		NameRead:         challenge.Challenger.ReadingName,
		Twitter:          challenge.SNS.Twitter,
		Discord:          challenge.SNS.Discord,
		IsStream:         challenge.Stream.IsStream,
		StreamURL:        challenge.Stream.URL.URL().String(),
		Comment:          challenge.Comment,
		StreamStatus:     streamStatus,
		StreamSite:       challenge.Stream.URL.StreamSite().String(),
		ChallengeDetails: details,
	}
}
