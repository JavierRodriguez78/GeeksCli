package digitalocean

import (
	"context"
	"net/http"
	"fmt"
	"github.com/digitalocean/godo"
)

type ManagamentDroplet struct {
	OauthClient *http.Client
}

func  (m *ManagamentDroplet) CheckAuthentication()  {
	client := godo.NewClient(m.OauthClient)
	account, res, err := client.Account.Get(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(account.Email)
	fmt.Println(res.StatusCode)
}

func (m *ManagamentDroplet)  CreateDroplet() {
	client := godo.NewClient(m.OauthClient)

	dropletName := "super-cool-droplet"
	createRequest := &godo.DropletCreateRequest {
		Name:   dropletName,
		Region: "nyc3",
		Size:   "s-1vcpu-1gb",
		Image: godo.DropletCreateImage {
			Slug: "ubuntu-14-04-x64",
		},
	}

	ctx := context.TODO()
	newDroplet, _, err := client.Droplets.Create(ctx, createRequest)
	if err != nil {
		fmt.Printf("Something bad happened: %s\n\n", err)

	}

	fmt.Printf(newDroplet.PublicIPv4())
}
