package notification_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vietnam-immigrations/go-utils/pkg/logger"
	"github.com/vietnam-immigrations/go-utils/pkg/notification"
)

func TestCreate(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	err := notification.Create(context.TODO(), logger.New(), "prod", notification.Notification{
		ID:         uuid.New().String(),
		CSSClasses: "bg-positive text-white",
		Message:    "Test message",
	})
	assert.NoError(t, err)
}
