package digitalocean

import (
	"os"
	"net/http"
	"GeeksCli/authentication/oauth2"
)

type Authenticate struct {
}

func  (a *Authenticate) NewClient() (*http.Client) {
	oauth2 := oauth2.Authenticate{}
	return oauth2.NewClient(os.Getenv("ACCESS_TOKEN_DG"));
}