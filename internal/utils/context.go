package utils

import "context"

type ContextKey string

const (
	CtxKeyID ContextKey = "id"
)

type MyContext struct {
	context.Context
}

func (ctx MyContext) GetUserID() int64 {
	return int64(ctx.Value(CtxKeyID).(int))
}
