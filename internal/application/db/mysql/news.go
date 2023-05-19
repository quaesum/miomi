package mysql

import (
	"context"
	"madmax/internal/entity"
)

func GetNewsInfo(ctx context.Context) ([]entity.News, error) {
	rows, err := mioDB.QueryContext(ctx, `
SELECT N.id, N.label, N.description, P.filename
FROM news AS N
  INNER JOIN photos AS P 
  LEFT JOIN news_photos AS NP ON N.id = NP.newsID AND P.id = NP.photoID
WHERE N.id = NP.newsID
GROUP BY 
    N.id, 
    N.label, 
    N.description, 
    P.filename
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
			return nil, err
		}
		news = append(news, newOne)
	}
	return news, nil
}
