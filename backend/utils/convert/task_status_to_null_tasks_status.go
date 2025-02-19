package convert

import (
	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/infrastructure/persistence/model/db"
)

// taskStatusToNullTasksStatus is a function that converts entity.TaskStatus to db.NullTasksStatus.
func TaskStatusToNullTasksStatus(status entity.TaskStatus) db.NullTasksStatus {
	if status != "" {
		return db.NullTasksStatus{TasksStatus: db.TasksStatus(status), Valid: true}
	}
	return db.NullTasksStatus{}
}
