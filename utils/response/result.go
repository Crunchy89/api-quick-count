package response

import (
	"errors"
	"net/http"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type Result struct {
	Code         int         `json:"code"`
	Data         interface{} `json:"data,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
}

func NewResultData(data interface{}) *Result {
	return &Result{
		Code:         http.StatusOK,
		Data:         data,
		ErrorMessage: "",
	}
}

func NewResultMessage(message string) *Result {
	return &Result{
		Code:         http.StatusOK,
		Data:         message,
		ErrorMessage: "",
	}
}

func NewResultError(code int, err error) *Result {
	return &Result{
		Code:         code,
		Data:         nil,
		ErrorMessage: err.Error(),
	}
}

func Auto(c echo.Context, data interface{}, err error) error {
	if err != nil {
		println(err.Error())
		return Error(c, err)
	} else if data == nil {
		return Accepted(c)
	} else {
		return ResultWithData(c, data)
	}
}

func Accepted(c echo.Context) error {
	result := &Result{
		Code:         http.StatusAccepted,
		Data:         "Accepted",
		ErrorMessage: "",
	}
	return c.JSON(result.Code, result)
}

func Error(c echo.Context, err error) error {
	var code int = 500
	if strings.Contains(err.Error(), "forbidden") {
		code = 403
	}

	if strings.Contains(err.Error(), "not found") {
		code = 404
	}

	if strings.Contains(err.Error(), "already exist") {
		code = 400
	}
	result := NewResultError(code, err)
	println(result.ErrorMessage)
	return c.JSON(result.Code, result)
}

func From(c echo.Context, err error) error {
	if err == nil {
		return Accepted(c)
	}
	return Error(c, err)
}

func ErrorForbidden(c echo.Context) error {
	result := NewResultError(403, errors.New("forbidden"))
	println(result.ErrorMessage)
	return c.JSON(result.Code, result)
}

func ErrorWithMessage(c echo.Context, message string) error {
	return Error(c, errors.New(message))
}

func ErrorWithMessageStatus(c echo.Context, code int, message string) error {
	result := NewResultError(code, errors.New(message))
	return c.JSON(result.Code, result)
}

func ResultWithData(c echo.Context, data interface{}) error {
	result := NewResultData(data)
	return c.JSON(result.Code, result)
}
