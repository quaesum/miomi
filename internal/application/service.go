package application

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/samber/lo"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
	"sort"
)

func ServiceCreate(ctx context.Context, userID int64, service *entity.ServiceCreateRequest) (int64, error) {
	serviceID, err := mysql.CreateService(ctx, userID, service)
	if err != nil {
		return 0, err
	}

	for _, photoID := range service.Photos {
		err = mysql.AddServicePhotos(ctx, serviceID, photoID)
		if err != nil {
			return 0, err
		}
	}
	return serviceID, nil
}

func ServiceByID(ctx context.Context, id int64) (*entity.Service, error) {
	service, err := mysql.GetServiceInfo(ctx, id)
	if err != nil && err == sql.ErrNoRows {
		fmt.Println(err)
		return nil, errors.New("service not exist")
	}
	photos, err := mysql.GetPhotosByServiceID(ctx, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	service.Photos = photos
	return service, nil
}

func RemoveServiceByID(ctx context.Context, id int64) error {
	err := mysql.RemoveServiceByID(ctx, id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func ServiceUpdate(ctx context.Context, userID, serviceID int64, serviceData *entity.ServiceCreateRequest) error {
	_, err := mysql.GetUserByID(ctx, userID)
	if err != nil && err != sql.ErrNoRows {
		return errors.New("user exist")
	}
	err = mysql.UpdateService(ctx, userID, serviceID, serviceData)
	if err != nil {
		return err
	}

	err = mysql.RemoveServicePhotos(ctx, serviceID)
	if err != nil {
		return err
	}
	for _, photoID := range serviceData.Photos {
		err = mysql.AddServicePhotos(ctx, serviceID, photoID)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetAllServices(ctx context.Context) ([]entity.Service, error) {
	services, err := mysql.GetAllServices(ctx)
	if err != nil {
		return nil, err
	}
	return services, err
}

func GetServicesSearchResult(searchTerm string, services []entity.Service) ([]entity.Service, error) {
	searchTerm = cleanQuery(searchTerm)

	for i := range services {
		services[i].Score = calculateServicesScore(services[i], searchTerm)
	}

	sort.Slice(services, func(i, j int) bool {
		return services[i].Score > services[j].Score
	})

	services = lo.Filter(services, func(service entity.Service, _ int) bool {
		return service.Score > 0
	})

	return services, nil
}
