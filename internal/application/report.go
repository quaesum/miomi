package application

import (
	"context"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
)

type ReportApplication interface {
	Create(c context.Context, report *entity.ReportCreateRequest, uID int64) (int64, error)
	GetAll(c context.Context) ([]entity.Report, error)
	GetByID(c context.Context, id int64) (*entity.Report, error)
	Remove(c context.Context, id int64) error
}

type ReportApp struct {
	ReportApplication
}

func (r *ReportApp) GetAll(ctx context.Context) ([]entity.Report, error) {
	return mysql.GetAllReports(ctx)
}

func (r *ReportApp) GetByID(ctx context.Context, id int64) (*entity.Report, error) {
	return mysql.GetReportByID(ctx, id)
}

func (r *ReportApp) Remove(ctx context.Context, id int64) error {
	return mysql.RemoveReport(ctx, id)
}

func (r *ReportApp) Create(ctx context.Context, report *entity.ReportCreateRequest, uID int64) (int64, error) {
	return mysql.CreateReport(ctx, report, uID)
}
