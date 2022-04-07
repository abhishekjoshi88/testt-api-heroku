package usecase_test

import (
	"cleanarch/domain"
	"cleanarch/mocks"
	"testing"
	"time"

	usecase "cleanarch/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetTrendingStickers(t *testing.T) {
	mockStickerRepo := new(mocks.StickerRepository)
	mockSticker := &domain.Sticker{
		ID: 1, Name: "name1", StartTime: "11:00:00",
		EndTime: "20:00:00", CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}

	mockListSticker := make([]*domain.Sticker, 0)
	mockListSticker = append(mockListSticker, mockSticker)

	t.Run("success", func(t *testing.T) {
		mockStickerRepo.On("GetTrendingStickers", mock.AnythingOfType("string")).Return(mockListSticker, nil).Once()

		u := usecase.StickerConstructor(mockStickerRepo)

		list, err := u.GetTrendingStickers()
		//err = errors.New("hello")
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListSticker))

		mockStickerRepo.AssertExpectations(t)

	})

	// t.Run("error-failed", func(t *testing.T) {
	// 	mockStickerRepo.On("GetTrendingStickers", mock.AnythingOfType("string")).Return(nil, errors.New("Unexpexted Error")).Once()

	// 	u := usecase.StickerConstructor(mockStickerRepo)

	// 	list, err := u.GetTrendingStickers()

	// 	assert.Error(t, err)
	// 	assert.Len(t, list, 0)
	// 	mockStickerRepo.AssertExpectations(t)
	// })
}
