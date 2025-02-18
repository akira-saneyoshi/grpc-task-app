package dto

import (
	"time"
)

type UpdateTaskDueDateParams struct {
	id      IDParam
	userID  IDParam
	dueDate *time.Time
}

func NewUpdateTaskDueDateParams(id string, userID string, dueDate *time.Time) *UpdateTaskDueDateParams {
	return &UpdateTaskDueDateParams{
		id:      *NewIDParam(id),
		userID:  *NewIDParam(userID),
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
