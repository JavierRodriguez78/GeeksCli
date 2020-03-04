package main

import (
	L "GeeksCli/lib"
	"GeeksCli/services/digitalocean"
	"GeeksCli/services/github"
	"fmt"
	// "gopkg.in/yaml.v3"
	// "io/ioutil"
	// "os"
	// "path/filepath"
)

type Config struct {
	Env string `yaml:"env"`
}

func main() {
	L.Info()
	digitaloceanAuthentication := digitalocean.Authenticate{}
	digitalocean := digitalocean.ManagamentDroplet{
		OauthClient: digitaloceanAuthentication.NewClient(),
	}

	fmt.Println("- digitalocean - ")
	digitalocean.CheckAuthentication()
	// digitalocean.CreateDroplet()
	fmt.Println("-----------------------------------------------------------------")


	githubAuthentication := github.Authenticate{}
	github := github.ManagamentRepository{
		OauthClient: githubAuthentication.NewClient(),
	}
  
	fmt.Println("- github - ")
	github.CheckAuthentication()
	fmt.Println("-----------------------------------------------------------------")

  	// os.Mkdir("."+string(filepath.Separator) + "geeks",0777)
	// f,  err := os.Create("geeks/config.yml")
	// f.Name()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// filename, _ := filepath.Abs("geeks/config.yml")
	// yamlFile, err := ioutil.ReadFile(filename)
	// var config Config
	// err = yaml.Unmarshal(yamlFile, &config)
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }
}

