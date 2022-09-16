package authorization

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"github.com/MikelSot/metal-bat/model"
)

type AuthValidator struct {
	logger model.Logger
}

func NewAuthValidator(logger model.Logger) *AuthValidator {
	return &AuthValidator{logger}
}

func (a *AuthValidator) ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tokenString string
		tokenString, err := getTokenString(c)
		if err != nil {
			return err
		}

		verifyFunction := func(token *jwt.Token) (interface{}, error) {
			return model.VerifyKey(), nil
		}

		token, err := jwt.ParseWithClaims(tokenString, &model.Claim{}, verifyFunction)
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

		userID := token.Claims.(*model.Claim).UserID
		email := token.Claims.(*model.Claim).Email
		userType := token.Claims.(*model.Claim).UserType

		c.Set("userID", userID)
		c.Set("email", email)
		c.Set("userType", userType)

		return next(c)
	}
}

func getTokenString(c echo.Context) (string, error) {
	tokenString, err := getHeaderToken(c.Request())
	if err == nil {
		return tokenString, nil
	}

	tokenString, err = getTokenFromURLParams(c.Request())
	if err == nil {
		return tokenString, nil
	}

	return "", echo.NewHTTPError(http.StatusUnauthorized, "Se encontró un error al tratar de leer el token")
}

// getHeaderToken busca el token del header Authorization
func getHeaderToken(r *http.Request) (string, error) {
	ah := r.Header.Get("Authorization")
	if ah == "" {
		return "", errors.New("el encabezado no contiene la autorización")
	}

	if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
		return ah[7:], nil
	}

	return "", errors.New("el header no contiene la palabra Bearer")
}

// getTokenFromURLParams busca el token de la URL
func getTokenFromURLParams(r *http.Request) (string, error) {
	ah := r.URL.Query().Get("authorization")
	if ah == "" {
		return "", errors.New("la URL no contiene la autorización")
	}

	return ah, nil
}

func GenerateToken(u model.User, userType uint, IP string) (string, error) {
	var token string

	claim := model.Claim{
		UserID:   u.ID,
		Email:    u.Email,
		IPClient: IP,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "autoPro",
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	token, err := newToken.SignedString(model.SignKey())
	if err != nil {
		return "", err
	}

	return token, nil
}
