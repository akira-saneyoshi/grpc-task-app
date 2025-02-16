package entity

import (
	"errors"
	"testing"

	"github.com/akira-saneyoshi/task-app/domain"
	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/domain/object/value"
	task "github.com/akira-saneyoshi/task-app/domain/object/value/task"
	"github.com/stretchr/testify/require"
)

func TestTaskEntity_Validate(tt *testing.T) {
	testcases := []struct {
		title string
		arg   *entity.Task
		err   error
	}{
		{"正常系: 正しい入力の場合", &entity.Task{ID: value.NewID("id"), UserID: value.NewID("uid"), Title: task.NewTitle("task")}, nil},
		{"準正常系: IDが空の場合", &entity.Task{ID: value.NewID(""), UserID: value.NewID("uid"), Title: task.NewTitle("task")}, &domain.ErrValidationFailed{Msg: "id is empty"}},
		{"準正常系: UserIDが空の場合", &entity.Task{ID: value.NewID("id"), UserID: value.NewID(""), Title: task.NewTitle("task")}, &domain.ErrValidationFailed{Msg: "user-id is empty"}},
		{"準正常系: titleが空の場合", &entity.Task{ID: value.NewID("id"), UserID: value.NewID("uid"), Title: task.NewTitle("")}, errors.New("title is empty")},
	}
	for _, v := range testcases {
		tt.Run(v.title, func(t *testing.T) {
			err := v.arg.Validate()

			if v.err == nil {
				require.NoError(t, err, "エラーが発生しないこと")
			} else {
				require.EqualError(t, err, v.err.Error(), "エラーが一致すること")
			}
		})
	}
}
