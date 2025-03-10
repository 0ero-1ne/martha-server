package services

import (
	"server/db"
	"server/models"
)

type TagService struct{}

func NewTagService() TagService {
	return TagService{}
}

func (service TagService) GetAll() ([]models.Tag, error) {
	var tags []models.Tag
	tx := db.GetDB().Find(&tags)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return tags, nil
}

func (service TagService) GetById(id uint) (models.Tag, error) {
	var tag models.Tag
	tx := db.GetDB().First(&tag, id)

	if tx.Error != nil {
		return tag, tx.Error
	}

	return tag, nil
}

func (service TagService) Create(tag models.Tag) (models.Tag, error) {
	tx := db.GetDB().Create(&tag)

	if tx.Error != nil {
		return tag, tx.Error
	}

	return tag, nil
}

func (service TagService) Update(id uint, newTag models.Tag) (models.Tag, error) {
	var tag models.Tag
	tx := db.GetDB().First(&tag, id)

	if tx.Error != nil {
		return tag, tx.Error
	}

	tag.Title = newTag.Title

	tx = db.GetDB().Save(&tag)

	if tx.Error != nil {
		return tag, tx.Error
	}

	return tag, nil
}

func (service TagService) Delete(id uint) error {
	var tag models.Tag
	tx := db.GetDB().First(&tag, id)

	if tx.Error != nil {
		return tx.Error
	}

	tx = db.GetDB().Delete(&tag)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
