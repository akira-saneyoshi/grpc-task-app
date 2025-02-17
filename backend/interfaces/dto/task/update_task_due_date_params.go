package dto

import (
	"time"

	"github.com/akira-saneyoshi/task-app/interfaces/dto"
)

type UpdateTaskDueDateParams struct {
	id      dto.IDParam
	userID  dto.IDParam
	dueDate *time.Time
}

func NewUpdateTaskDueDateParams(id string, userID string, dueDate *time.Time) *UpdateTaskDueDateParams {
	return &UpdateTaskDueDateParams{
		id:      *dto.NewIDParam(id),
		userID:  *dto.NewIDParam(userID),
		dueDate: dueDate,
	}
}

func (p *UpdateTaskDueDateParams) ID() string {
	return p.id.Value()
}

func (p *UpdateTaskDueDateParams) UserID() string {
	return p.userID.Value()
}

func (p *UpdateTaskDueDateParams) DueDate() *time.Time {
	return p.dueDate
}

func (p *UpdateTaskDueDateParams) Validate() error {
	if err := p.id.Validate(); err != nil {
		return err
	}
	if err := p.userID.Validate(); err != nil {
		return err
	}
	return nil
}
