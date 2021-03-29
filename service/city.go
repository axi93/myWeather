package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/axi93/myWeather/api"
	"github.com/axi93/myWeather/models"
)

func Check() (models.Coordenadas, error) {
	city := os.Args[1]
	body := api.ObtainWeather(city)

	if strings.Contains(body, "404") {
		return models.Coordenadas{}, errors.New(" city not found, please enter a valid city ")
	}

	return convertJson(body)
}

func convertJson(body string) (models.Coordenadas, error) {
	replace1 := strings.ReplaceAll(body, "test(", "")
	body = strings.ReplaceAll(replace1, ")", "")
	body = strings.ReplaceAll(body, "[", "")
	body = strings.ReplaceAll(body, "]", "")

	var jsonWeather = []byte(body)

	var coordenates models.Coordenadas
	err := json.Unmarshal(jsonWeather, &coordenates)
	if err != nil {
		return models.Coordenadas{}, err
	}

	return coordenates, nil
}

func ObtainSports(c models.Coordenadas) {
	if c.Main.Temp >= 15 && c.Main.Temp <= 19.99 {
		fmt.Println("You can do exercise outdoor")
	} else if c.Main.Temp >= 10 && c.Main.Temp <= 14.99 {
		fmt.Println("You can do exercise indoor")
	} else {
		fmt.Println("Stay home better bro")
	}
}
