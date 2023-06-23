package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	pb "projects_tasks/proto/gen"
	"time"
)

type Project struct {
	db *sqlx.DB
}

func NewProject(db *sqlx.DB) *Project {
	return &Project{
		db: db,
	}
}

func (p Project) CreateProject(project *pb.CreateProjectRequest) (uint64, error) {
	var projectId uint64
	query := `INSERT INTO projects 
				(name, priority, img, description, company_id, reporter_id, assignees_ids, deadline_at) 
			  VALUES 
				($1, $2, $3, $4, $5, $6, $7, $8) 
			  RETURNING id`

	deadlineAt := time.Unix(project.DeadlineAt.GetSeconds(), int64(project.DeadlineAt.GetNanos()))

	err := p.db.QueryRowx(query,
		project.Name,
		project.Priority,
		project.Img,
		project.Description,
		project.CompanyId,
		project.ReporterId,
		pq.Array(project.AssigneesIds),
		deadlineAt,
	).Scan(&projectId)
	if err != nil {
		return 0, err
	}

	return projectId, nil
}
