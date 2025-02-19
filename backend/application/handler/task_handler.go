package handler

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/akira-saneyoshi/task-app/application"
	"github.com/akira-saneyoshi/task-app/application/usecase"
	"github.com/akira-saneyoshi/task-app/domain"
	"github.com/akira-saneyoshi/task-app/interfaces/dto"
	dto_task "github.com/akira-saneyoshi/task-app/interfaces/dto/dto_task"
	task_v1 "github.com/akira-saneyoshi/task-app/interfaces/proto/task/v1"
	"github.com/akira-saneyoshi/task-app/utils/contextkey"
	"github.com/akira-saneyoshi/task-app/utils/convert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskHandler struct {
	usecase.ITaskUsecase
	contextkey.IContextReader
}

func NewTaskHandler(uc usecase.ITaskUsecase, cr contextkey.IContextReader) *TaskHandler {
	return &TaskHandler{uc, cr}
}

func (h *TaskHandler) GetTaskList(ctx context.Context, arg *connect.Request[task_v1.GetTaskListRequest]) (*connect.Response[task_v1.GetTaskListResponse], error) {
	var uid string
	var err error
	if uid, err = h.IContextReader.GetUserID(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	res, err := h.ITaskUsecase.FindTasksByUserID(ctx, dto.NewIDParam(uid))
	if err != nil {
		switch e := err.(type) {
		case *application.ErrInputValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrQueryFailed:
			return nil, connect.NewError(connect.CodeAborted, e)
		default:
			return nil, connect.NewError(connect.CodeUnknown, e)
		}
	}
	tasks := make([]*task_v1.Task, len(res))
	for i, v := range res {
		task := &task_v1.Task{
			Id:          v.ID.Value(),
			UserId:      v.UserID.Value(),
			Title:       v.Title,
			Description: v.Description,
			Status:      convert.ConvertStatus(v.Status),
			CreatedAt:   timestamppb.New(v.CreatedAt),
			UpdatedAt:   timestamppb.New(v.UpdatedAt),
		}
		if v.DueDate != nil {
			task.DueDate = timestamppb.New(*v.DueDate)
		}
		tasks[i] = task
	}
	return connect.NewResponse(&task_v1.GetTaskListResponse{
		Tasks: tasks,
	}), nil
}

func (h *TaskHandler) CreateTask(ctx context.Context, arg *connect.Request[task_v1.CreateTaskRequest]) (*connect.Response[task_v1.CreateTaskResponse], error) {
	var uid string
	var err error
	if uid, err = h.IContextReader.GetUserID(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	dueDateStr := arg.Msg.GetDueDate()
	var dueDate *time.Time
	if dueDateStr != "" {
		parsedDueDate, parseErr := time.Parse(time.RFC3339, dueDateStr)
		if parseErr != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, parseErr)
		}
		dueDate = &parsedDueDate
	}

	description := arg.Msg.GetDescription()
	var descriptionPtr *string = nil
	if description != "" {
		descriptionPtr = &description
	}

	createdID, err := h.ITaskUsecase.CreateTask(ctx, dto_task.NewCreateTaskParams(
		uid,
		arg.Msg.GetTitle(),
		descriptionPtr,
		arg.Msg.GetStatus(),
		dueDate,
	))
	if err != nil {
		switch e := err.(type) {
		case *application.ErrInputValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrQueryFailed:
			return nil, connect.NewError(connect.CodeAborted, e)
		default:
			return nil, connect.NewError(connect.CodeUnknown, e)
		}
	}
	return connect.NewResponse(&task_v1.CreateTaskResponse{
		CreatedId: createdID,
	}), nil
}

func (h *TaskHandler) UpdateTaskDetails(ctx context.Context, arg *connect.Request[task_v1.UpdateTaskDetailsRequest]) (*connect.Response[task_v1.UpdateTaskDetailsResponse], error) {
	var uid string
	var err error
	if uid, err = h.IContextReader.GetUserID(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	dueDateStr := arg.Msg.GetDueDate()
	var dueDate *time.Time
	if dueDateStr != "" {
		parsedDueDate, parseErr := time.Parse(time.RFC3339, dueDateStr)
		if parseErr != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, parseErr)
		}
		dueDate = &parsedDueDate
	}

	descriptionPtr := stringToPtr(arg.Msg.GetDescription())

	if err := h.ITaskUsecase.UpdateTaskDetails(ctx, dto_task.NewUpdateTaskDetailsParams(
		arg.Msg.GetTaskId(),
		uid,
		arg.Msg.GetTitle(),
		descriptionPtr,
		arg.Msg.GetDueDate(),
		dueDate,
	)); err != nil {
		switch e := err.(type) {
		case *application.ErrInputValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrNotFound:
			return nil, connect.NewError(connect.CodeNotFound, e)
		case *domain.ErrPermissionDenied:
			return nil, connect.NewError(connect.CodePermissionDenied, e)
		case *domain.ErrQueryFailed:
			return nil, connect.NewError(connect.CodeAborted, e)
		default:
			return nil, connect.NewError(connect.CodeUnknown, e)
		}
	}
	return connect.NewResponse(&task_v1.UpdateTaskDetailsResponse{}), nil
}

func (h *TaskHandler) DeleteTask(ctx context.Context, arg *connect.Request[task_v1.DeleteTaskRequest]) (*connect.Response[task_v1.DeleteTaskResponse], error) {
	var uid string
	var err error
	if uid, err = h.IContextReader.GetUserID(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	if err := h.ITaskUsecase.DeleteTask(ctx, dto.NewIDParam(arg.Msg.TaskId), dto.NewIDParam(uid)); err != nil {
		switch e := err.(type) {
		case *application.ErrInputValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrNotFound:
			return nil, connect.NewError(connect.CodeNotFound, e)
		case *domain.ErrPermissionDenied:
			return nil, connect.NewError(connect.CodePermissionDenied, e)
		case *domain.ErrQueryFailed:
			return nil, connect.NewError(connect.CodeAborted, e)
		default:
			return nil, connect.NewError(connect.CodeUnknown, e)
		}
	}
	return connect.NewResponse(&task_v1.DeleteTaskResponse{}), nil

}

func stringToPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
