package convert

import (
	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	task_v1 "github.com/akira-saneyoshi/task-app/interfaces/proto/task/v1"
)

func ConvertStatus(entityStatus entity.TaskStatus) task_v1.TaskStatus {
	switch entityStatus {
	case entity.StatusPending:
		return task_v1.TaskStatus_TASK_STATUS_PENDING_UNSPECIFIED
	case entity.StatusInProgress:
		return task_v1.TaskStatus_TASK_STATUS_IN_PROGRESS
	case entity.StatusCompleted:
		return task_v1.TaskStatus_TASK_STATUS_COMPLETED
	default:
		return task_v1.TaskStatus_TASK_STATUS_PENDING_UNSPECIFIED
	}
}
