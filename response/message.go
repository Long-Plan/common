package response

type ResponseMessage map[ResponseCode]string

var responseMessage = ResponseMessage{
	SuccessCode: "Success",

	// HTTP Status Code
	BadRequestCode:          "Bad request",
	UnauthorizedCode:        "Unauthorized",
	ForbiddenCode:           "Forbidden",
	NotFoundCode:            "Not found",
	InternalServerErrorCode: "Internal server error",

	// Handler level
	GetBodyErrorCode:      "Failed to get body",
	ValidateBodyErrorCode: "Failed to validate body",
	PermissionDeniedCode:  "Permission denied",
	ParseParamErrorCode:   "Failed to parse param",

	// Service level
	DatabaseErrorCode:    "Database error",
	EncryptDataErrorCode: "Failed to encrypt data",
	HashErrorCode:        "Failed to hash data",

	UnknownErrorCode: "Unknown error",
}

func (r ResponseMessage) GetMessage(code ResponseCode) string {
	if message, ok := r[code]; ok {
		return message
	}
	return ""
}
