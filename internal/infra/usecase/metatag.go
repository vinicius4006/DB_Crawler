package usecase

import (
	model "db_crawler/internal/entity"
	"errors"
	"reflect"
)

type MetaTagUseCase struct {
	metatagRepository model.MetaTagRepository
}

func NewMetaTagUseCase(metatagRepository model.MetaTagRepository) *MetaTagUseCase {
	return &MetaTagUseCase{metatagRepository: metatagRepository}
}

func (m *MetaTagUseCase) ExecuteCreate(input model.MetaTag) (uint64, error) {
	err := m.metatagRepository.Create(&input)
	if err != nil {
		return 0, err
	}

	primaryKey := reflect.ValueOf(&input).Elem().FieldByName("ID").Uint()
	return primaryKey, nil
}

func (m *MetaTagUseCase) ExecuteFindBySiteID(id uint64) ([]*model.MetaTag, error) {
	metatags, err := m.metatagRepository.FindBySiteID(id)
	if err != nil {
		return metatags, err
	}
	return metatags, nil
}

func (m *MetaTagUseCase) ExecuteFindByTag(tag string, siteid uint64) ([]*model.MetaTag, error) {
	if len(tag) == 0 {
		return []*model.MetaTag{}, errors.New("Tag is empty")
	}
	metatags, err := m.metatagRepository.FindByTag(tag, siteid)
	if err != nil {
		return metatags, err
	}
	return metatags, nil
}

func (m *MetaTagUseCase) ExecuteFindAll() ([]*model.MetaTag, error) {
	metatags, err := m.metatagRepository.FindAll()
	if err != nil {
		return metatags, nil
	}
	return metatags, nil
}
