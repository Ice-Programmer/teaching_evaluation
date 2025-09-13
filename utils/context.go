package utils

import "context"

func ContextSetKeyValue(ctx context.Context, key, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

func ContextGetKeyValue(ctx context.Context, key interface{}) interface{} {
	return ctx.Value(key)
}
