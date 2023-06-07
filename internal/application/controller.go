package application

import (
	"context"
	"database/sql"
	"errors"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
)

func AnimalsAll(ctx context.Context) ([]entity.Animal, error) {
	return mysql.GetAllAnimals(ctx)
}

func NewsAll(ctx context.Context) ([]entity.News, error) {
	return mysql.GetNewsInfo(ctx)
}

func NewsCreate(ctx context.Context, userID int64, newsData *entity.NewsCreateRequest) (int64, error) {
	_, err := mysql.GetUserByID(ctx, userID)
	if err != nil && err != sql.ErrNoRows {
		return 0, errors.New("user exist")
	}
	newsId, err := mysql.CreateNews(ctx, newsData)
	if err != nil {
		return 0, err
	}
	err = mysql.AddNewsPhoto(ctx, newsId, newsData.Photo)
	if err != nil {
		return 0, err
	}
	return newsId, nil
}
