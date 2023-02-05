package mysrtafes_backend

import (
	"mysrtafes-backend/pkg/errors"
	"mysrtafes-backend/pkg/game/tag"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TagMaster interface {
	Create(*gorm.DB) error
	Read(db *gorm.DB) error
	Update(db *gorm.DB) error
	Delete(db *gorm.DB) error
	NewEntity() *tag.Tag
}

type tagMaster struct {
	ID          tag.ID `gorm:"primaryKey;autoIncrement"`
	Name        tag.Name
	Description tag.Description
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTagMaster(tag *tag.Tag) TagMaster {
	return &tagMaster{
		ID:          tag.ID,
		Name:        tag.Name,
		Description: tag.Description,
	}
}

func NewTagMasterFromID(tagID tag.ID) TagMaster {
	return &tagMaster{
		ID: tagID,
	}
}

func (tagMaster) TableName() string {
	return "tag_masters"
}

func (t *tagMaster) Create(db *gorm.DB) error {
	result := db.Create(t)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBCreateError,
				result.Error.Error(),
				nil,
			),
			"create tag_masters error",
		)
	}
	return nil
}

func (t *tagMaster) Read(db *gorm.DB) error {
	result := db.Where("id = ?", t.ID).Find(&t)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBReadError,
				result.Error.Error(),
				nil,
			),
			"read tag_masters error",
		)
	}
	return nil
}

func (t *tagMaster) Update(db *gorm.DB) error {
	// TODO: 更新の時だけCreatedAtがなぜか入ってこない問題があるっぽい。
	result := db.Updates(t)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBUpdateError,
				result.Error.Error(),
				nil,
			),
			"update tag_masters error",
		)
	}
	return nil
}

func (t *tagMaster) Delete(db *gorm.DB) error {
	result := db.Delete(t)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBDeleteError,
				result.Error.Error(),
				nil,
			),
			"delete tag_masters error",
		)
	}
	return nil
}

func (t *tagMaster) NewEntity() *tag.Tag {
	return &tag.Tag{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

type tagMasters []*tagMaster

func NewTagMasters() tagMasters {
	return []*tagMaster{}
}

func (t *tagMasters) Find(db *gorm.DB, findOption *tag.FindOption) error {
	// 検索モードで調整
	switch findOption.SearchMode {
	case tag.SearchMode_Pagination:
		db = db.Limit(findOption.Pagination.Limit).Offset(findOption.Pagination.Offset)
	case tag.SearchMode_Seek:
		db = db.Where("id > ?", findOption.Seek.LastID).Limit(findOption.Seek.Count)
	}

	switch findOption.OrderOption.Order {
	case tag.Order_Name:
		db.Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}, Desc: findOption.OrderOption.Desc})
	}

	result := db.Find(&t)
	if result.Error != nil {
		return errors.NewInternalServerError(
			errors.Layer_Model,
			errors.NewInformation(
				errors.ID_DBReadError,
				result.Error.Error(),
				nil,
			),
			"find tag_masters error",
		)
	}
	return nil
}
