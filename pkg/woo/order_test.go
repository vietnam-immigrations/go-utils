package woo

import (
	"context"
	"testing"

	vscontext "github.com/vietnam-immigrations/go-utils/v2/pkg/context"
)

func TestGetOrder(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.Background(), vscontext.KeyStage, "dev")
	_, err := GetOrder(ctx, "110")
	if err != nil {
		t.Fail()
	}
}
