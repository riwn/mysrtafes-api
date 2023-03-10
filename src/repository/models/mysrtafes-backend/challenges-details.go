package mysrtafes_backend

import (
	"mysrtafes-backend/pkg/challenge"
	"mysrtafes-backend/pkg/challenge/detail"
	"mysrtafes-backend/pkg/errors"
	"mysrtafes-backend/pkg/game"
	"time"

	"gorm.io/gorm"
)

type MysChallenge2ChallengeDetails interface {
	Update(db *gorm.DB) error
	Delete(db *gorm.DB) error
	NewEntity() (*detail.Detail, error)
}

type mysChallenge2ChallengeDetails struct {
	ID           detail.ID `gorm:"primaryKey;autoIncrement"`
	ChallengeID  challenge.ID
	GameMasterID game.ID
	GameMaster   gameMaster `gorm:"foreignKey:GameMasterID"`
	GoalDetail   detail.GoalDetail
	GameName     game.Name
	Department   detail.Department
	Goals        []*goalGenreMaster `gorm:"many2many:mys_challenge2_goal_genre_detail_links;"`
	// TODO: Resultの紐付け
	CreatedAt time.Time
	UpdatedAt time.Time
}

func newMysChallenge2ChallengeDetails(challengeID challenge.ID, detail *detail.Detail, goals []*goalGenreMaster) *mysChallenge2ChallengeDetails {
	// TODO: GoalGenreの紐付けを追加する
	return &mysChallenge2ChallengeDetails{
		ID:           detail.ID,
		ChallengeID:  challengeID,
		GameMasterID: detail.Game.ID,
		GoalDetail:   detail.GoalDetail,
		GameName:     detail.Game.Name,
		Department:   detail.Department,
		Goals:        goals,
	}
}

func (mysChallenge2ChallengeDetails) TableName() string {
	return "mys_challenge2_challenge_details"
}

func (g *mysChallenge2ChallengeDetails) Update(db *gorm.DB) error {
	return nil
}

func (g *mysChallenge2ChallengeDetails) Delete(db *gorm.DB) error {
	return nil
}

func (g *mysChallenge2ChallengeDetails) NewEntity() (*detail.Detail, error) {
	// TODO: GameNameの扱いをどうするか。
	game, err := g.GameMaster.NewEntity()
	if err != nil {
		return nil, err
	}
	// IDがないときはNameだけでも付与する
	if game.ID == 0 {
		game.Name = g.GameName
	}
	return &detail.Detail{
		ID:   g.ID,
		Game: *game,
		// TODO: Goals:      goals,
		GoalDetail: g.GoalDetail,
		Department: g.Department,
		// TODO: Result:     result,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}, nil
}

type MysChallenge2ChallengeDetailsList interface {
	Create(*gorm.DB) error
	Find(*gorm.DB) error
	NewEntities() ([]*detail.Detail, error)
	newModels() []*mysChallenge2ChallengeDetails
}

type mysChallenge2ChallengeDetailsList []*mysChallenge2ChallengeDetails

func NewMysChallenge2ChallengeDetailsList(challenge *challenge.Challenge) MysChallenge2ChallengeDetailsList {
	models := make(mysChallenge2ChallengeDetailsList, 0, len(challenge.Details))
	for _, challengeDetail := range challenge.Details {
		goalModels := NewGoalGenreMasterFromGoals(challengeDetail.Goals)
		models = append(models, newMysChallenge2ChallengeDetails(challenge.ID, challengeDetail, goalModels))
	}
	return models
}

func NewWrapMysChallenge2ChallengeDetailsList(details []*mysChallenge2ChallengeDetails) MysChallenge2ChallengeDetailsList {
	return mysChallenge2ChallengeDetailsList(details)
}

func NewEmptyMysChallenge2ChallengeDetailsList() mysChallenge2ChallengeDetailsList {
	return mysChallenge2ChallengeDetailsList{}
}

func (d mysChallenge2ChallengeDetailsList) BeforeCreate(db *gorm.DB) error {
	gameIDs := make([]game.ID, 0, len(d))
	existIDs := make(map[game.ID]struct{})
	for _, detail := range d {
		if _, ok := existIDs[detail.GameMasterID]; !ok {
			existIDs[detail.GameMasterID] = struct{}{}
			gameIDs = append(gameIDs, detail.GameMasterID)
		}
	}

	var gameModels []*gameMaster
	db.Where("id IN ?", gameIDs).Find(&gameModels)
	if len(gameModels) != len(gameIDs) {
		return errors.NewInvalidValidate(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBReadError,
				"game.id model is nothing error",
				[]errors.InvalidParams{
					errors.NewInvalidParams("game.id", gameIDs),
				},
			),
			"game.id model is nothing error",
		)
	}
	return nil
}

func (d mysChallenge2ChallengeDetailsList) Create(db *gorm.DB) error {
	// 中間テーブルの設定
	err := db.SetupJoinTable(&mysChallenge2ChallengeDetails{}, "Goals", detailGoalLink{})
	if err != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBTableJoinError,
				err.Error(),
				nil,
			),
			"set many to many table error",
		)
	}
	result := db.Omit("Goals.*").Create(d)
	if result.Error != nil {
		// FIXME: ここ、ModelのBeforeCreateのエラーが伝播できてない。
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBCreateError,
				result.Error.Error(),
				nil,
			),
			"create mys_challenge2_challenge_details error",
		)
	}
	return nil
}

func (d mysChallenge2ChallengeDetailsList) Find(db *gorm.DB) error {
	// TODO: ChallengeIDでの検索は必須。他にFindOptionをつけて取得できるようにするのかを要検討
	return nil
}

func (d mysChallenge2ChallengeDetailsList) NewEntities() ([]*detail.Detail, error) {
	entities := make([]*detail.Detail, 0, len(d))
	for _, model := range d {
		entity, err := model.NewEntity()
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}

func (d mysChallenge2ChallengeDetailsList) newModels() []*mysChallenge2ChallengeDetails {
	return d
}
