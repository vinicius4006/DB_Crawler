package usecase

import (
	model "db_crawler/internal/entity"
	"errors"
	"reflect"
)

type SiteUseCase struct {
	siteRepository model.SiteRepository
}

func NewSiteUseCase(siteRepository model.SiteRepository) *SiteUseCase {
	return &SiteUseCase{siteRepository: siteRepository}
}

func (c *SiteUseCase) ExecuteCreate(input model.Site) (uint64, error) {
	err := c.siteRepository.Create(&input)
	if err != nil {
		return 0, err
	}
	primaryKey := reflect.ValueOf(&input).Elem().FieldByName("ID").Uint()
	return primaryKey, nil
}

func (c *SiteUseCase) ExecuteFindByID(id uint64) (*model.Site, error) {
	site, err := c.siteRepository.FindByID(id)
	if err != nil {
		return &model.Site{}, err
	}
	return site, nil
}

func (c *SiteUseCase) ExecuteFindByURL(url string) ([]*model.Site, error) {
	if len(url) == 0 {
		return []*model.Site{}, errors.New("URL is empty")
	}
	sites, err := c.siteRepository.FindByURL(url)
	if err != nil {
		return sites, err
	}
	return sites, nil
}

func (c *SiteUseCase) ExecuteFindAll() ([]*model.Site, error) {
	sites, err := c.siteRepository.FindAll()
	if err != nil {
		return sites, nil
	}
	return sites, nil
}
