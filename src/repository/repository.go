package repository

import (
	"errors"
	"mysrtafes-backend/pkg/challenge"
	"mysrtafes-backend/pkg/game"
	"mysrtafes-backend/pkg/game/platform"
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
	game.Repository
	// link.Repository
	platform.Repository
	tag.Repository
	Close() error
}

func New(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) ChallengeCreate(c *challenge.Challenge) (*challenge.Challenge, error) {
	challenge, err := mysrtafes_backend.NewMysChallenge2Challenges(c)
	if err != nil {
		return nil, err
	}
	// TODO: GoalGenre紐付けのInsert
	err = r.DB.Transaction(func(tx *gorm.DB) error {
		err := challenge.Create(tx)
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return challenge.NewEntity()
}

func (r *repository) ChallengeRead(id challenge.ID) (*challenge.Challenge, error) {
	challenge := mysrtafes_backend.NewNewMysChallenge2ChallengesFromID(id)
	if err := challenge.Read(r.DB); err != nil {
		return nil, err
	}

	return challenge.NewEntity()
}

func (r *repository) ChallengeUpdate(*challenge.Challenge) (*challenge.Challenge, error) {
	return nil, errors.New("not implemented ChallengeUpdate")
}

func (r *repository) ChallengeDelete(challenge.ID) error {
	return errors.New("not implemented ChallengeDelete")
}

func (r *repository) TagCreate(tag *tag.Tag) (*tag.Tag, error) {
	model := mysrtafes_backend.NewTagMaster(tag)
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := model.Create(tx)
		if err != nil {
			return err
		}
		return nil
	})
	return model.NewEntity(), err
}

func (r *repository) TagRead(tagID tag.ID) (*tag.Tag, error) {
	model := mysrtafes_backend.NewTagMasterFromID(tagID)
	err := model.Read(r.DB)
	return model.NewEntity(), err
}

func (r *repository) TagFind(f *tag.FindOption) ([]*tag.Tag, error) {
	models := mysrtafes_backend.NewTagMasters()
	err := models.Find(r.DB, f)
	return models.NewEntities(), err
}

func (r *repository) TagUpdate(tag *tag.Tag) (*tag.Tag, error) {
	model := mysrtafes_backend.NewTagMaster(tag)
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := model.Update(tx)
		if err != nil {
			return err
		}
		return nil
	})
	return model.NewEntity(), err
}

func (r *repository) TagDelete(tagID tag.ID) error {
	model := mysrtafes_backend.NewTagMasterFromID(tagID)
	return model.Delete(r.DB)
}

func (r *repository) PlatformCreate(platform *platform.Platform) (*platform.Platform, error) {
	model := mysrtafes_backend.NewPlatformMaster(platform)
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := model.Create(tx)
		if err != nil {
			return err
		}
		return nil
	})
	return model.NewEntity(), err
}

func (r *repository) PlatformRead(platformID platform.ID) (*platform.Platform, error) {
	model := mysrtafes_backend.NewPlatformMasterFromID(platformID)
	err := model.Read(r.DB)
	return model.NewEntity(), err
}

func (r *repository) PlatformFind(p *platform.FindOption) ([]*platform.Platform, error) {
	models := mysrtafes_backend.NewPlatformMasters()
	err := models.Find(r.DB, p)
	return models.NewEntities(), err
}

func (r *repository) PlatformUpdate(platform *platform.Platform) (*platform.Platform, error) {
	model := mysrtafes_backend.NewPlatformMaster(platform)
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := model.Update(tx)
		if err != nil {
			return err
		}
		return nil
	})
	return model.NewEntity(), err
}

func (r *repository) PlatformDelete(platformID platform.ID) error {
	model := mysrtafes_backend.NewPlatformMasterFromID(platformID)
	return model.Delete(r.DB)
}

func (r *repository) GameCreate(game *game.Game, platformIDs []platform.ID, tagIDs []tag.ID) (*game.Game, error) {
	tags := mysrtafes_backend.NewTagMasterListFromIDs(tagIDs)
	platforms := mysrtafes_backend.NewPlatformListFromIDs(platformIDs)
	model := mysrtafes_backend.NewGameMaster(game, platforms, tags)
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := model.Create(tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return model.NewEntity()
}

func (r *repository) GameRead(gameID game.ID) (*game.Game, error) {
	model := mysrtafes_backend.NewGameMasterFromID(gameID)
	err := model.Read(r.DB)
	if err != nil {
		return nil, err
	}
	return model.NewEntity()
}

func (r *repository) GameFind(f *game.FindOption) ([]*game.Game, error) {
	models := mysrtafes_backend.NewGameMasters()
	err := models.Find(r.DB, f)
	if err != nil {
		return nil, err
	}
	return models.NewEntities()
}

func (r *repository) GameUpdate(game *game.Game, platformIDs []platform.ID, tagIDs []tag.ID) (*game.Game, error) {
	tags := mysrtafes_backend.NewTagMasterListFromIDs(tagIDs)
	platforms := mysrtafes_backend.NewPlatformListFromIDs(platformIDs)
	model := mysrtafes_backend.NewGameMaster(game, platforms, tags)
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := model.Update(tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return model.NewEntity()
}

func (r *repository) GameDelete(id game.ID) error {
	model := mysrtafes_backend.NewGameMasterFromID(id)
	return model.Delete(r.DB)
}

func (r *repository) Close() error {
	db, err := r.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
