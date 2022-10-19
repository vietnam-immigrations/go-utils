package newrelic_test

import (
	"context"
	"testing"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/stretchr/testify/assert"
	mycontext "github.com/vietnam-immigrations/go-utils/v2/pkg/context"
	mynewrelic "github.com/vietnam-immigrations/go-utils/v2/pkg/newrelic"
)

func TestNewApplication(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(
		context.Background(),
		mycontext.KeyStage, "dev",
	)

	app, err := mynewrelic.NewApplication(ctx)
	defer mynewrelic.Shutdown(ctx)
	assert.Nil(t, err)

	tx := app.StartTransaction("Test_Transaction")
	ctxWithNewrelic := newrelic.NewContext(ctx, tx)
	time.Sleep(1 * time.Second)

	txLater := newrelic.FromContext(ctxWithNewrelic)
	segment := txLater.StartSegment("Test_Segment")
	time.Sleep(1 * time.Second)
	segment.End()

	tx.End()
}
