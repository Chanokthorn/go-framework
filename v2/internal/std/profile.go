package std

import (
	"context"
	"fmt"
)

const profileKey = "profile"

type UserPermission struct {
	UserUUID  string `json:"userUUID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Profile contains all request-related user information
type Profile struct {
	UserPermission
}

func SetProfile(ctx context.Context, profile Profile) context.Context {
	return context.WithValue(ctx, profileKey, profile)
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
