package service

import (
	"context"
	"testing"
	"time"

	"github.com/akira-saneyoshi/task-app/domain"
	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/domain/object/value"
	"github.com/akira-saneyoshi/task-app/domain/service"
	"github.com/akira-saneyoshi/task-app/tests/mocks"
	"github.com/stretchr/testify/require"
)

func TestTaskService_NewTaskService(tt *testing.T) {
	tt.Run("異常系: structがinterfaceを実装しているか", func(t *testing.T) {
		var _ service.ITaskService = (*service.TaskService)(nil)
	})
}

func TestTaskService_FindTaskByID(tt *testing.T) {
	ctx := context.Background()
	now := time.Now().UTC()
	id := "id"
	uid := "uid"
	dueDate := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	task := &entity.Task{
		ID:          value.NewID(id),
		UserID:      value.NewID(uid),
		Title:       "task",
		Description: ptrString("description"),
		Status:      entity.StatusPending,
		DueDate:     &dueDate,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tt.Run("正常系: 正しい入力の場合", func(t *testing.T) {
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(task, nil)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		srv := service.NewTaskService(repo, im, cm)
		ret, err := srv.FindTaskByID(ctx, id)

		require.NoError(t, err, "エラーが発生しないこと")
		require.Equal(t, id, ret.ID.Value())
		require.Equal(t, task.UserID.Value(), ret.UserID.Value())
		require.Equal(t, task.Title, ret.Title)
		require.Equal(t, task.Description, ret.Description)
		require.Equal(t, task.Status, ret.Status)
		require.Equal(t, *task.DueDate, *ret.DueDate)
		require.Equal(t, task.CreatedAt, ret.CreatedAt)
		require.Equal(t, task.UpdatedAt, ret.UpdatedAt)
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: TaskIDが空の場合", func(t *testing.T) {
		errExp := &domain.ErrValidationFailed{Msg: "id is empty"}
		id := ""
		repo := new(mocks.ITaskRepository)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		srv := service.NewTaskService(repo, im, cm)
		_, err := srv.FindTaskByID(ctx, id)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: 存在しないTaskIDの場合", func(t *testing.T) {
		errExp := &domain.ErrNotFound{Msg: "task not found"}
		id := "another"
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(nil, errExp)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		srv := service.NewTaskService(repo, im, cm)
		_, err := srv.FindTaskByID(ctx, id)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
}

func TestTaskService_FindTasksByUserID(tt *testing.T) {
	ctx := context.Background()
	now := time.Now().UTC()
	uid := "uid"
	dueDate1 := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	dueDate2 := time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC)
	tasks := []*entity.Task{
		{ID: value.NewID("t1"), UserID: value.NewID(uid), Title: "task1", Description: ptrString("desc1"), Status: entity.StatusPending, DueDate: &dueDate1, CreatedAt: now, UpdatedAt: now},
		{ID: value.NewID("t2"), UserID: value.NewID(uid), Title: "task2", Description: ptrString("desc2"), Status: entity.StatusInProgress, DueDate: &dueDate2, CreatedAt: now, UpdatedAt: now},
	}

	tt.Run("正常系: 正しい入力の場合", func(t *testing.T) {
		repo := new(mocks.ITaskRepository)
		repo.On("FindTasksByUserID", ctx, uid).Return(tasks, nil)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		srv := service.NewTaskService(repo, im, cm)
		ret, err := srv.FindTasksByUserID(ctx, uid)

		require.NoError(t, err, "エラーが発生しないこと")
		require.ElementsMatch(t, tasks, ret)
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: UserIDが空の場合", func(t *testing.T) {
		errExp := &domain.ErrValidationFailed{Msg: "id is empty"}
		uid := ""
		repo := new(mocks.ITaskRepository)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		srv := service.NewTaskService(repo, im, cm)
		_, err := srv.FindTasksByUserID(ctx, uid)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: 存在しないUserIDの場合", func(t *testing.T) {
		errExp := &domain.ErrQueryFailed{}
		uid := "another"
		repo := new(mocks.ITaskRepository)
		repo.On("FindTasksByUserID", ctx, uid).Return(nil, errExp)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		srv := service.NewTaskService(repo, im, cm)
		_, err := srv.FindTasksByUserID(ctx, uid)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
}

func TestTaskService_CreateTask(tt *testing.T) {
	ctx := context.Background()
	id := "id"
	uid := "uid"
	now := time.Now().UTC()
	task := &entity.Task{
		ID:          value.NewID(id),
		UserID:      value.NewID(uid),
		Title:       "task",
		Description: ptrString("description"),
		Status:      entity.StatusPending,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	afTitle := "task title after"
	description := ptrString("description")
	status := string(entity.StatusPending)
	var dueDate *time.Time

	tt.Run("正常系: 正しい入力の場合", func(t *testing.T) {
		repo := new(mocks.ITaskRepository)
		repo.On("CreateTask", ctx, task).Return(id, nil)
		im := new(mocks.IIDManager)
		im.On("GenerateID").Return(id)
		cm := new(mocks.IClockManager)
		cm.On("GetNow").Return(now)
		srv := service.NewTaskService(repo, im, cm)
		ret, err := srv.CreateTask(ctx, uid, afTitle, description, status, dueDate)

		require.NoError(t, err, "エラーが発生しないこと")
		require.Equal(t, id, ret)
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: 不正な入力の場合", func(t *testing.T) {
		errExp := &domain.ErrValidationFailed{Msg: "title is empty"}
		repo := new(mocks.ITaskRepository)
		im := new(mocks.IIDManager)
		im.On("GenerateID").Return(id)
		cm := new(mocks.IClockManager)
		cm.On("GetNow").Return(now)
		srv := service.NewTaskService(repo, im, cm)
		_, err := srv.CreateTask(ctx, uid, "", description, status, dueDate)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: クエリエラーの場合", func(t *testing.T) {
		errExp := &domain.ErrQueryFailed{}
		repo := new(mocks.ITaskRepository)
		repo.On("CreateTask", ctx, task).Return("", errExp)
		im := new(mocks.IIDManager)
		im.On("GenerateID").Return(id)
		cm := new(mocks.IClockManager)
		cm.On("GetNow").Return(now)
		srv := service.NewTaskService(repo, im, cm)
		_, err := srv.CreateTask(ctx, uid, afTitle, description, status, dueDate)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
}

func TestTaskService_UpdateTaskDetails(tt *testing.T) {
	ctx := context.Background()
	id := "id"
	uid := "uid"
	now := time.Now().UTC()
	task := &entity.Task{
		ID:          value.NewID(id),
		UserID:      value.NewID(uid),
		Title:       "task",
		Description: ptrString("description"),
		Status:      entity.StatusPending,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	upd := now.Add(time.Second)
	newDescription := ptrString("new description")
	newDueDate := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)

	tt.Run("正常系: 正しい入力の場合", func(t *testing.T) {
		arg := &entity.Task{
			ID:          task.ID,
			UserID:      task.UserID,
			Title:       task.Title,
			Description: newDescription,
			Status:      entity.StatusInProgress,
			DueDate:     &newDueDate,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   upd,
		}
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(task, nil)
		repo.On("UpdateTask", ctx, arg).Return(nil)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		cm.On("GetNow").Return(upd)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.UpdateTaskDetails(ctx, arg.ID.Value(), arg.UserID.Value(), arg.Title, arg.Description, string(arg.Status), arg.DueDate)

		require.NoError(t, err, "エラーが発生しないこと")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})

	tt.Run("準正常系: 不正な入力(Titleが空)の場合", func(t *testing.T) {
		errExp := &domain.ErrValidationFailed{Msg: "title is empty"}
		arg := &entity.Task{
			ID:          task.ID,
			UserID:      task.UserID,
			Title:       "",
			Description: newDescription,
			Status:      entity.StatusInProgress,
			DueDate:     &newDueDate,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   upd,
		}
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(task, nil)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		cm.On("GetNow").Return(upd)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.UpdateTaskDetails(ctx, arg.ID.Value(), arg.UserID.Value(), arg.Title, arg.Description, string(arg.Status), arg.DueDate)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})

	tt.Run("準正常系: 不正な入力(Statusが不正)の場合", func(t *testing.T) {
		errExp := &domain.ErrValidationFailed{Msg: "invalid status value"}
		arg := &entity.Task{
			ID:          task.ID,
			UserID:      task.UserID,
			Title:       "new task title",
			Description: newDescription,
			Status:      entity.Status("invalid-status"),
			DueDate:     &newDueDate,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   upd,
		}
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(task, nil)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		cm.On("GetNow").Return(upd)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.UpdateTaskDetails(ctx, arg.ID.Value(), arg.UserID.Value(), arg.Title, arg.Description, string(arg.Status), arg.DueDate)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})

	tt.Run("準正常系: 存在しないTaskIDの場合", func(t *testing.T) {
		errExp := &domain.ErrNotFound{Msg: "task not found"}
		arg := &entity.Task{
			ID:          value.NewID("another"),
			UserID:      task.UserID,
			Title:       "new task title",
			Description: newDescription,
			Status:      entity.StatusInProgress,
			DueDate:     &newDueDate,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   upd,
		}
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, arg.ID.Value()).Return(nil, errExp)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.UpdateTaskDetails(ctx, arg.ID.Value(), arg.UserID.Value(), arg.Title, arg.Description, string(arg.Status), arg.DueDate)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})

	tt.Run("準正常系: アクセス権がない場合", func(t *testing.T) {
		errExp := &domain.ErrPermissionDenied{}
		arg := &entity.Task{
			ID:          task.ID,
			UserID:      value.NewID("another"),
			Title:       "new task title",
			Description: newDescription,
			Status:      entity.StatusInProgress,
			DueDate:     &newDueDate,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   upd,
		}
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(task, nil)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.UpdateTaskDetails(ctx, arg.ID.Value(), arg.UserID.Value(), arg.Title, arg.Description, string(arg.Status), arg.DueDate)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: クエリエラーの場合", func(t *testing.T) {
		errExp := &domain.ErrQueryFailed{}
		arg := &entity.Task{
			ID:          task.ID,
			UserID:      task.UserID,
			Title:       "new task title",
			Description: newDescription,
			Status:      entity.StatusInProgress,
			DueDate:     &newDueDate,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   upd,
		}
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(task, nil)
		repo.On("UpdateTask", ctx, arg).Return(errExp)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		cm.On("GetNow").Return(upd)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.UpdateTaskDetails(ctx, arg.ID.Value(), arg.UserID.Value(), arg.Title, arg.Description, string(arg.Status), arg.DueDate)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
}

func TestTaskService_UpdateTaskStatus(tt *testing.T) {
	ctx := context.Background()
	id := "id"
	uid := "uid"
	now := time.Now().UTC()
	task := &entity.Task{
		ID:          value.NewID(id),
		UserID:      value.NewID(uid),
		Title:       "task title",
		Description: ptrString("description"),
		Status:      entity.StatusPending,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	upd := now.Add(time.Second)

	tt.Run("正常系: 正しい入力の場合", func(t *testing.T) {
		arg := &entity.Task{
			ID:          task.ID,
			UserID:      task.UserID,
			Title:       task.Title,
			Description: task.Description,
			Status:      entity.StatusInProgress,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   upd,
		}
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(task, nil)
		repo.On("UpdateTask", ctx, arg).Return(nil)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		cm.On("GetNow").Return(upd)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.UpdateTaskStatus(ctx, id, uid, string(entity.StatusInProgress))

		require.NoError(t, err, "エラーが発生しないこと")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: 不正な入力の場合", func(t *testing.T) {
		errExp := &domain.ErrValidationFailed{Msg: "invalid status value"}
		repo := new(mocks.ITaskRepository)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		cm.On("GetNow").Return(upd)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.UpdateTaskStatus(ctx, id, uid, "invalid-status")

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: 存在しないTaskIDの場合", func(t *testing.T) {
		errExp := &domain.ErrNotFound{Msg: "task not found"}
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(nil, errExp)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.UpdateTaskStatus(ctx, id, uid, string(entity.StatusInProgress))

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: アクセス権がない場合", func(t *testing.T) {
		errExp := &domain.ErrPermissionDenied{}
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(task, nil)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.UpdateTaskStatus(ctx, id, "another-uid", string(entity.StatusInProgress))

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: クエリエラーの場合", func(t *testing.T) {
		errExp := &domain.ErrQueryFailed{}
		arg := &entity.Task{
			ID:          task.ID,
			UserID:      task.UserID,
			Title:       task.Title,
			Description: task.Description,
			Status:      entity.StatusInProgress,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   upd,
		}
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(task, nil)
		repo.On("UpdateTask", ctx, arg).Return(errExp)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		cm.On("GetNow").Return(upd)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.UpdateTaskStatus(ctx, id, uid, string(entity.StatusInProgress))

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
}

func TestTaskService_ChangeTaskDueDate(tt *testing.T) {
	ctx := context.Background()
	id := "id"
	uid := "uid"
	now := time.Now().UTC()
	task := &entity.Task{
		ID:          value.NewID(id),
		UserID:      value.NewID(uid),
		Title:       "task title",
		Description: ptrString("description"),
		Status:      entity.StatusPending,
		DueDate:     nil,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	upd := now.Add(time.Second)
	newDueDate := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)

	tt.Run("正常系: 正しい入力の場合", func(t *testing.T) {
		arg := &entity.Task{
			ID:          task.ID,
			UserID:      task.UserID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			DueDate:     &newDueDate,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   upd,
		}
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(task, nil)
		repo.On("UpdateTask", ctx, arg).Return(nil)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		cm.On("GetNow").Return(upd)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.UpdateTaskDueDate(ctx, id, uid, &newDueDate)

		require.NoError(t, err, "エラーが発生しないこと")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: 存在しないTaskIDの場合", func(t *testing.T) {
		errExp := &domain.ErrNotFound{Msg: "task not found"}
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(nil, errExp)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.UpdateTaskDueDate(ctx, id, uid, &newDueDate)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: アクセス権がない場合", func(t *testing.T) {
		errExp := &domain.ErrPermissionDenied{}
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(task, nil)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.UpdateTaskDueDate(ctx, id, "another-uid", &newDueDate)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: クエリエラーの場合", func(t *testing.T) {
		errExp := &domain.ErrQueryFailed{}
		arg := &entity.Task{
			ID:          task.ID,
			UserID:      task.UserID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			DueDate:     &newDueDate,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   upd,
		}
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(task, nil)
		repo.On("UpdateTask", ctx, arg).Return(errExp)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		cm.On("GetNow").Return(upd)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.UpdateTaskDueDate(ctx, id, uid, &newDueDate)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
}

func TestTaskService_DeleteTask(tt *testing.T) {
	ctx := context.Background()
	id := "id"
	uid := "uid"
	now := time.Now().UTC()
	task := &entity.Task{
		ID:          value.NewID(id),
		UserID:      value.NewID(uid),
		Title:       "task title",
		Description: ptrString("description"),
		Status:      entity.StatusPending,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tt.Run("正常系: 正しい入力の場合", func(t *testing.T) {
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(task, nil)
		repo.On("DeleteTask", ctx, id).Return(nil)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.DeleteTask(ctx, id, uid)

		require.NoError(t, err, "エラーが発生しないこと")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: 不正な入力の場合", func(t *testing.T) {
		errExp := &domain.ErrValidationFailed{Msg: "id is empty"}
		id := ""
		repo := new(mocks.ITaskRepository)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.DeleteTask(ctx, id, uid)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: 存在しないTaskIDの場合", func(t *testing.T) {
		errExp := &domain.ErrNotFound{Msg: "task not found"}
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(nil, errExp)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.DeleteTask(ctx, id, uid)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: アクセス権がない場合", func(t *testing.T) {
		errExp := &domain.ErrPermissionDenied{}
		uid := "another"
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(task, nil)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.DeleteTask(ctx, id, uid)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
	tt.Run("準正常系: クエリエラーの場合", func(t *testing.T) {
		errExp := &domain.ErrQueryFailed{}
		repo := new(mocks.ITaskRepository)
		repo.On("FindTaskByID", ctx, id).Return(task, nil)
		repo.On("DeleteTask", ctx, id).Return(errExp)
		im := new(mocks.IIDManager)
		cm := new(mocks.IClockManager)
		srv := service.NewTaskService(repo, im, cm)
		err := srv.DeleteTask(ctx, id, uid)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
		im.AssertExpectations(t)
		cm.AssertExpectations(t)
	})
}

func ptrString(s string) *string {
	return &s
}
