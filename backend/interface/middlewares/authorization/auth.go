package authorization

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"net/http"
	"strings"
)

func FirebaseAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			opt := option.WithCredentialsFile("./backend/config/development/development.json")
			app, err := firebase.NewApp(context.Background(), nil, opt)
			if err != nil {
				u := fmt.Sprintf("error invalid credential file: %v\n", err)
				return c.JSON(http.StatusUnauthorized, u)
			}
			client, err := app.Auth(context.Background())
			if err != nil {
				u := fmt.Sprintf("error firebase unauthorized: %v\n", err)
				return c.JSON(http.StatusUnauthorized, u)
			}
			header := c.Request().Header.Get("Authorization")
			replaced := strings.Replace(header, "Bearer ", "", 1)
			if replaced == "" {
				u := fmt.Sprintf("error verifying ID token: %v\n", err)
				return c.JSON(http.StatusUnauthorized, u)
			}
			_, err = client.VerifyIDToken(context.Background(), replaced)
			if err != nil {
				u := fmt.Sprintf("error verifying ID token: %v\n", err)
				return c.JSON(http.StatusUnauthorized, u)
			}
			err = next(c)
			return err
		}
	}
}
