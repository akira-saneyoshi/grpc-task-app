package contextkey

import (
	"context"
)

type IContextWriter interface {
	SetUserID(ctx context.Context, userID string) context.Context
}

type ContextWriter struct{}

func NewContextWriter() *ContextWriter {
	return &ContextWriter{}
}

func (w *ContextWriter) SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, ContextKeyUserID, userID)
}
