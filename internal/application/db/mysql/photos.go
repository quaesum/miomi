package mysql

import (
	"context"
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
