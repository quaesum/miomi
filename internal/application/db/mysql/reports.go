package mysql

import (
	"context"
	"madmax/internal/entity"
)

func GetAllReports(ctx context.Context) ([]entity.Report, error) {
	rows, err := mioDB.QueryContext(ctx, `SELECT 
		r.id,
		r.senderID,
		r.label,
		r.description,
		v.firstName,
		v.lastName,
		v.email,
		v.createdAt
	FROM reports r
	JOIN volunteers v ON r.senderID = v.id`)
	if err != nil {
		return nil, err
	}

	var reports []entity.Report
	for rows.Next() {
		var report entity.Report
		var senderID, createdAt int64
		var firstName, lastName, email, label, description string
		if err = rows.Scan(&report.ID, &senderID, &label, &description, &firstName, &lastName, &email, &createdAt); err != nil {
			return nil, err
		}
		report.Label = label
		report.Description = description
		report.Sender = entity.UserResponse{
			ID:        senderID,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			CreatedAt: createdAt,
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func GetReportByID(ctx context.Context, reportID int64) (*entity.Report, error) {
	row := mioDB.QueryRowContext(ctx, `SELECT 
		r.id,
		r.senderID,
		r.label,
		r.description,
		v.firstName,
		v.lastName,
		v.email,
		v.createdAt
	FROM reports r
	JOIN volunteers v ON r.senderID = v.id WHERE r.id = ?`, reportID)
	report := new(entity.Report)
	err := row.Scan(&report.ID, &report.Sender.ID, &report.Label, &report.Description, &report.Sender.FirstName, &report.Sender.FirstName, &report.Sender.Email, &report.Sender.CreatedAt)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func CreateReport(ctx context.Context, report *entity.ReportCreateRequest, uID int64) (int64, error) {
	res, err := mioDB.ExecContext(ctx, `
	INSERT INTO reports
		SET senderID = ?,
		    label = ?,
		    description = ?
`, uID, report.Label, report.Description)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func RemoveReport(ctx context.Context, reportID int64) error {
	_, err := mioDB.ExecContext(ctx, `
DELETE FROM reports
WHERE id = ?
`, reportID)
	if err != nil {
		return err
	}
	return nil
}
