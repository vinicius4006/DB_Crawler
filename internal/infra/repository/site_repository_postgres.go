package repository

import (
	model "db_crawler/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type SiteRepositoryPostgres struct {
	db *gorm.DB
}

func NewSiteRepositoryPostgres(db *gorm.DB) *SiteRepositoryPostgres {
	return &SiteRepositoryPostgres{db: db}
}

func (s *SiteRepositoryPostgres) Create(site *model.Site) error {
	result := s.db.Create(&site)
	return result.Error
}

func (s *SiteRepositoryPostgres) FindByID(id uint64) (*model.Site, error) {
	var site model.Site
	result := s.db.First(&site, id)
	if result != nil {
		return &site, result.Error
	}

	return &site, nil
}
func (s *SiteRepositoryPostgres) FindByURL(url string) ([]*model.Site, error) {
	var sites []*model.Site

	result := s.db.Where("url LIKE ?", fmt.Sprintf("%%%s%%", url)).Find(&sites)

	if result.Error != nil {
		return sites, result.Error
	}
	return sites, nil
}

func (s *SiteRepositoryPostgres) FindAll() ([]*model.Site, error) {
	var sites []*model.Site
	result := s.db.Find(&sites)
	if result.Error != nil {
		return sites, result.Error
	}
	return sites, nil
}
