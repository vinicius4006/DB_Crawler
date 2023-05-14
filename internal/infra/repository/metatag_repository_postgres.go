package repository

import (
	model "db_crawler/internal/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type MetaTagRepositoryPostgres struct {
	db *gorm.DB
}

func NewMetaTagRepositoryPostgres(db *gorm.DB) *MetaTagRepositoryPostgres {
	return &MetaTagRepositoryPostgres{db: db}
}

func (m *MetaTagRepositoryPostgres) Create(metatag *model.MetaTag) error {
	rows := m.db.Where("id = ?", &metatag.SiteID).First(&[]model.Site{}).RowsAffected
	if rows == 0 {
		return errors.New("Don't find site by id")
	}
	result := m.db.Create(&metatag)
	return result.Error
}

func (m *MetaTagRepositoryPostgres) FindBySiteID(id uint64) ([]*model.MetaTag, error) {
	var metatags []*model.MetaTag
	result := m.db.Where("site_id = ?", id).Find(&metatags)
	if result.Error != nil {
		return metatags, result.Error
	}

	return metatags, nil
}
func (m *MetaTagRepositoryPostgres) FindByTag(tag string) ([]*model.MetaTag, error) {
	var metatags []*model.MetaTag
	result := m.db.Where("tag LIKE ?", fmt.Sprintf("%%%s%%", tag)).Find(&metatags)

	if result.Error != nil {
		return metatags, result.Error
	}
	return metatags, nil
}

func (m *MetaTagRepositoryPostgres) FindAll() ([]*model.MetaTag, error) {
	var metatags []*model.MetaTag
	result := m.db.Find(&metatags)
	if result.Error != nil {
		return metatags, result.Error
	}
	return metatags, nil
}
