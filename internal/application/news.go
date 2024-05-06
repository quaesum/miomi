package application

import (
	"context"
	"database/sql"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
)

func NewsAll(ctx context.Context) ([]entity.News, error) {
	return mysql.GetNewsInfo(ctx)
}

func NewsCreate(ctx context.Context, newsData *entity.NewsCreateRequest) (int64, error) {
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

func RemoveNewsById(ctx context.Context, id int64) error {
	err := mysql.RemoveNewsByID(ctx, id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func NewsUpdate(ctx context.Context, animalID int64, newsData *entity.NewsCreateRequest) error {
	err := mysql.UpdateNews(ctx, animalID, newsData)
	if err != nil {
		return err
	}

	err = mysql.RemoveNewsPhotos(ctx, animalID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	err = mysql.AddNewsPhotos(ctx, animalID, newsData.Photo)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}
