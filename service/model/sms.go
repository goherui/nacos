package model

import (
	__ "day6/proto"

	"gorm.io/gorm"
)

type CmsContent struct {
	gorm.Model
	Title      string `gorm:"column:title;not null" json:"title"`
	CategoryID uint64 `gorm:"column:category_id;not null" json:"categoryId"`
	Content    string `gorm:"column:content;type:longtext;not null" json:"content"`
	Status     int8   `gorm:"column:status;default:0" json:"status"`
	ViewCount  uint   `gorm:"column:view_count;default:0" json:"viewCount"`
	Creator    uint64 `gorm:"column:creator;not null" json:"creator"`
}

func (CmsContent) TableName() string {
	return "cms_content"
}

func (c *CmsContent) FindTitle(db *gorm.DB, title string) error {
	return db.Where("title=?", title).First(&c).Error
}

func (c *CmsContent) CmsCreate(db *gorm.DB) error {
	return db.Create(&c).Error
}

func (c *CmsContent) FindId(db *gorm.DB, id int64) interface{} {
	return db.Where("id=?", id).First(&c).Error
}

func (c *CmsContent) CmsDel(db *gorm.DB, id int64) interface{} {
	return db.Delete(&c, id).Error
}

func (c *CmsContent) CmsUpdate(db *gorm.DB, id int64) error {
	return db.Where("id=?", id).Updates(&c).Error
}

func (c *CmsContent) FindContent(db *gorm.DB, in *__.CmsListReq) ([]__.Cms, error) {
	var list []__.Cms
	if in.Page != 0 || in.Size != 0 {
		offset := (in.Page - 1) * in.Size
		err := db.Model(&CmsContent{}).Select("cms_content.*,cms_category.category_name").
			Joins("LEFT JOIN cms_category ON cms_content.category_id = cms_category.id").
			Offset(int(offset)).Limit(int(in.Size)).Find(&list).Error

		return list, err
	}
	err := db.Model(&CmsContent{}).Select("cms_content.*,cms_category.category_name").
		Joins("LEFT JOIN cms_category ON cms_content.category_id = cms_category.id").
		Find(&list).Error
	return list, err
}
