package middlewares

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/hikayat13/alterra-agcm/day-2/submission/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"strconv"
	"time"
)

type JwtCustomClaims struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(user models.User) (string, error) {

	var claims = &JwtCustomClaims{
		UserId: int(user.ID),
		Name:   user.Name,
		Email:  user.Email,
	}
	exp, err := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME")) // in minute
	if err != nil {
		exp = 60
	}
	claims.ExpiresAt = time.Now().Add(time.Minute * time.Duration(exp)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token.SignedString([]byte(os.Getenv("JWT_SECRET"))))
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ExtractToken(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(*JwtCustomClaims)
		return claims
	}
	return nil
}

func init() {
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "Unauthorized"
}

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(
		middleware.JWTConfig{
			Claims:     &JwtCustomClaims{},
			SigningKey: []byte(os.Getenv("JWT_SECRET")),
		},
	)
}
