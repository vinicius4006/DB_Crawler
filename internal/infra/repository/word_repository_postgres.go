package repository

import (
	model "db_crawler/internal/entity"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type WordRepositoryPostgres struct {
	db *gorm.DB
}

func NewWordRepositoryPostgres(db *gorm.DB) *WordRepositoryPostgres {
	return &WordRepositoryPostgres{db: db}
}

func (w *WordRepositoryPostgres) Create(word *model.Word) error {
	rows := w.db.Where("id = ?", &word.SiteID).First(&[]model.Site{}).RowsAffected
	if rows == 0 {
		return errors.New("Don't find site by id")
	}
	result := w.db.Create(&word)
	return result.Error
}

func (w *WordRepositoryPostgres) FindBySiteID(id uint64) ([]*model.Word, error) {
	var words []*model.Word
	result := w.db.Preload("Ref").Where("site_id = ?", id).Find(&words)
	if result.Error != nil {
		return words, result.Error
	}

	return words, nil
}
func (w *WordRepositoryPostgres) FindByValues(q string) ([]*model.Word, error) {
	var words []*model.Word
	removeSpaceQ := strings.Trim((q), " ")
	result := w.db.Preload("Ref").Where("value ILIKE ?", fmt.Sprintf("%%%s%%", removeSpaceQ)).Find(&words)
	fmt.Println("ESPACO", removeSpaceQ)
	if result.Error != nil {
		return words, result.Error
	}

	return words, nil
}

func (w *WordRepositoryPostgres) FindAll() ([]*model.Word, error) {
	var words []*model.Word
	result := w.db.Find(&words)
	if result.Error != nil {
		return words, result.Error
	}
	return words, nil
}

func (w *WordRepositoryPostgres) Update(word *model.Word) error {
	result := w.db.Model(word).Where("id = ? AND site_id = ?", word.ID, word.SiteID).Updates(&word)

	if result.Error != nil {
		return errors.New(fmt.Sprintf("Erro ao atualizar: %v", result.Error))
	}
	return nil
}
