package mysrtafes_backend

import (
	"mysrtafes-backend/pkg/challenge"
	"mysrtafes-backend/pkg/errors"

	"gorm.io/gorm"
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
	IsStream  challenge.IsStream
	StreamURL string
	Twitter   challenge.Twitter
	Discord   challenge.Discord
	Comment   challenge.Comment
}

func NewMysChallenge2Challenges(challenge *challenge.Challenge) (MysChallenge2Challenges, error) {
	hashPassword, err := challenge.Challenger.Password.Hash()
	if err != nil {
		return nil, err
	}
	return &mysChallenge2Challenges{
		Name:      challenge.Challenger.Name,
		NameRead:  challenge.Challenger.ReadingName,
		Password:  hashPassword,
		IsStream:  challenge.Stream.IsStream,
		StreamURL: challenge.Stream.URL.URL().String(),
		Twitter:   challenge.SNS.Twitter,
		Discord:   challenge.SNS.Discord,
		Comment:   challenge.Comment,
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
	result := db.Create(m)
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
	return nil
}

func (m *mysChallenge2Challenges) Update(db *gorm.DB) error {
	return nil
}

func (m *mysChallenge2Challenges) Delete(db *gorm.DB) error {
	return nil
}

func (m *mysChallenge2Challenges) NewEntity() (*challenge.Challenge, error) {
	url, err := challenge.NewURL(m.StreamURL)
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
			IsStream: m.IsStream,
			URL:      url,
		},
		SNS: challenge.SNS{
			Discord: m.Discord,
			Twitter: m.Twitter,
		},
		Comment: m.Comment,
	}, nil
}
