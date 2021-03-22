package main

import (
	"fmt"
	"os"

	"github.com/axi93/myWeather/routers"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Give me a country")
		return
	}
	body := routers.CheckCountry()

	json := routers.ConvertJson(body)

	routers.ObtainSports(json)
}
