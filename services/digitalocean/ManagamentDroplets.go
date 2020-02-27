package digitalocean

import (
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)


type ManagamentDroplet struct{

}

func  (m *ManagamentDroplet) Authenticate()  {
	//DigitalOcean
	tokenSource := &TokenSource{
		AccessToken: pat,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)
	account, res, err := client.Account.Get(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(account.Email)
	fmt.Println(res.StatusCode)
}

func (m *ManagamentDroplet)  CreateDroplet(){

	tokenSource := &TokenSource{
		AccessToken: pat,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	dropletName := "super-cool-droplet"

	createRequest := &godo.DropletCreateRequest{
		Name:   dropletName,
		Region: "nyc3",
		Size:   "s-1vcpu-1gb",
		Image: godo.DropletCreateImage{
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
