package dto

import (
	"time"

	"github.com/akira-saneyoshi/task-app/application"
	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/interfaces/dto"
)

type UpdateTaskDetailsParams struct {
	id          dto.IDParam
	userID      dto.IDParam
	title       string
	description *string
	status      entity.Status
	dueDate     *time.Time
}

func NewUpdateTaskDetailsParams(id string, userID string, title string, description *string, status string, dueDate *time.Time) *UpdateTaskDetailsParams {
	taskStatus := entity.StatusPending
	if status != "" {
		taskStatus = entity.Status(status)
	}
	return &UpdateTaskDetailsParams{
		id:          *dto.NewIDParam(id),
		userID:      *dto.NewIDParam(userID),
		title:       title,
		description: description,
		status:      taskStatus,
		dueDate:     dueDate,
	}
}

func (p *UpdateTaskDetailsParams) ID() string {
	return p.id.Value()
}

func (p *UpdateTaskDetailsParams) UserID() string {
	return p.userID.Value()
}

func (p *UpdateTaskDetailsParams) Title() string {
	return p.title
}

func (p *UpdateTaskDetailsParams) Description() *string {
	return p.description
}

func (p *UpdateTaskDetailsParams) Status() entity.Status {
	return p.status
}

func (p *UpdateTaskDetailsParams) DueDate() *time.Time {
	return p.dueDate
}

func (p *UpdateTaskDetailsParams) Validate() error {
	if err := p.id.Validate(); err != nil {
		return err
	}
	if err := p.userID.Validate(); err != nil {
		return err
	}
	if len([]rune(p.title)) > 100 {
		return &application.ErrInputValidationFailed{Msg: "[ERROR] title must be 100 characters or less"}
	}
	if p.description != nil && len([]rune(*p.description)) > 200 {
		return &application.ErrInputValidationFailed{Msg: "[ERROR] description must be 200 characters or less"}
	}
	if !entity.IsValidStatus(p.status) {
		return &application.ErrInputValidationFailed{Msg: "[ERROR] invalid status value"}
	}
	return nil
}
