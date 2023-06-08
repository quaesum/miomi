package mysql

import (
	"context"
	"madmax/internal/entity"
)

func GetPhotosByAnimalID(ctx context.Context, animalID int64) ([]string, error) {
	rows, err := mioDB.QueryContext(ctx, `
SELECT P.filename 
 FROM photos AS P 
  INNER JOIN animals_photos AS PH ON PH.animalID = ?
  AND P.id = PH.photoID`, animalID)
	if err != nil {
		return nil, err
	}
	var files []string
	for rows.Next() {
		var fileName string
		err := rows.Scan(
			&fileName,
		)
		if err != nil {
			return nil, err
		}
		files = append(files, fileName)
	}

	return files, err
}

func CreateFile(ctx context.Context, name string) (int64, error) {
	res, err := mioDB.ExecContext(ctx, `
INSERT INTO miomi.photos
(filename, origin, file_type)
VALUES(?, '', '');
`, name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetUrlByID(ctx context.Context, id int64) (*entity.PhotoRequest, error) {
	row := mioDB.QueryRowContext(ctx, `
SELECT P.filename
	FROM photos AS P
	WHERE P.id = ?`, id)
	res := new(entity.PhotoRequest)
	err := row.Scan(&res.Filename)
	if err != nil {
		return nil, err
	}
	return res, nil
}
