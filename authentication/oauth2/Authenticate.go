package oauth2

import (
	"context"
	"net/http"
	"golang.org/x/oauth2"
)

type Authenticate struct {
}

func  (o *Authenticate) NewClient(token string) (*http.Client) {
	tokenSource := &TokenSource {
		AccessToken: token,
	}

	return oauth2.NewClient(context.Background(), tokenSource)
}