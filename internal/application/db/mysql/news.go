package mysql

import (
	"context"
	"madmax/internal/entity"
	"time"
)

func GetNewsInfo(ctx context.Context) ([]entity.News, error) {
	rows, err := mioDB.QueryContext(ctx, `
SELECT N.id, N.label, N.description, P.filename, N.created_at
FROM news AS N
  INNER JOIN photos AS P 
  LEFT JOIN news_photos AS NP ON N.id = NP.newsID AND P.id = NP.photoID
WHERE N.id = NP.newsID
GROUP BY 
    N.id, 
    N.label, 
    N.description, 
    P.filename,
    N.created_at
`)
	if err != nil {
		return nil, err
	}
	var news []entity.News
	for rows.Next() {
		var newOne entity.News
		var createdAt int64
		err = rows.Scan(
			&newOne.ID,
			&newOne.Label,
			&newOne.Description,
			&newOne.Photo,
			&createdAt,
		)
		newOne.CreatedAt = time.Unix(createdAt, 0).Format("02.01.2006")
		if err != nil {
			return nil, err
		}
		news = append(news, newOne)
	}
	return news, nil
}

func CreateNews(ctx context.Context, news *entity.NewsCreateRequest) (int64, error) {
	res, err := mioDB.ExecContext(ctx, `
INSERT INTO news
	SET label = ?,
	description = ?,
	created_at = UNIX_TIMESTAMP()
	`, news.Label, news.Description)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func AddNewsPhoto(ctx context.Context, newsId, photoId int64) error {
	_, err := mioDB.ExecContext(ctx, `
	INSERT INTO news_photos
	(newsId, photoId)
	VALUES (?,?);
`, newsId, photoId)
	return err
}

func RemoveNewsByID(ctx context.Context, newsId int64) error {
	_, err := mioDB.ExecContext(ctx, `
DELETE FROM miomi.news
WHERE id = ?
`, newsId)
	if err != nil {
		return err
	}
	return nil
}
