package application

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/samber/lo"
	"madmax/internal/application/db/bleve"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
	"sort"
	"strconv"
)

func ProductCreate(ctx context.Context, product *entity.ProductCreateRequest) (int64, error) {
	productID, err := mysql.CreateProduct(ctx, product)
	if err != nil {
		return 0, err
	}

	for _, photoID := range product.Photos {
		err = mysql.AddProductPhotos(ctx, productID, photoID)
		if err != nil {
			return 0, err
		}
	}

	productInfo, err := ProductByID(ctx, productID)
	if err != nil {
		return 0, err
	}

	productBleve := entity.InsertProductReqToCreate(*productInfo)
	err = bleve.AddProduct(strconv.Itoa(int(productID)), productBleve)
	if err != nil {
		return 0, err
	}

	return productID, nil
}

func ProductByID(ctx context.Context, id int64) (*entity.Product, error) {
	service, err := mysql.GetProductInfo(ctx, id)
	if err != nil && err == sql.ErrNoRows {
		fmt.Println(err)
		return nil, errors.New("service not exist")
	}
	photos, err := mysql.GetPhotosByProductID(ctx, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	service.Photos = photos
	return service, nil
}

func RemoveProductByID(ctx context.Context, id int64) error {
	err := mysql.RemoveProductByID(ctx, id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	err = bleve.DeleteProduct(strconv.Itoa(int(id)))
	if err != nil {
		return err
	}
	return nil
}

func ProductUpdate(ctx context.Context, productID int64, productData *entity.ProductCreateRequest) error {
	err := mysql.UpdateProduct(ctx, productID, productData)
	if err != nil {
		return err
	}

	err = mysql.RemoveProductPhotos(ctx, productID)
	if err != nil {
		return err
	}
	for _, photoID := range productData.Photos {
		err = mysql.AddProductPhotos(ctx, productID, photoID)
		if err != nil {
			return err
		}
	}

	productInfo, err := ProductByID(ctx, productID)
	if err != nil {
		return err
	}
	productBleve := entity.InsertProductReqToCreate(*productInfo)
	err = bleve.AddProduct(strconv.Itoa(int(productID)), productBleve)
	if err != nil {
		return err
	}
	return nil
}

func GetProductsFromBleve(searchQuery string) ([]entity.ProductSearch, error) {
	res, err := bleve.SearchProducts(searchQuery)
	if err != nil {
		return nil, err
	}
	var products []entity.ProductSearch
	for _, item := range res.Hits {
		result := item.Fields
		product := entity.ProductSearch{
			ID:          item.ID,
			Name:        result["name"].(string),
			Description: result["description"].(string),
			Photos:      result["photos"].([]string),
			Link:        result["link"].(string),
		}
		products = append(products, product)
	}
	return products, err
}

func GetProductsSearchResult(searchTerm string, products []entity.Product) ([]entity.Product, error) {
	searchTerm = cleanQuery(searchTerm)

	for i := range products {
		products[i].Score = calculateProductsScore(products[i], searchTerm)
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].Score > products[j].Score
	})

	products = lo.Filter(products, func(service entity.Product, _ int) bool {
		return service.Score > 0
	})

	return products, nil
}
