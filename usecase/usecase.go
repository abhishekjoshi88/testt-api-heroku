package usecase

import (
	"cleanarch/config"
	"cleanarch/domain"

	"log"
	"time"
)

type useCase struct {
	repository domain.StickerRepository
}

// StickerConstructor will create an object that represent the StickerUsecase interface
func StickerConstructor(repository domain.StickerRepository) domain.StickerUseCase {
	return &useCase{repository: repository}
}

func (useCase useCase) GetTrendingStickers() (result []*domain.Sticker, err error) {
	now := time.Now()
	currentTime := now.Format(config.Time)

	stickers, err := useCase.repository.GetTrendingStickers(currentTime)

	if err != nil {

		log.Fatal(err)
	}
	for _, sticker := range stickers {
		result = append(result, sticker)
	}
	return
}
