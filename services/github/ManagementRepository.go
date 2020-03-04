package github

import (
	"context"
	"net/http"
	"fmt"
	"github.com/google/go-github/v29/github"
)

type ManagamentRepository struct {
	OauthClient *http.Client
}

func  (m *ManagamentRepository) CheckAuthentication()  {
	client := github.NewClient(m.OauthClient)
	repos, res, err := client.Repositories.List(context.Background(), "", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(repos)
	fmt.Println(res.StatusCode)
}