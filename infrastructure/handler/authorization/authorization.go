package authorization

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"github.com/MikelSot/metal-bat/infrastructure/postgres/sessiontoken"
	"github.com/MikelSot/metal-bat/model"
	"github.com/MikelSot/metal-bat/model/authorization"
)

type AuthValidator struct {
	param                  model.RemoteConfig
	logger                 model.Logger
	serviceAccessCodeParam string
	tokenParamName         string
	sessionService         sessiontoken.UseCase
}

func NewAuthValidator(param model.RemoteConfig, logger model.Logger, serviceAccessCodeParam string, tokenParamName string, sessionService sessiontoken.UseCase) *AuthValidator {
	return &AuthValidator{param, logger, serviceAccessCodeParam, tokenParamName, sessionService}
}

func (t *AuthValidator) ValidateServiceAccessCode(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAuthorized := false

		etlServiceAccessCode, err := t.param.GetByName(t.serviceAccessCodeParam)
		if err != nil {
			t.logger.Warnf("no se encontro el access-code en el parametro %s", t.serviceAccessCodeParam)
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Errorf("¡Upps! hubo un problema al buscar el access code del servicio"))
		}

		reqAccessCode := c.QueryParam(t.tokenParamName)
		if etlServiceAccessCode == reqAccessCode {
			isAuthorized = true
		}

		if !isAuthorized {
			t.logger.Warnf("el access code enviado no es validos %s", reqAccessCode)
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Errorf("¡Upps! el access code que nos enviaste no es valido"))

		}

		return next(c)
	}
}

// ValidateJWT Middleware para validar los JWT token
func (t *AuthValidator) ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tokenString string
		tokenString, err := getTokenFromAuthorizationHeader(c.Request())
		if err != nil {
			tokenString, err = getTokenFromURLParams(c.Request())
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Se encontró un error al tratar de leer el token")
			}
		}

		verifyFunction := func(token *jwt.Token) (interface{}, error) {
			return authorization.VerifyKey(), nil
		}

		token, err := jwt.ParseWithClaims(tokenString, &authorization.Claim{}, verifyFunction)
		if err != nil {
			status := http.StatusUnauthorized
			var msg string
			switch err.(type) {
			case *jwt.ValidationError:
				vErr := err.(*jwt.ValidationError)

				switch vErr.Errors {
				case jwt.ValidationErrorExpired:
					msg = "Su token ha expirado, por favor vuelva a ingresar"
				default:
					msg = "Error de validación del token"
				}
			default:
				status = http.StatusInternalServerError
				msg = "Error al procesar el token"
			}

			return echo.NewHTTPError(status, msg)
		}
		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token de acceso no válido")
		}

		userID := token.Claims.(*authorization.Claim).UserID
		email := token.Claims.(*authorization.Claim).Email
		sessionID := token.Claims.(*authorization.Claim).SessionID
		userType := token.Claims.(*authorization.Claim).UserType

		if t.sessionService != nil {
			_, err = t.sessionService.GetByIDAndExpiresDate(sessionID)
			if err == sql.ErrNoRows {
				return echo.NewHTTPError(http.StatusUnauthorized, "La sesión ha finalizado")
			}
			if err != nil {
				t.logger.Warnf("no fue posible validar la sesión del usuario: %v", err)
			}
		}

		c.Set("userID", userID)
		c.Set("email", email)
		c.Set("sessionID", sessionID)
		c.Set("userType", userType)

		return next(c)
	}
}

// getTokenFromAuthorizationHeader busca el token del header Authorization
func getTokenFromAuthorizationHeader(r *http.Request) (string, error) {
	ah := r.Header.Get("Authorization")
	if ah == "" {
		return "", errors.New("el encabezado no contiene la autorización")
	}

	// Should be a bearer token
	if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
		return ah[7:], nil
	} else {
		return "", errors.New("el header no contiene la palabra Bearer")
	}
}

// getTokenFromURLParams busca el token de la URL
func getTokenFromURLParams(r *http.Request) (string, error) {
	ah := r.URL.Query().Get("authorization")
	if ah == "" {
		return "", errors.New("la URL no contiene la autorización")
	}

	return ah, nil
}
