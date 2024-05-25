package utils

import "context"

type ContextKey string

const (
	CtxKeyID ContextKey = "id"
)

type MyContext struct {
	context.Context
}

func (myCtx MyContext) GetUserID() int64 {
	return int64(myCtx.Value(CtxKeyID).(int))
}
