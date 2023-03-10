package mysrtafes_backend

import (
	"mysrtafes-backend/pkg/errors"
	"mysrtafes-backend/pkg/game/platform"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PlatformMaster interface {
	Create(*gorm.DB) error
	Read(db *gorm.DB) error
	Update(db *gorm.DB) error
	Delete(db *gorm.DB) error
	NewEntity() *platform.Platform
}

type platformMaster struct {
	ID          platform.ID `gorm:"primaryKey;autoIncrement"`
	Name        platform.Name
	Description platform.Description
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewPlatformMaster(platform *platform.Platform) PlatformMaster {
	return &platformMaster{
		ID:          platform.ID,
		Name:        platform.Name,
		Description: platform.Description,
	}
}

func NewPlatformMasterFromID(platformID platform.ID) PlatformMaster {
	return &platformMaster{
		ID: platformID,
	}
}

func NewPlatformListFromIDs(platformIDs []platform.ID) []*platformMaster {
	platforms := make([]*platformMaster, 0, len(platformIDs))
	for _, platformID := range platformIDs {
		platforms = append(platforms, &platformMaster{
			ID: platformID,
		})
	}
	return platforms
}

func (platformMaster) TableName() string {
	return "platform_masters"
}

func (p *platformMaster) Create(db *gorm.DB) error {
	result := db.Create(p)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBCreateError,
				result.Error.Error(),
				nil,
			),
			"create platform_masters error",
		)
	}
	return nil
}

func (p *platformMaster) Read(db *gorm.DB) error {
	result := db.Where("id = ?", p.ID).Find(&p)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBReadError,
				result.Error.Error(),
				nil,
			),
			"read platform_masters error",
		)
	}
	return nil
}

func (p *platformMaster) Update(db *gorm.DB) error {
	// TODO: 更新の時だけCreatedAtが入ってこない問題がある。
	result := db.Updates(p)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBUpdateError,
				result.Error.Error(),
				nil,
			),
			"update platform_masters error",
		)
	}
	return nil
}

func (p *platformMaster) Delete(db *gorm.DB) error {
	result := db.Delete(p)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBDeleteError,
				result.Error.Error(),
				nil,
			),
			"delete platform_masters error",
		)
	}
	return nil
}

func (p *platformMaster) NewEntity() *platform.Platform {
	return &platform.Platform{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

type PlatformMasters interface {
	Find(db *gorm.DB, findOption *platform.FindOption) error
	NewEntities() []*platform.Platform
}

type platformMasters []*platformMaster

func NewPlatformMasters() PlatformMasters {
	return platformMasters{}
}

func (p platformMasters) Find(db *gorm.DB, findOption *platform.FindOption) error {
	// 検索モードで調整
	switch findOption.SearchMode {
	case platform.SearchMode_Pagination:
		db = db.Limit(findOption.Pagination.Limit).Offset(findOption.Pagination.Offset)
	case platform.SearchMode_Seek:
		db = db.Where("id > ?", findOption.Seek.LastID).Limit(findOption.Seek.Count)
	}

	switch findOption.OrderOption.Order {
	case platform.Order_Name:
		db.Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}, Desc: findOption.OrderOption.Desc})
	}

	result := db.Find(&p)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBReadError,
				result.Error.Error(),
				nil,
			),
			"find platform_masters error",
		)
	}
	return nil
}

func (p platformMasters) NewEntities() []*platform.Platform {
	entities := make([]*platform.Platform, 0, len(p))
	for _, model := range p {
		entities = append(entities, model.NewEntity())
	}
	return entities
}
