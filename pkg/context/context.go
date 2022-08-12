package context

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"github.com/vietnam-immigrations/go-utils/v2/pkg/mongodb"
)

// WithRequest adds request data to context
func WithRequest(ctx context.Context, request *events.APIGatewayProxyRequest) context.Context {
	stage := request.RequestContext.Stage
	result := context.WithValue(ctx, KeyStage, stage)
	correlationID, ok := request.Headers["X-Correlation-Id"]
	if !ok {
		correlationID = uuid.New().String()
	}
	result = context.WithValue(result, KeyCorrelationID, correlationID)
	return result
}

// WithOrder adds order data to context
func WithOrder(ctx context.Context, order mongodb.Order) context.Context {
	result := context.WithValue(ctx, KeyOrderID, order.ID)
	result = context.WithValue(result, KeyOrderWooID, order.OrderID)
	result = context.WithValue(result, KeyOrderNumber, order.Number)
	return result
}
