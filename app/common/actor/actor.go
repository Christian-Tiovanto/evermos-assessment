// Package actor provides request actor context functionality.
package actor

import (
	"context"

	"github.com/mughieams/evermos-assessment/app/common/util"
)

type actorCtxKey util.Empty

// Actor represents request actor
type Actor struct {
	ID    int64
	Name  string
	Email string
	Phone string
}

// FromContext to get actor from context
func FromContext(ctx context.Context) Actor {
	if u, ok := ctx.Value(actorCtxKey{}).(Actor); ok {
		return u
	}
	return Actor{}
}

// SetToContext to set actor to context
func SetToContext(ctx context.Context, actor Actor) context.Context {
	return context.WithValue(ctx, actorCtxKey{}, actor)
}

// IsEmpty checks if actor is empty
func (a Actor) IsEmpty() bool {
	return a.ID == 0
}
