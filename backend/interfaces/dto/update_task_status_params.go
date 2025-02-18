package dto

import (
	"github.com/akira-saneyoshi/task-app/application"
	"github.com/akira-saneyoshi/task-app/domain/object/entity"
)

type UpdateTaskStatusParams struct {
	id     IDParam
	userID IDParam
	status entity.Status
}

func NewUpdateTaskStatusParams(id string, userID string, status string) *UpdateTaskStatusParams {
	taskStatus := entity.StatusPending
	if status != "" {
		taskStatus = entity.Status(status)
	}
	return &UpdateTaskStatusParams{
		id:     *NewIDParam(id),
		userID: *NewIDParam(userID),
		status: taskStatus,
	}
}

func (p *UpdateTaskStatusParams) ID() string {
	return p.id.Value()
}

func (p *UpdateTaskStatusParams) UserID() string {
	return p.userID.Value()
}

func (p *UpdateTaskStatusParams) Status() entity.Status {
	return p.status
}

func (p *UpdateTaskStatusParams) Validate() error {
	if err := p.id.Validate(); err != nil {
		return err
	}
	if err := p.userID.Validate(); err != nil {
		return err
	}
	if !entity.IsValidStatus(p.status) {
		return &application.ErrInputValidationFailed{Msg: "[ERROR] invalid status value"}
	}
	return nil
}
