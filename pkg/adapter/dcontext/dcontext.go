package dcontext

import (
	"context"
)

type dkey struct{}

var dk dkey = struct{}{}

// SetUserID ContextへユーザIDを保存する
func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, dk, userID)
}

// GetUserIDFromContext ContextからユーザIDを取得する
func GetUserIDFromContext(ctx context.Context) string {
	var userID string
	if ctx.Value(dk) != nil {
		userID = ctx.Value(dk).(string)
	}
	return userID
}
