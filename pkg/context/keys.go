package context

const (
	KeyStage                   = "stage"
	KeyFunction                = "function"
	KeyRequestID               = "request_id"
	KeyRequestPath             = "request_path"
	KeyRequestMethod           = "request_method"
	KeyCorrelationID           = "correlation_id"
	KeyOrderID                 = "order_id"
	KeyOrderWooID              = "order_woo_id"
	KeyOrderNumber             = "order_number"
	KeyApplicantPassportNumber = "applicant_passport_number"
)

var Keys = []string{
	KeyStage,
	KeyFunction,
	KeyRequestID,
	KeyRequestPath,
	KeyRequestMethod,
	KeyCorrelationID,
	KeyOrderID,
	KeyOrderWooID,
	KeyOrderNumber,
	KeyApplicantPassportNumber,
}
