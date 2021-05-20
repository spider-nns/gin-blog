package model

import (
	"github.com/jinzhu/gorm"
	"http-gin/pkg/app"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ? ", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
func (t Tag) List(db *gorm.DB, pageNo, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageNo >= 0 && pageSize > 0 {
		db = db.Offset(pageNo).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ? ", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, err
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}
func (t Tag) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(t).Where("id = ? AND is_del = ?", t.ID, 0).Update(values).Error; err != nil {
		return err
	}
	return nil
}
func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id =? ", t.Model.ID).Delete(&t).Error
}
