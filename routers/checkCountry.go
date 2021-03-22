package routers

import (
	"fmt"
	"os"
	"strings"

	"github.com/axi93/myWeather/api"
)

func CheckCountry() string {

	country := os.Args[1]
	//fmt.Println(country)
	body := api.ObtainWeather(country)

	if strings.Contains(body, "404") {
		fmt.Printf("City not found, please enter an existing city")

	}
	return body
}
