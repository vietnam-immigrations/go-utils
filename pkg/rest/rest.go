package rest

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/sirupsen/logrus"
)

func Response(log *logrus.Entry, status int, request *events.APIGatewayProxyRequest, body interface{}) *events.APIGatewayProxyResponse {
	log.Infof("http response: [%d] %v", status, body)
	bodyString, err := json.Marshal(body)
	if err != nil {
		return ResponseError(log, 500, request, err)
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(bodyString),
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "*",
		},
	}
}

func ResponseRaw(log *logrus.Entry, status int, body string) *events.APIGatewayProxyResponse {
	log.Infof("http response: [%d] %v", status, body)
	return &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       body,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "*",
		},
	}
}

type ErrorBody struct {
	Message     string `json:"message"`
	Method      string `json:"method"`
	Path        string `json:"path"`
	RequestID   string `json:"requestId"`
	RequestTime string `json:"requestTime"`
}

func ResponseError(log *logrus.Entry, status int, request *events.APIGatewayProxyRequest, err error) *events.APIGatewayProxyResponse {
	log.Infof("http error response: [%d] %v", status, err)
	body := ErrorBody{
		Message:     err.Error(),
		Method:      request.HTTPMethod,
		Path:        request.Path,
		RequestID:   request.RequestContext.RequestID,
		RequestTime: request.RequestContext.RequestTime,
	}
	bodyString, err := json.Marshal(body)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: status,
			Body:       "{\"message\": \"Could not construct error object.\"}",
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "*",
			},
		}
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(bodyString),
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "*",
		},
	}
}
