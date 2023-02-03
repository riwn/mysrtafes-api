package repository

import (
	"errors"
	"mysrtafes-backend/pkg/challenge"

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
	// tag.Repository
	Close() error
}

func New(db *gorm.DB) Repository {
	return &repository{db}
}

func (r repository) ChallengeCreate(*challenge.Challenge) (*challenge.Challenge, error) {
	return nil, errors.New("not implemented ChallengeCreate")
}

func (r repository) ChallengeRead(challenge.ID) (*challenge.Challenge, error) {
	return nil, errors.New("not implemented ChallengeRead")
}

func (r repository) ChallengeUpdate(*challenge.Challenge) (*challenge.Challenge, error) {
	return nil, errors.New("not implemented ChallengeUpdate")
}

func (r repository) ChallengeDelete(challenge.ID) error {
	return errors.New("not implemented ChallengeDelete")
}

func (r *repository) Close() error {
	// TODO: DBのClose
	return nil
}
