package contextkey

import (
	"context"
	"errors"
)

type IContextReader interface {
	GetUserID(ctx context.Context) (string, error)
}

type ContextReader struct{}

func NewContextReader() *ContextReader {
	return &ContextReader{}
}

func (r *ContextReader) GetUserID(ctx context.Context) (string, error) {
	if v := ctx.Value(ContextKeyUserID); v != nil {
		if userID, ok := v.(string); ok {
			return userID, nil
		}
		return "", errors.New("[ERROR] context value not of type string for user-id")
	}
	return "", errors.New("[ERROR] context value not found for user-id")
}
