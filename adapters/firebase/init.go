package firebase

import (
	"context"

	firebase "firebase.google.com/go"
)

// Init create firebase app
func Init() (*firebase.App, error) {
	return firebase.NewApp(context.Background(), nil)
}
