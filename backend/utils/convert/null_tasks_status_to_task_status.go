package convert

import (
	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/infrastructure/persistence/model/db"
)

// nullTasksStatusToTaskStatus is a function that converts db.NullTasksStatus to entity.TaskStatus.
func NullTasksStatusToTaskStatus(nts db.NullTasksStatus) entity.TaskStatus {
	return entity.TaskStatus(nts.TasksStatus)
}
