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

func (r *repository) TagRead(tagID tag.ID) (*tag.Tag, error) {
	model := mysrtafes_backend.NewTagMasterFromID(tagID)
	err := model.Read(r.DB)
	return model.NewEntity(), err
}

func (r *repository) TagFind(f *tag.FindOption) ([]*tag.Tag, error) {
	models := mysrtafes_backend.NewTagMasters()
	err := models.Find(r.DB, f)
	entities := make([]*tag.Tag, 0, len(models))
	for _, model := range models {
		entities = append(entities, model.NewEntity())
	}
	return entities, err
}

func (r *repository) TagUpdate(tag *tag.Tag) (*tag.Tag, error) {
	model := mysrtafes_backend.NewTagMaster(tag)
	r.DB.Transaction(func(tx *gorm.DB) error {
		err := model.Update(tx)
		if err != nil {
			return err
		}
		return nil
	})
	return model.NewEntity(), nil
}

func (r *repository) TagDelete(tagID tag.ID) error {
	model := mysrtafes_backend.NewTagMasterFromID(tagID)
	return model.Delete(r.DB)
}

func (r *repository) Close() error {
	db, err := r.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
