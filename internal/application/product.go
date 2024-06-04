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
	"madmax/internal/utils"
	"sort"
	"strconv"
)

type ProductApplication struct {
	bleve *bleve.ProductBleve
}

func NewProductApplication() *ProductApplication {
	return &ProductApplication{
		bleve: bleve.NewProductBleve(),
	}
}

func (p *ProductApplication) Create(ctx context.Context, product *entity.ProductCreateRequest) (int64, error) {
	productID, err := mysql.CreateProduct(ctx, product)
	if err != nil {
		return 0, err
	}

	for _, photoLink := range product.Photos {
		err = mysql.AddProductPhotos(ctx, productID, photoLink)
		if err != nil {
			return 0, err
		}
	}

	productInfo, err := p.GetByID(ctx, productID)
	if err != nil {
		return 0, err
	}

	productBleve := entity.InsertProductReqToCreate(*productInfo)
	err = p.bleve.Add(strconv.Itoa(int(productID)), productBleve)
	if err != nil {
		return 0, err
	}

	return productID, nil
}

func (p *ProductApplication) GetByID(ctx context.Context, id int64) (*entity.Product, error) {
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

func (p *ProductApplication) Remove(ctx context.Context, id int64) error {
	err := mysql.RemoveProductByID(ctx, id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	err = p.bleve.Remove(strconv.Itoa(int(id)))
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductApplication) Update(ctx context.Context, productID int64, productData *entity.ProductCreateRequest) error {
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

	productInfo, err := p.GetByID(ctx, productID)
	if err != nil {
		return err
	}
	productBleve := entity.InsertProductReqToCreate(*productInfo)
	err = p.bleve.Add(strconv.Itoa(int(productID)), productBleve)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductApplication) GetFromBleve(req *entity.SearchRequest, limit int) ([]entity.ProductSearch, error) {
	res, err := p.bleve.Search(req, limit)
	if err != nil {
		return nil, err
	}
	var products []entity.ProductSearch
	for _, item := range res.Hits {
		result := item.Fields
		id, err := strconv.ParseInt(item.ID, 10, 64)
		if err != nil {
			return nil, err
		}
		product := entity.ProductSearch{
			ID:          id,
			Name:        result["name"].(string),
			Description: result["description"].(string),
			Link:        result["link"].(string),
		}
		product.Photos, err = utils.ProcessPhotos(result["photos"])
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, err
}

//func (p *ProductApplication) GetAllFromBleve() ([]entity.ProductSearch, error) {
//	res, err := p.bleve.SearchWOQuery()
//	if err != nil {
//		return nil, err
//	}
//	var products []entity.ProductSearch
//	for _, item := range res.Hits {
//		id, err := strconv.ParseInt(item.ID, 10, 64)
//		if err != nil {
//			return nil, err
//		}
//		result := item.Fields
//		product := entity.ProductSearch{
//			ID:          id,
//			Name:        result["name"].(string),
//			Description: result["description"].(string),
//			Link:        result["link"].(string),
//		}
//		product.Photos, err = utils.ProcessPhotos(result["photos"])
//		if err != nil {
//			return nil, err
//		}
//		products = append(products, product)
//	}
//	return products, err
//}

func GetProductsSearchResult(searchTerm string, products []entity.Product) ([]entity.Product, error) {
	searchTerm = utils.CleanQuery(searchTerm)

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
