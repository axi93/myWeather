package main

import (
	"fmt"
	"os"

	"github.com/axi93/myWeather/service"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Give me a city")
		return
	}
	coor, err := service.Check()
	if err != nil {
		fmt.Printf("Error:%s", err)
		return
	}
	service.ObtainSports(coor)
}
