package authctx

import (
	"context"
)

// The key type is unexported to prevent collisions
// with context keys defiend in other packages.
type key int

// ContextUserID is the context key for the request's user id
const ContextUserID key = 0

// NewContext returns a Context that carries the request's user id
func NewContext(ctx context.Context, uid int64) context.Context {
	return context.WithValue(ctx, ContextUserID, uid)
}

// FromContext extracts the request's user id from a Context
func FromContext(ctx context.Context) (int64, bool) {
	uid, ok := ctx.Value(ContextUserID).(int64)
	return uid, ok
}
