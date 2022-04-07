package delivery

import (
	"cleanarch/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	UseCase domain.StickerUseCase
}

func StickerConstructor(e *echo.Echo, useCase domain.StickerUseCase) {
	handler := &Delivery{
		UseCase: useCase,
	}

	e.GET("/v1/trendingStickers", handler.GetTrendingStickers)
}

func (delivery *Delivery) GetTrendingStickers(ctx echo.Context) error {
	sticker, err := delivery.UseCase.GetTrendingStickers()
	if err == nil {
		return ctx.JSON(http.StatusOK, sticker)
	}

	return ctx.JSON(http.StatusInternalServerError, domain.UnexpectedError)
}

//TODO MAKE FILE
