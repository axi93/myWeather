package routers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/axi93/myWeather/models"
)

func ConvertJson(body string) models.Coordenadas {
	replace1 := strings.ReplaceAll(body, "test(", "")
	body = strings.ReplaceAll(replace1, ")", "")

	var jsonWeather = []byte(body)

	var coordenates models.Coordenadas
	err := json.Unmarshal(jsonWeather, &coordenates)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Printf("%+v", coordenates)
	return coordenates
}
