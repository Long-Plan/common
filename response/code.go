package response

type ResponseCode string

const (
	SuccessCode ResponseCode = "0000"

	// HTTP Status Code
	BadRequestCode          ResponseCode = "4000"
	UnauthorizedCode        ResponseCode = "4010"
	ForbiddenCode           ResponseCode = "4030"
	NotFoundCode            ResponseCode = "4040"
	InternalServerErrorCode ResponseCode = "5000"

	// Handler level
	GetBodyErrorCode      ResponseCode = "1000"
	ValidateBodyErrorCode ResponseCode = "1001"
	PermissionDeniedCode  ResponseCode = "1002"
	ParseParamErrorCode   ResponseCode = "1003"

	// Service level
	DatabaseErrorCode    ResponseCode = "2000"
	EncryptDataErrorCode ResponseCode = "2001"
	HashErrorCode        ResponseCode = "2002"
	DecryptDataErrorCode ResponseCode = "2003"

	UnknownErrorCode ResponseCode = "9999"
)
