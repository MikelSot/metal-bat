package response

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/MikelSot/metal-bat/model"
)

func HTTPErrorHandler(err error, c echo.Context) {
	// custom error
	e, ok := err.(*model.Error)
	if ok {
		err = c.JSON(getResponseError(e))
		return
	}

	if echoError, ok := err.(*echo.HTTPError); ok {
		var msgErr string
		msgErr, ok := echoError.Message.(string)
		if !ok {
			msgErr = "¡Upps! algo inesperado ocurrió"
		}

		err = c.JSON(echoError.Code, model.Response{Message: msgErr})
		return
	}

	// if the handler not returns a "model.Error" then it returns a generic error JSON response
	err = c.JSON(http.StatusInternalServerError, model.MessageResponse{
		Errors: model.Responses{
			{Code: UnexpectedError, Message: "¡Uy! metimos la pata, disculpanos lo solucionaremos"},
		},
	})
}

func getResponseError(err *model.Error) (outputStatus int, outputResponse model.MessageResponse) {
	if !err.HasCode() {
		err.SetCode(UnexpectedError)
	}

	if !err.HasAPIMessage() {
		err.SetErrorAsAPIMessage()
	}

	if err.HasData() {
		outputResponse.Data = err.Data()
	}

	if !err.HasStatusHTTP() {
		err.SetStatusHTTP(http.StatusInternalServerError)
	}

	outputStatus = err.StatusHTTP()
	outputResponse.Errors = model.Responses{model.Response{
		Code:    err.Code(),
		Message: err.APIMessage(),
	}}

	return
}
