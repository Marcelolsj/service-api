package http_errors

import (
	"errors"
	"service-api/domain/bu_errors"

	"github.com/gofiber/fiber/v2"
)

type HttpError struct {
	Msg    string
	Status int
}

func (this *HttpError) Error() string {
	return this.Msg
}

func NewHttpErrorMsg(msg string, status int) HttpError {
	return HttpError{
		Msg:    msg,
		Status: status,
	}
}

func NewHttpErrorStatus(err error, status int) HttpError {
	return HttpError{
		Msg:    err.Error(),
		Status: status,
	}
}

func NewHttpError(err error) *HttpError {
	var buErr *bu_errors.BusinessError
	if errors.As(err, &buErr) {
		return &HttpError{
			Msg:    err.Error(),
			Status: fiber.StatusBadRequest,
		}
	}

	return &HttpError{
		Msg:    err.Error(),
		Status: fiber.StatusInternalServerError,
	}
}
