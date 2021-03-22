package routers

import (
	"fmt"

	"github.com/axi93/myWeather/models"
)

func ObtainSports(c models.Coordenadas) {

	fmt.Println(c.Main.Temp)
	if c.Main.Temp >= 15 && c.Main.Temp <= 19.99 {
		fmt.Println("You can do exercise outdoor")
	} else if c.Main.Temp >= 10 && c.Main.Temp <= 14.99 {
		fmt.Println("You can do exercise indoor")
	} else {
		fmt.Println("Stay home better bro")
	}

}
