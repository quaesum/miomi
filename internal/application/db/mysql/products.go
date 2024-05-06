package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"madmax/internal/entity"
)

func GetProductInfo(ctx context.Context, productID int64) (*entity.Product, error) {
	row := mioDB.QueryRowContext(ctx, `
SELECT 
  P.id, 
  P.name, 
  P.description,
  P.link,
  P.created_at, 
  IFNULL(P.deleted_at, 0) AS deleted_at, 
  IFNULL(P.updated_at, 0) AS updated_at
FROM 
  products AS P 
WHERE 
  P.id = ?
GROUP BY 
  P.id, 
  P.name, 
  P.description,
  P.link,
  P.created_at,
  P.deleted_at,
  P.updated_at`, productID)
	product := new(entity.Product)
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Link,
		&product.CreatedAt,
		&product.DeletedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	photos, err := GetPhotosByProductID(ctx, productID)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}
	product.Photos = photos
	return product, err
}

func GetAllProducts(ctx context.Context) ([]entity.Product, error) {
	rows, err := mioDB.QueryContext(ctx, `
SELECT 
  P.id, 
  P.name, 
  P.description,
  P.link,
  P.created_at, 
  IFNULL(P.deleted_at, 0) AS deleted_at, 
  IFNULL(P.updated_at, 0) AS updated_at
FROM 
  products AS P 
GROUP BY 
  P.id, 
  P.name, 
  P.description,
  P.link,
  P.created_at, 
  P.deleted_at,
  P.updated_at
`)

	if err != nil {
		return nil, err
	}
	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Link,
			&product.CreatedAt,
			&product.DeletedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		photos, err := GetPhotosByProductID(ctx, product.ID)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(err)
		}
		product.Photos = photos
		products = append(products, product)
	}

	return products, nil
}

func CreateProduct(ctx context.Context, product *entity.ProductCreateRequest) (int64, error) {
	res, err := mioDB.ExecContext(ctx, `
INSERT INTO products
		SET name = ?,
 			description = ?,
 			link = ?,
   			created_at = UNIX_TIMESTAMP()
`, product.Name, product.Description, product.Link)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func RemoveProductByID(ctx context.Context, productID int64) error {
	_, err := mioDB.ExecContext(ctx, `
DELETE FROM products
WHERE id = ?

`, productID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(ctx context.Context, productID int64, product *entity.ProductCreateRequest) error {
	_, err := mioDB.ExecContext(ctx, `
UPDATE products
   SET  name = ?,
 		description = ?,
 		link = ?,
        updated_at = UNIX_TIMESTAMP()
 WHERE id = ? 
`, product.Name, product.Description, product.Link, productID)
	if err != nil {
		return err
	}
	return nil
}

func GetProductsCount() (int64, error) {
	rows, err := mioDB.Query("SELECT count(*) FROM products")
	if err != nil {
		return 0, err
	}
	var count int64
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&count)
	}
	return count, nil
}
