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
func AddAnimalsPhotos(ctx context.Context, animalID, photoID int64) error {
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

func GetPhotosByProductID(ctx context.Context, productID int64) ([]string, error) {
	rows, err := mioDB.QueryContext(ctx, `
SELECT P.filename 
 FROM photos AS P 
  INNER JOIN products_photos AS PH ON PH.productID = ?
  AND P.id = PH.photoID`, productID)
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

func AddNewsPhotos(ctx context.Context, newsID, photoID int64) error {
	_, err := mioDB.ExecContext(ctx, `
INSERT INTO news_photos
(newsID, photoID)
VALUES(?, ?);
`, newsID, photoID)
	return err
}

func AddServicePhotos(ctx context.Context, serviceID, photoID int64) error {
	_, err := mioDB.ExecContext(ctx, `
INSERT INTO service_photos
(serviceID, photoID)
VALUES(?, ?);
`, serviceID, photoID)
	return err
}

func AddProductPhotos(ctx context.Context, productID, photoID int64) error {
	_, err := mioDB.ExecContext(ctx, `
INSERT INTO products_photos
(productID, photoID)
VALUES(?, ?);
`, productID, photoID)
	return err
}

func RemoveServicePhotos(ctx context.Context, serviceID int64) error {
	_, err := mioDB.ExecContext(ctx, `
DELETE FROM service_photos
		WHERE  serviceID = ?
`, serviceID)
	return err
}

func RemoveProductPhotos(ctx context.Context, productID int64) error {
	_, err := mioDB.ExecContext(ctx, `
DELETE FROM products_photos
		WHERE  productID = ?
`, productID)
	return err
}

func RemoveNewsPhotos(ctx context.Context, newsID int64) error {
	_, err := mioDB.ExecContext(ctx, `
DELETE FROM news_photos
		WHERE  newsID = ?
`, newsID)
	return err
}
