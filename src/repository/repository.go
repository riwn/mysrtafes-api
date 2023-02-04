package repository

import (
	"errors"
	"mysrtafes-backend/pkg/challenge"
	"mysrtafes-backend/pkg/game/tag"
	mysrtafes_backend "mysrtafes-backend/repository/models/mysrtafes-backend"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

type Repository interface {
	challenge.Repository
	// stream.Repository
	// detail.Repository
	// goal.Repository
	// result.Repository
	// game.Repository
	// link.Repository
	// platform.Repository
	tag.Repository
	Close() error
}

func New(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) ChallengeCreate(*challenge.Challenge) (*challenge.Challenge, error) {
	return nil, errors.New("not implemented ChallengeCreate")
}

func (r *repository) ChallengeRead(challenge.ID) (*challenge.Challenge, error) {
	return nil, errors.New("not implemented ChallengeRead")
}

func (r *repository) ChallengeUpdate(*challenge.Challenge) (*challenge.Challenge, error) {
	return nil, errors.New("not implemented ChallengeUpdate")
}

func (r *repository) ChallengeDelete(challenge.ID) error {
	return errors.New("not implemented ChallengeDelete")
}

func (r *repository) TagCreate(tag *tag.Tag) (*tag.Tag, error) {
	model := mysrtafes_backend.NewTagMaster(tag)
	r.DB.Transaction(func(tx *gorm.DB) error {
		err := model.Create(tx)
		if err != nil {
			return err
		}
		return nil
	})
	return model.NewEntity(), nil
}

func (r *repository) TagRead(tag.ID) (*tag.Tag, error) {
	return nil, errors.New("not implemented TagRead")
}

func (r *repository) TagUpdate(*tag.Tag) (*tag.Tag, error) {
	return nil, errors.New("not implemented TagUpdate")
}

func (r *repository) TagDelete(tag.ID) error {
	return errors.New("not implemented TagDelete")
}

func (r *repository) Close() error {
	db, err := r.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
