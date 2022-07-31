package logger

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/vietnam-immigrations/go-utils/pkg/mongodb"
)

type LogField string

const (
	LogFieldStage                  LogField = "stage"
	LogFieldFunction               LogField = "function"
	LogFieldRequestID              LogField = "request_id"
	LogFieldRequestPath            LogField = "request_path"
	LogFieldRequestMethod          LogField = "request_method"
	LogFieldCorrelationID          LogField = "correlation_id"
	LogFieldOrderID                LogField = "order_id"
	LogFieldOrderWooID             LogField = "order_woo_id"
	LogFieldOrderNumber            LogField = "order_number"
	LogFieldCustomerPassportNumber LogField = "customer_passport_number"
)

func New() *logrus.Entry {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	return log.WithFields(nil)
}

func NewFromRequest(request *events.APIGatewayProxyRequest) *logrus.Entry {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	stage := request.RequestContext.Stage
	correlationID, ok := request.Headers["X-Correlation-Id"]
	if !ok {
		correlationID = uuid.New().String()
	}

	return log.WithField(string(LogFieldStage), stage).
		WithField(string(LogFieldRequestID), request.RequestContext.RequestID).
		WithField(string(LogFieldRequestPath), request.Path).
		WithField(string(LogFieldRequestMethod), request.HTTPMethod).
		WithField(string(LogFieldCorrelationID), correlationID)
}

func InstrumentOrderData(log *logrus.Entry, order mongodb.Order) *logrus.Entry {
	return log.WithField(string(LogFieldOrderID), order.ID).
		WithField(string(LogFieldOrderWooID), order.OrderID).
		WithField(string(LogFieldOrderNumber), order.Number)
}
