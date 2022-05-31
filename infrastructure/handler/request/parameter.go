package request

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

// ExtractIDFromURLParam extracts id of the c.Param
func ExtractIDFromURLParam(c echo.Context) (int, error) {
	value := c.Param("id")
	if value == "" {
		return 0, nil
	}

	return strconv.Atoi(value)
}

// ExtractNumberFromURLParamByName extracts id of the c.Param
func ExtractNumberFromURLParamByName(name string, c echo.Context) (int, error) {
	value := c.QueryParam(name)
	if value == "" {
		return 0, nil
	}

	return strconv.Atoi(value)
}

// ExtractIDFromFormParamByName extracts id of the c.Param
func ExtractIDFromFormParamByName(name string, c echo.Context) (int, error) {
	paramValue := c.FormValue(name)
	if paramValue == "" {
		return 0, nil
	}

	return strconv.Atoi(paramValue)
}
