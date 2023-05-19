package mysql

import (
	"context"
	"fmt"
	"madmax/internal/entity"
)

func GetNewsInfo(ctx context.Context) ([]entity.News, error) {
	rows, err := mioDB.QueryContext(ctx, `
SELECT N.id, N.label, N.description, NP.photoID,
FROM 
    news AS N
	INNER JOIN news_photos AS NP
WHERE N.id = NP.newsID
GROUP BY 
    N.id, 
    N.label, 
    N.description, 
    NP.photoID
`)
	if err != nil {
		return nil, err
	}
	var news []entity.News
	for rows.Next() {
		var newOne entity.News
		err := rows.Scan(
			&newOne.ID,
			&newOne.Label,
			&newOne.Description,
			&newOne.Photo,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		news = append(news, newOne)
	}
	return news, nil
}
