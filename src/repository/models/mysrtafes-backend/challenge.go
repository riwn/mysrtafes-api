package mysrtafes_backend

import (
	"mysrtafes-backend/pkg/challenge"
	"mysrtafes-backend/pkg/errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MysChallenge2Challenges interface {
	Create(*gorm.DB) error
	Read(db *gorm.DB) error
	Update(db *gorm.DB) error
	Delete(db *gorm.DB) error
	NewEntity() (*challenge.Challenge, error)
}

type mysChallenge2Challenges struct {
	ID        challenge.ID `gorm:"primaryKey;autoIncrement"`
	Name      challenge.Name
	NameRead  challenge.ReadingName
	Password  challenge.Password
	IsStream  bool
	StreamURL string
	Twitter   challenge.Twitter
	Discord   challenge.Discord
	Comment   challenge.Comment
	Details   []*mysChallenge2ChallengeDetails `gorm:"foreignKey:ChallengeID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewMysChallenge2Challenges(challenge *challenge.Challenge) (MysChallenge2Challenges, error) {
	hashPassword, err := challenge.Challenger.Password.Hash()
	if err != nil {
		return nil, err
	}
	details := NewMysChallenge2ChallengeDetailsList(challenge)
	return &mysChallenge2Challenges{
		Name:      challenge.Challenger.Name,
		NameRead:  challenge.Challenger.ReadingName,
		Password:  hashPassword,
		IsStream:  bool(challenge.Stream.IsStream),
		StreamURL: challenge.Stream.URL.URL().String(),
		Twitter:   challenge.SNS.Twitter,
		Discord:   challenge.SNS.Discord,
		Comment:   challenge.Comment,
		Details:   details.newModels(),
	}, nil
}

func NewNewMysChallenge2ChallengesFromID(id challenge.ID) MysChallenge2Challenges {
	return &mysChallenge2Challenges{
		ID: id,
	}
}

func (mysChallenge2Challenges) TableName() string {
	return "mys_challenge2_challenges"
}

func (m *mysChallenge2Challenges) Create(db *gorm.DB) error {
	result := db.Omit("Details.GameMaster").Create(m)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBCreateError,
				result.Error.Error(),
				nil,
			),
			"create mys_challenge2_challenges error",
		)
	}
	return nil
}

func (m *mysChallenge2Challenges) Read(db *gorm.DB) error {
	// TODO: GoalGenreやStatus, Result等もPreloadで取得する予定
	result := db.Preload("Details.GameMaster").Preload(clause.Associations).
		Where("id = ?", m.ID).
		Find(&m)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBReadError,
				result.Error.Error(),
				nil,
			),
			"read mys_challenge2_challenges error",
		)
	}

	return nil
}

func (m *mysChallenge2Challenges) Update(db *gorm.DB) error {
	return nil
}

func (m *mysChallenge2Challenges) Delete(db *gorm.DB) error {
	return nil
}

func (m *mysChallenge2Challenges) NewEntity() (*challenge.Challenge, error) {
	var url challenge.URL
	var err error
	if m.StreamURL != "" {
		url, err = challenge.NewURL(m.StreamURL)
		if err != nil {
			return nil, err
		}
	}
	details := NewWrapMysChallenge2ChallengeDetailsList(m.Details)
	detailEntities, err := details.NewEntities()
	if err != nil {
		return nil, err
	}
	return &challenge.Challenge{
		ID: m.ID,
		Challenger: challenge.Challenger{
			Name:        m.Name,
			ReadingName: m.NameRead,
			Password:    m.Password,
		},
		Stream: challenge.Stream{
			IsStream: challenge.IsStream(m.IsStream),
			URL:      url,
		},
		SNS: challenge.SNS{
			Discord: m.Discord,
			Twitter: m.Twitter,
		},
		Details:   detailEntities,
		Comment:   m.Comment,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}, nil
}
