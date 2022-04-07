package repository_test

import (
	"testing"
	"time"

	"gorm.io/driver/mysql"

	"cleanarch/domain"

	_repo "cleanarch/repository"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/gorm"
)

// Test code for Sticker Repository
func TestGetTrendingStickers(t *testing.T) { // Testing the Fetch function with all trending stickers

	DB, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      DB,
		SkipInitializeWithVersion: true,
	})
	db, err := gorm.Open(dialector, &gorm.Config{})

	defer DB.Close()

	mockStickers := []*domain.Sticker{
		{
			ID: 1, Name: "name1", StartTime: "11:00:00",
			EndTime: "20:00:00", CreatedAt: time.Now(), UpdatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "StartTime", "EndTime"}).
		AddRow(mockStickers[0].ID, mockStickers[0].Name, mockStickers[0].StartTime, mockStickers[0].EndTime)

	query := "[SELECT * FROM stickers DESC LIMIT ?]"

	mock.ExpectQuery(query).WillReturnRows(rows)

	now := time.Now()
	currentTime := now.Format("15:04:05")

	a := _repo.StickerConstructor(db)
	list, err := a.GetTrendingStickers(currentTime)

	assert.NoError(t, err)
	assert.Len(t, list, 1)

}
