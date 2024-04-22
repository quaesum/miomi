package application

import (
	"context"
	"database/sql"
	"errors"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
	"math"
)

func GetAnimalsOnCurrentPage(req entity.AnimalsRequest, animals []entity.Animal) ([]entity.Animal, error) {
	page := req.Page
	perPage := req.PerPage
	page -= 1
	if page < 0 {
		return nil, errors.New("page must be greater than 0")
	}

	leftBorder := page * perPage
	rightBorder := page*perPage + perPage
	if rightBorder > int8(len(animals)) {
		rightBorder = int8(len(animals))
	}

	return animals[leftBorder:rightBorder], nil
}

func GetMaxPagesAnimals(length int, perPage int8) (int8, error) {
	pages := math.Ceil(float64(length) / float64(perPage))
	return int8(pages), nil
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

func RemoveNewsById(ctx context.Context, id int64) error {
	err := mysql.RemoveNewsByID(ctx, id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}
