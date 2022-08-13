package notification_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	vscontext "github.com/vietnam-immigrations/go-utils/v2/pkg/context"
	"github.com/vietnam-immigrations/go-utils/v2/pkg/notification"
)

func TestCreate(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.Background(), vscontext.KeyStage, "prod")
	err := notification.Create(ctx, notification.Notification{
		ID:         uuid.New().String(),
		CSSClasses: "bg-positive text-white",
		Message:    "Test message",
	})
	assert.NoError(t, err)
}
