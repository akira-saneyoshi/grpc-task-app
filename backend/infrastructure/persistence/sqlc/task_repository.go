package sqlc

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/domain/object/value"
	"github.com/akira-saneyoshi/task-app/infrastructure/persistence/model/db"
	"github.com/akira-saneyoshi/task-app/utils/convert"
)

type SQLCTaskRepository struct {
	db.Querier
}

func NewSQLCTaskRepository(qry db.Querier) *SQLCTaskRepository {
	return &SQLCTaskRepository{qry}
}

func (r *SQLCTaskRepository) FindTaskByID(ctx context.Context, id string) (*entity.Task, error) {
	res, err := r.Querier.FindTaskByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &entity.Task{
		ID:          value.NewID(res.ID),
		UserID:      value.NewID(res.UserID),
		Title:       res.Title,
		Description: convert.ConvertNullString(res.Description),
		Status:      convert.ConvertNullStatus(res.Status),
		DueDate:     convert.ConvertNullTimePtr(res.DueDate),
		CreatedAt:   convert.ConvertNullTimeValue(res.CreatedAt),
		UpdatedAt:   convert.ConvertNullTimeValue(res.UpdatedAt),
	}, nil
}

func (r *SQLCTaskRepository) FindTasksByUserID(ctx context.Context, userID string) ([]*entity.Task, error) {
	res, err := r.Querier.FindTasksByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	tasks := make([]*entity.Task, len(res))
	for i, v := range res {
		tasks[i] = &entity.Task{
			ID:          value.NewID(v.ID),
			UserID:      value.NewID(v.UserID),
			Title:       v.Title,
			Description: convert.ConvertNullString(v.Description),
			Status:      convert.ConvertNullStatus(v.Status),
			DueDate:     convert.ConvertNullTimePtr(v.DueDate),
			CreatedAt:   convert.ConvertNullTimeValue(v.CreatedAt),
			UpdatedAt:   convert.ConvertNullTimeValue(v.UpdatedAt),
		}
	}
	return tasks, nil
}

func (r *SQLCTaskRepository) CreateTask(ctx context.Context, arg *entity.Task) (string, error) {
	params := db.CreateTaskParams{
		ID:          arg.ID.Value(),
		UserID:      arg.UserID.Value(),
		Title:       arg.Title,
		Description: convert.NewSQLNullString(arg.Description),
		Status: db.NullTasksStatus{
			TasksStatus: db.TasksStatus(arg.Status),
			Valid:       true,
		},
		DueDate: convert.NewSQLNullTime(arg.DueDate),
	}
	err := r.Querier.CreateTask(ctx, params)
	if err != nil {
		return "", err
	}

	lastInsertID, err := r.Querier.GetLastInsertID(ctx)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(lastInsertID, 10), nil
}

func (r *SQLCTaskRepository) UpdateTask(ctx context.Context, arg *entity.Task) error {
	params := db.UpdateTaskParams{
		ID:          arg.ID.Value(),
		Title:       arg.Title,
		Description: convert.NewSQLNullString(arg.Description),
		Status: db.NullTasksStatus{
			TasksStatus: db.TasksStatus(arg.Status),
			Valid:       true,
		},
		DueDate: convert.NewSQLNullTime(arg.DueDate),
	}
	return r.Querier.UpdateTask(ctx, params)
}

func (r *SQLCTaskRepository) DeleteTask(ctx context.Context, id string) error {
	return r.Querier.DeleteTask(ctx, id)
}
