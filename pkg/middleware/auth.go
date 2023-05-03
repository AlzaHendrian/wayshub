package middleware

import (
	"net/http"
	"strings"
	dto "wayshub/dto/result"
	jwtToken "wayshub/pkg/jwt"

	"github.com/labstack/echo/v4"
)

// Declare Result struct here ...
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Create Auth function here ...
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization") //mengambil nilai dari header Authorization dari permintaan HTTP masuk menggunakan fungsi c.Request() dari objek echo.Context.

		if token == "" {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Code: http.StatusBadRequest, Message: "unauthorized"})
		}

		token = strings.Split(token, " ")[1] //Baris ini membagi string token menggunakan spasi sebagai delimiter dan mendapatkan bagian kedua dari string, yaitu token yang sebenarnya.
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, Result{Code: http.StatusUnauthorized, Message: "unathorized"})
		}

		c.Set("userLogin", claims)
		return next(c)
	}
}
