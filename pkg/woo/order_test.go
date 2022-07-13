package woo

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetOrder(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	_, err := GetOrder(context.TODO(), logrus.WithFields(nil), "dev", "110")
	if err != nil {
		t.Fail()
	}
}
