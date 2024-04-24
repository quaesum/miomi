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
INSERT INTO photos
(filename, origin, file_type)
VALUES(?, '', '');
`, name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
func RemoveAnimalsPhotos(ctx context.Context, animalID int64) error {
	_, err := mioDB.ExecContext(ctx, `
DELETE FROM animals_photos
		WHERE  animalID = ?
`, animalID)
	return err
}
func AddanimalsPhotos(ctx context.Context, animalID, photoID int64) error {
	_, err := mioDB.ExecContext(ctx, `
INSERT INTO animals_photos
(animalID, photoID)
VALUES(?, ?);
`, animalID, photoID)
	return err
}

func GetUrlAndId(ctx context.Context) ([]entity.PhotoRequest, error) {
	rows, err := mioDB.QueryContext(ctx, `
SELECT P.id,P.filename
	FROM photos AS P
`)
	if err != nil {
		return nil, err
	}
	var res []entity.PhotoRequest
	for rows.Next() {
		var newOne entity.PhotoRequest
		err = rows.Scan(&newOne.ID, &newOne.Filename)
		if err != nil {
			return nil, err
		}
		res = append(res, newOne)
	}
	return res, nil
}

func GetPhotosByServiceID(ctx context.Context, serviceID int64) ([]string, error) {
	rows, err := mioDB.QueryContext(ctx, `
SELECT P.filename 
 FROM photos AS P 
  INNER JOIN service_photos AS PH ON PH.serviceID = ?
  AND P.id = PH.photoID`, serviceID)
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
