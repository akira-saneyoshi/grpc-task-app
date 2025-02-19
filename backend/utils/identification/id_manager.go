package identification

import "github.com/google/uuid"

type IIDManager interface {
	GenerateID() string
}

type UUIDManager struct{}

func NewUUIDManager() *UUIDManager {
	return &UUIDManager{}
}

func (m *UUIDManager) GenerateID() string {
	return uuid.NewString()
}
