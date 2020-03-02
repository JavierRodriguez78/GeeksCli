package oauth2

import (
	"os"
	"context"
	"net/http"
	"golang.org/x/oauth2"
)

type Authenticate struct {
}

func  (a *Authenticate) LogIn() (*http.Client) {
	tokenSource := &TokenSource {
		AccessToken: os.Getenv("ACCESS_TOKEN"),
	}

	return oauth2.NewClient(context.Background(), tokenSource)
}
