package request

import "time"

type ReportRequest struct {
	StartDate string `form:"start"`
	EndDate   string `form:"end"`
}

type ReportQuery struct {
	StartDate time.Time `json:"start"`
	EndDate   time.Time `json:"end"`
}

func (r ReportRequest) ToQuery() (*ReportQuery, error) {
	var start time.Time
	var end *time.Time
	if r.StartDate != "" {
		from, err := time.Parse("02-01-2006", r.StartDate)
		if err != nil {
			return nil, err
		}
		start = from
	} else {
		from, err := time.Parse("02-01-2006", "01-01-2000")
		if err != nil {
			return nil, err
		}
		start = from
	}
	
	if r.EndDate != "" {
		to, err := time.Parse("02-01-2006", r.EndDate)
		if err != nil {
			return nil, err
		}
		end = &to
	} else {
		now := time.Now()
		end = &now
	}

	return &ReportQuery{
		StartDate: start,
		EndDate: *end,
	}, nil
}