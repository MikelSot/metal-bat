package response

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"

	"github.com/labstack/echo/v4"

	"github.com/MikelSot/metal-bat/infrastructure/handler/request"
	"github.com/MikelSot/metal-bat/model"
)

const (
	BindFailed model.StatusCode = "bind_failed"
	// Failure sends the custom error message and API message from the logic
	Failure        model.StatusCode = "failure"
	Ok             model.StatusCode = "ok"
	RecordCreated  model.StatusCode = "record_created"
	RecordUpdated  model.StatusCode = "record_updated"
	RecordDeleted  model.StatusCode = "record_deleted"
	RecordNotFound model.StatusCode = "record_not_found"
	// UnexpectedError is a server error
	UnexpectedError model.StatusCode = "unexpected_error"
	// AuthError is any of authorization errors
	AuthError model.StatusCode = "authorization_error"
)

type Response interface {
	OK(data interface{}) (int, model.MessageResponse)
	Created(data interface{}) (int, model.MessageResponse)
	Updated(data interface{}) (int, model.MessageResponse)
	Deleted(data interface{}) (int, model.MessageResponse)
	NotFound() (int, string)
	BindFailed(c echo.Context, err error) error
	UnexpectedError(c echo.Context, who string, err error) *model.Error
	ErrorHandled(c echo.Context, who string, err *model.Error) *model.Error
	AuthorizationError(c echo.Context, who string, err *model.Error, statusCode int) *model.Error
	Error(c echo.Context, who string, err error) *model.Error
}

type API struct {
	logger model.Logger
}

func New(logger model.Logger) API {
	return API{logger}
}

func (a API) OK(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: Ok, Message: "¡Listo!"}},
	}
}

func (a API) Created(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: RecordCreated, Message: "¡listo!"}},
	}
}

func (a API) Updated(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: RecordUpdated, Message: "¡listo!"}},
	}
}

func (a API) Deleted(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: RecordDeleted, Message: "¡listo!"}},
	}
}

func (a API) NotFound() (int, string) {
	return http.StatusNotFound, ""
}

func (a API) BindFailed(c echo.Context, err error) error {
	fun, _, line, _ := runtime.Caller(1)

	e := model.NewError()
	e.SetError(err)
	e.SetCode(BindFailed)
	e.SetStatusHTTP(http.StatusBadRequest)
	e.SetEndpoint(c.Path())
	e.SetQueryParams(c.QueryParams().Encode())
	e.SetWhere(fmt.Sprintf("%s:%d", runtime.FuncForPC(fun).Name(), line))
	e.SetWho("c.Bind()")

	a.logger.Warnf("%s", e.Error())

	return e
}

func (a API) UnexpectedError(c echo.Context, who string, err error) *model.Error {
	fun, _, line, _ := runtime.Caller(1)

	e := model.NewError()
	e.SetError(err)
	e.SetAPIMessage("¡Uy! metimos la pata, disculpanos lo solucionaremos")
	e.SetCode(UnexpectedError)
	e.SetStatusHTTP(http.StatusInternalServerError)
	e.SetEndpoint(c.Path())
	e.SetQueryParams(c.QueryParams().Encode())
	e.SetWhere(fmt.Sprintf("%s:%d", runtime.FuncForPC(fun).Name(), line))
	e.SetWho(who)
	e.SetUserID(request.GetUserID(c))

	a.logger.Errorf("%s", e.Error())

	return e
}

func (a API) ErrorHandled(c echo.Context, who string, e *model.Error) *model.Error {
	fun, _, line, _ := runtime.Caller(1)

	e.SetCode(Failure)
	e.SetEndpoint(c.Path())
	e.SetQueryParams(c.QueryParams().Encode())
	e.SetWhere(fmt.Sprintf("%s:%d", runtime.FuncForPC(fun).Name(), line))
	e.SetWho(who)
	e.SetUserID(request.GetUserID(c))

	if !e.HasStatusHTTP() {
		e.SetStatusHTTP(http.StatusBadRequest)
	}

	if e.StatusHTTP() < http.StatusInternalServerError {
		a.logger.Warnf("%s", e.Error())
		return e
	}

	a.logger.Errorf("%s", e.Error())

	return e
}

func (a API) AuthorizationError(c echo.Context, who string, e *model.Error, statusCode int) *model.Error {
	fun, _, line, _ := runtime.Caller(1)

	e.SetCode(AuthError)
	e.SetEndpoint(c.Path())
	e.SetQueryParams(c.QueryParams().Encode())
	e.SetWhere(fmt.Sprintf("%s:%d", runtime.FuncForPC(fun).Name(), line))
	e.SetWho(who)
	e.SetStatusHTTP(statusCode)
	e.SetUserID(request.GetUserID(c))

	if e.StatusHTTP() < http.StatusInternalServerError {
		a.logger.Warnf("%s", e.Error())
		return e
	}

	a.logger.Errorf("%s", e.Error())

	return e
}

func (a API) Error(c echo.Context, who string, err error) *model.Error {
	e := model.NewError()
	if errors.As(err, &e) {
		return a.ErrorHandled(c, who, e)
	}

	return a.UnexpectedError(c, who, err)
}
