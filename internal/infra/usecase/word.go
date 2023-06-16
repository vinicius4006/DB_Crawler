package usecase

import (
	model "db_crawler/internal/entity"
	"errors"
	"reflect"
)

type WordUseCase struct {
	wordrepository model.WordRepository
}

func NewWordUseCase(wordrepository model.WordRepository) *WordUseCase {
	return &WordUseCase{wordrepository: wordrepository}
}

func (w *WordUseCase) ExecuteCreate(input model.Word) (uint64, error) {
	err := w.wordrepository.Create(&input)
	if err != nil {
		return 0, err
	}
	primaryKey := reflect.ValueOf(&input).Elem().FieldByName("ID").Uint()
	return primaryKey, nil
}

func (w *WordUseCase) ExecuteFindBySiteID(id uint64) ([]*model.Word, error) {
	words, err := w.wordrepository.FindBySiteID(id)
	if err != nil {
		return words, err
	}
	return words, nil
}

func (w *WordUseCase) ExecuteFindByValues(value string) ([]*model.Word, error) {
	if len(value) == 0 {
		return []*model.Word{}, errors.New("Value is empty")
	}
	words, err := w.wordrepository.FindByValues(value)
	if err != nil {
		return words, err
	}
	return words, nil
}

func (w *WordUseCase) ExecuteFindAll() ([]*model.Word, error) {
	words, err := w.wordrepository.FindAll()
	if err != nil {
		return words, nil
	}
	return words, nil
}

func (w *WordUseCase) ExecuteUpdate(word *model.Word) error {
	err := w.wordrepository.Update(word)
	if err != nil {
		return err
	}
	return nil
}
