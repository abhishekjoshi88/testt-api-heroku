package repository

import (
	"cleanarch/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// StickerConstructor will create an object that represent the StickerRepository interface
func StickerConstructor(db *gorm.DB) domain.StickerRepository {
	return &repository{db: db}
}

// region Sticker

func (repository repository) GetTrendingStickers(currentTime string) ([]*domain.Sticker, error) {
	//make slice
	trendingStickers := make([]*domain.Sticker, 0)
	err := repository.db.Table("stickers").
		Where("stickers.startTime <= ?", currentTime).
		Where("stickers.endTime >= ?", currentTime).
		Order("priority desc").
		Find(&trendingStickers).Error

	return trendingStickers, err
}
