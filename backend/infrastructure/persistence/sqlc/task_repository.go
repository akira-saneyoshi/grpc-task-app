package sqlc

import (
	"context"
	"strconv"

	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/domain/object/value"
	"github.com/akira-saneyoshi/task-app/infrastructure/persistence/model/db"
	"github.com/akira-saneyoshi/task-app/utils/convert"
)

// SQLC implementation of task persistence.
type SQLCTaskRepository struct {
	db.Querier
}

func NewSQLCTaskRepository(qry db.Querier) *SQLCTaskRepository {
	return &SQLCTaskRepository{qry}
}

func (r *SQLCTaskRepository) FindTaskByID(ctx context.Context, id string) (*entity.Task, error) {
	res, err := r.Querier.FindTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.Task{
		ID:          value.NewID(res.ID),
		UserID:      value.NewID(res.UserID),
		Title:       res.Title,
		Description: convert.NullStringToPointer(res.Description),
		Status:      convert.NullTasksStatusToTaskStatus(res.Status),
		DueDate:     convert.NullTimeToTimePointer(res.DueDate),
		CreatedAt:   convert.ConvertNullTime(res.CreatedAt),
		UpdatedAt:   convert.ConvertNullTime(res.UpdatedAt),
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
			Description: convert.NullStringToPointer(v.Description),
			Status:      convert.NullTasksStatusToTaskStatus(v.Status),
			DueDate:     convert.NullTimeToTimePointer(v.DueDate),
			CreatedAt:   convert.ConvertNullTime(v.CreatedAt),
			UpdatedAt:   convert.ConvertNullTime(v.UpdatedAt),
		}
	}
	return tasks, nil
}

func (r *SQLCTaskRepository) CreateTask(ctx context.Context, arg *entity.Task) (string, error) {

	params := db.CreateTaskParams{
		ID:          arg.ID.Value(),
		UserID:      arg.UserID.Value(),
		Title:       arg.Title,
		Description: convert.StringToPointerNullString(arg.Description),
		Status:      convert.TaskStatusToNullTasksStatus(arg.Status),
		DueDate:     convert.TimeToPointerNullTime(arg.DueDate),
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
		Description: convert.StringToPointerNullString(arg.Description),
		Status:      convert.TaskStatusToNullTasksStatus(arg.Status),
		DueDate:     convert.TimeToPointerNullTime(arg.DueDate),
	}
	return r.Querier.UpdateTask(ctx, params)
}

func (r *SQLCTaskRepository) DeleteTask(ctx context.Context, id string) error {
	return r.Querier.DeleteTask(ctx, id)
}
