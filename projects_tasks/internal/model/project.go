package model

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

type Project struct {
	ID           uint64
	Name         string
	Priority     string
	Img          string
	Description  string
	CompanyId    uint64
	ReporterId   uint64
	AssigneesIds []uint64
	StartsAt     time.Time
	DeadlineAt   timestamp.Timestamp
	UpdatedAt    timestamp.Timestamp
}

type ProjectInput struct {
	Name         string
	Priority     string
	Img          *string
	Description  string
	CompanyId    uint64
	ReporterId   uint64
	AssigneesIds []uint64
	DeadlineAt   *timestamp.Timestamp
}
