package main

import (
	L "GeeksCli/lib"
	"GeeksCli/services/digitalocean"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
)


type Config struct {
	Env                 string `yaml:"env"`
}



func main() {
  L.Info()
  digital := digitalocean.ManagamentDroplet{}
  digital.CreateDroplet()

  os.Mkdir("."+string(filepath.Separator) + "geeks",0777)
	f,  err := os.Create("geeks/config.yml")
	f.Name()
	if err != nil {
		fmt.Println(err)
		return
	}
	filename, _ := filepath.Abs("geeks/config.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}


}

