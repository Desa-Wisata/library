package logger

import (
	"context"
	"fmt"
)

// Log ...
func Log(ctx context.Context, message ...interface{}) context.Context {
	v, ok := ctx.Value(LogKey).(*Data)
	if ok {
		v.Messages = append(v.Messages, fmt.Sprint(message...))

		ctx = context.WithValue(ctx, LogKey, v)

		return ctx
	}
	return ctx
}

// Logf ...
func Logf(ctx context.Context, message string, value ...interface{}) context.Context {
	v, ok := ctx.Value(LogKey).(*Data)
	if ok {
		msg := fmt.Sprintf(message, value...)
		v.Messages = append(v.Messages, msg)

		ctx = context.WithValue(ctx, LogKey, v)

		return ctx
	}
	return ctx
}
