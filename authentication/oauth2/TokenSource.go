package oauth2

import (
	"golang.org/x/oauth2"
)

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token {
		AccessToken: t.AccessToken,
	}
	
	return token, nil
}
