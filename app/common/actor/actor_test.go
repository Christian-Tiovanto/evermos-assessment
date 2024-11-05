package actor_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mughieams/evermos-assessment/app/common/actor"
)

func TestActor(t *testing.T) {
	t.Run("when actor is exist", func(t *testing.T) {
		ctx := context.Background()
		user := actor.Actor{
			ID:    1,
			Name:  "John Doe",
			Phone: "081234567890",
		}

		ctx = actor.SetToContext(ctx, user)
		a := actor.FromContext(ctx)

		assert.Equal(t, user.ID, a.ID)
		assert.False(t, a.IsEmpty())
		assert.Equal(t, user.Name, a.Name)
	})

	t.Run("when actor is empty", func(t *testing.T) {
		ctx := context.Background()
		a := actor.FromContext(ctx)

		assert.True(t, a.IsEmpty())
	})
}
