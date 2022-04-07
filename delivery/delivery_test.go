package delivery_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	StickerHttp "cleanarch/delivery"
	"cleanarch/domain"
	mocks "cleanarch/mocks"
)

// Test code for Sticker Handler

func TestGetTrendingStickers(t *testing.T) {

	mockSticker := &domain.Sticker{
		ID: 1, Name: "name1", StartTime: "11:00:00",
		EndTime: "20:00:00", CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}

	mockUCase := new(mocks.StickerUseCase)
	mockListSticker := make([]*domain.Sticker, 0)
	mockListSticker = append(mockListSticker, mockSticker)

	mockUCase.On("GetTrendingStickers").Return(mockListSticker, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/v1/trendingsticker", strings.NewReader(""))
	//err = errors.New("hello")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := StickerHttp.Delivery{
		UseCase: mockUCase,
	}
	err = handler.GetTrendingStickers(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

// func TestFetchError(t *testing.T) {

// 	mockUCase := new(mocks.StickerUseCase)
// 	mockUCase.On("GetTrendingStickers").Return(nil, "Error")

// 	e := echo.New()
// 	req, err := http.NewRequest(echo.GET, "/v1/trendingstickers", strings.NewReader(""))
// 	assert.NoError(t, err)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	handler := StickerHttp.Delivery{
// 		UseCase: mockUCase,
// 	}
// 	err = handler.GetTrendingStickers(c)
// 	require.NoError(t, err)

// 	assert.Equal(t, http.StatusInternalServerError, rec.Code)
// 	mockUCase.AssertExpectations(t)
// }
