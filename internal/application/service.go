package application

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"madmax/internal/application/db/bleve"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
	"madmax/internal/utils"
	"strconv"
)

type ServiceApplication struct {
	bleve *bleve.ServiceBleve
}

func NewServiceApplication() *ServiceApplication {
	return &ServiceApplication{
		bleve: bleve.NewSerivceBleve(),
	}
}

func (s *ServiceApplication) Create(ctx context.Context, userID int64, service *entity.ServiceCreateRequest) (int64, error) {
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

	serviceInfo, err := s.GetByID(ctx, serviceID)
	if err != nil {
		return 0, err
	}

	productBleve := entity.InsertServiceReqToCreate(*serviceInfo)
	err = s.bleve.Add(strconv.Itoa(int(serviceID)), productBleve)
	if err != nil {
		return 0, err
	}

	return serviceID, nil
}

func (s *ServiceApplication) GetByID(ctx context.Context, id int64) (*entity.Service, error) {
	service, err := mysql.GetServiceInfo(ctx, id)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("service not exist")
	}
	return service, nil
}

func (s *ServiceApplication) Remove(ctx context.Context, id int64) error {
	err := mysql.RemoveServiceByID(ctx, id)

	if err != nil {
		return err
	}

	err = s.bleve.Remove(strconv.Itoa(int(id)))
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceApplication) Update(ctx context.Context, userID, serviceID int64, serviceData *entity.ServiceCreateRequest) error {
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
	serviceInfo, err := s.GetByID(ctx, serviceID)
	if err != nil {
		return err
	}

	productBleve := entity.InsertServiceReqToCreate(*serviceInfo)
	err = s.bleve.Add(strconv.Itoa(int(serviceID)), productBleve)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceApplication) GetFromBleve(searchQuery string) ([]entity.ServiceBleve, error) {
	res, err := s.bleve.Search(searchQuery)
	if err != nil {
		return nil, err
	}
	var services []entity.ServiceBleve
	for _, item := range res.Hits {
		result := item.Fields
		id, err := strconv.ParseInt(item.ID, 10, 64)
		if err != nil {
			return nil, err
		}
		service := entity.ServiceBleve{
			ID:          id,
			VolunteerID: result["volunteer_id"].(float64),
			Name:        result["name"].(string),
			Description: result["description"].(string),
		}
		service.Photos, err = utils.ProcessPhotos(result["photos"])
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, err
}

func (s *ServiceApplication) GetAllFromBleve() ([]entity.ServiceBleve, error) {
	res, err := s.bleve.SearchWOQuery()
	if err != nil {
		return nil, err
	}
	var services []entity.ServiceBleve
	for _, item := range res.Hits {
		result := item.Fields
		id, err := strconv.ParseInt(item.ID, 10, 64)
		if err != nil {
			return nil, err
		}
		service := entity.ServiceBleve{
			ID:          id,
			VolunteerID: result["volunteer_id"].(float64),
			Name:        result["name"].(string),
			Description: result["description"].(string),
		}
		service.Photos, err = utils.ProcessPhotos(result["photos"])
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, err
}
