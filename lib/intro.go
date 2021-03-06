package lib

import (
	"fmt"
	"github.com/matishsiao/goInfo"
)

func Info(){
	asciiGeek :=
		`
  /$$$$$$                      /$$                
 /$$__  $$                    | $$                
| $$  \__/  /$$$$$$   /$$$$$$ | $$   /$$  /$$$$$$$
| $$ /$$$$ /$$__  $$ /$$__  $$| $$  /$$/ /$$_____/
| $$|_  $$| $$$$$$$$| $$$$$$$$| $$$$$$/ |  $$$$$$ 
| $$  \ $$| $$_____/| $$_____/| $$_  $$  \____  $$
|  $$$$$$/|  $$$$$$$|  $$$$$$$| $$ \  $$ /$$$$$$$/
 \______/  \_______/ \_______/|__/  \__/|_______/ 

`

	fmt.Printf(asciiGeek)
	fmt.Println("- Tu generador de código de Geekshubs para tus proyectos v.1.0 - ")
	fmt.Println("-----------------------------------------------------------------")
	gi := goInfo.GetInfo()
	gi.VarDump()
	fmt.Println("-----------------------------------------------------------------")

}