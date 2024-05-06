package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"madmax/internal/entity"
)

func GetServiceInfo(ctx context.Context, serviceID int64) (*entity.Service, error) {
	row := mioDB.QueryRowContext(ctx, `
SELECT 
  S.id, 
  S.volunteer_id, 
  S.label, 
  S.description, 
  S.created_at, 
  IFNULL(S.deleted_at, 0) AS deleted_at, 
  IFNULL(A.updated_at, 0) AS updated_at
FROM 
  serives AS S 
WHERE 
  S.id = ?
GROUP BY 
  S.id, 
  S.volunteer_id, 
  S.label, 
  S.description, 
  S.created_at,
  S.deleted_at,
  S.updated_at`, serviceID)
	service := new(entity.Service)
	err := row.Scan(
		&service.ID,
		&service.VolunteerID,
		&service.Label,
		&service.Description,
		&service.CreatedAt,
		&service.DeletedAt,
		&service.UpdatedAt,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	photos, err := GetPhotosByServiceID(ctx, service.ID)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}
	service.Photos = photos
	return service, err
}

func GetAllServices(ctx context.Context) ([]entity.Service, error) {
	rows, err := mioDB.QueryContext(ctx, `
SELECT 
  S.id, 
  S.volunteer_id, 
  S.label, 
  S.description, 
  S.created_at, 
  IFNULL(S.deleted_at, 0) AS deleted_at, 
  IFNULL(S.updated_at, 0) AS updated_at
FROM 
  serives AS S 
GROUP BY 
  S.id, 
  S.volunteer_id, 
  S.label, 
  S.description, 
  S.created_at,
  S.deleted_at,
  S.updated_at
`)

	if err != nil {
		return nil, err
	}
	var services []entity.Service
	for rows.Next() {
		var service entity.Service
		err = rows.Scan(
			&service.ID,
			&service.VolunteerID,
			&service.Label,
			&service.Description,
			&service.CreatedAt,
			&service.DeletedAt,
			&service.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		photos, err := GetPhotosByServiceID(ctx, service.ID)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(err)
		}
		service.Photos = photos
		services = append(services, service)
	}

	return services, nil
}

func CreateService(ctx context.Context, userID int64, service *entity.ServiceCreateRequest) (int64, error) {
	res, err := mioDB.ExecContext(ctx, `
INSERT INTO serives  
		SET  volunteer_id = ?,
		  	label = ?,
 			description = ?,
   			created_at = UNIX_TIMESTAMP()
`, userID, service.Label, service.Description)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func RemoveServiceByID(ctx context.Context, serviceID int64) error {
	_, err := mioDB.ExecContext(ctx, `
DELETE FROM serives
WHERE id = ?

`, serviceID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateService(ctx context.Context, userID, serviceID int64, service *entity.ServiceCreateRequest) error {
	_, err := mioDB.ExecContext(ctx, `
UPDATE serives
   SET  volunteer_id = ?,
		label = ?,
 		description = ?,
        updated_at = UNIX_TIMESTAMP()
 WHERE id = ? 
`, userID, service.Label, service.Description, serviceID)
	if err != nil {
		return err
	}
	return nil
}

func GetServicesCount() (int64, error) {
	rows, err := mioDB.Query("SELECT count(*) FROM serives")
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
