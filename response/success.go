package response

import "github.com/gofiber/fiber/v2"

type SuccessResponse struct {
	Code    ResponseCode `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
}

func WriteSuccess(ctx *fiber.Ctx, data interface{}) error {
	response := SuccessResponse{
		Code:    SuccessCode,
		Message: responseMessage[SuccessCode],
		Data:    data,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
