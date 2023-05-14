package model

import (
	"gorm.io/gorm"
)

type SiteRepository interface {
	Create(site *Site) error
	FindByID(id uint64) (*Site, error)
	FindByURL(url string) ([]*Site, error)
	FindAll() ([]*Site, error)
}
type MetaTagRepository interface {
	Create(metatag *MetaTag) error
	FindBySiteID(id uint64) ([]*MetaTag, error)
	FindByTag(url string, siteid uint64) ([]*MetaTag, error)
	FindAll() ([]*MetaTag, error)
}
type WordRepository interface {
	Create(word *Word) error
	FindBySiteID(id uint64) ([]*Word, error)
	FindByValue(url string, siteid uint64) ([]*Word, error)
	FindAll() ([]*Word, error)
}

type Site struct {
	gorm.Model
	Url  string
	Body []byte `gorm:"type:bytea"`
}

type MetaTag struct {
	gorm.Model
	SiteID  uint64
	Ref     Site `gorm:"foreignKey:SiteID"`
	Tag     string
	Content []byte `gorm:"type:bytea"`
}

type Word struct {
	gorm.Model
	SiteID  uint64
	Ref     Site   `gorm:"foreignKey:SiteID"`
	Value   string `gorm:"type:text"`
	Counter uint64
}
