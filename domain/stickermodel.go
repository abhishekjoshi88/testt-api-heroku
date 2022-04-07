package domain

import "time"

// Table Names

const (
	StickerTable = "stickers"
)

//tables Structs

type Sticker struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Priority  int       `json:"priority"`
	StartTime string    `json:"-"`
	EndTime   string    `json:"-"`
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}

func (sticker *Sticker) TableName() string {
	return StickerTable
}

// interface Structs

type StickerUseCase interface {
	GetTrendingStickers() ([]*Sticker, error)
}

type StickerRepository interface {
	GetTrendingStickers(currentTime string) ([]*Sticker, error)
}
