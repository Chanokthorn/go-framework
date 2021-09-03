package std

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"reflect-test/v3/internal/lib/user"
	"regexp"
)

const (
	profileKey = "profile"
)

type JWTAccessToken struct {
	UserWithPermission user.UserPermission `json:"userWithPermission"`
	jwt.StandardClaims
}

// Profile contains all request-related user information
type Profile struct {
	user.UserPermission
}

// ProfileProvider is echo.MiddlewareFunc that decodes Profile from JWT, the process does not validate the JWT
// for example:
//	e := echo.New()
//	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
//		SigningKey: []byte(config.App.KeyEncryption),
//	}))
//	api.Use(std.ProfileProvider)
func ProfileProvider(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		m1 := regexp.MustCompile(`^Bearer `)
		tokenString := m1.ReplaceAllString(authHeader, "")

		token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf(`invalid token: %v`, err)})
		}

		var accessToken JWTAccessToken

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			err = mapstructure.Decode(claims, &accessToken)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf(`invalid token structure: %v`, err)})
			}
		} else {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf(`unable to map token to map claims: %v`, err)})
		}

		ctx := c.Request().Context()

		ctx = context.WithValue(ctx, profileKey, Profile{accessToken.UserWithPermission})

		r := c.Request()
		*r = *r.WithContext(ctx)

		return next(c)
	}
}

// WithProfile arbitrarily sets profile into context. This is only for testing purpose.
func WithProfile(ctx context.Context, profile Profile, next func(ctx context.Context) error) error {
	ctx = context.WithValue(ctx, profileKey, profile)
	return next(ctx)
}

// UseProfile is a hook that retrieves Profile from ctx, it is required that ProfileProvider is set as middleware
// and the ctx is originally generated from WithContext closure or UseContext
// for example:
//	func (p *productRepository) CreateTx(ctx context.Context, product product.ProductRelational) (id int, err error) {
//		profile, err := std.UseProfile(ctx)
//		if err != nil {
//			return 0, err
//		}
//
//		// do something with product and profile...
//	}
func UseProfile(ctx context.Context) (Profile, error) {
	profile, ok := ctx.Value(profileKey).(Profile)
	if !ok {
		return Profile{}, fmt.Errorf(`unable to retrieve profile from context`)
	}

	return profile, nil
}
