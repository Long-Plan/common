package response

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Code    ResponseCode `json:"code"`
	Message string       `json:"message"`
	Error   interface{}  `json:"error"`
}

type ErrorData struct {
	Code    ResponseCode `json:"code"`
	Message string       `json:"message"`
	Detail  error        `json:"detail"`
}

func (e ErrorData) Error() string {
	return fmt.Sprintf("Code: %s | Message: %s", e.Code, e.Message)
}

func NewErrorData(code ResponseCode, err error) *ErrorData {
	var errorData ErrorData
	if errors.As(err, &errorData) {
		return &errorData
	}

	return &ErrorData{
		Code:    code,
		Message: responseMessage.GetMessage(code),
		Detail:  err,
	}
}

func WriteError(ctx *fiber.Ctx, err error) error {
	response := ErrorResponse{
		Code:    UnknownErrorCode,
		Message: responseMessage.GetMessage(UnknownErrorCode),
	}

	if errData, ok := err.(*ErrorData); ok {
		response.Code = errData.Code
		response.Message = errData.Message
		if errData.Detail != nil {
			response.Error = errData.Detail.Error()
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
