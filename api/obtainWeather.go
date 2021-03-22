package api

import (
	"io/ioutil"
	"net/http"
)

func ObtainWeather(weather string) string {

	url := "https://community-open-weather-map.p.rapidapi.com/weather?q=" + weather + "&lat=0&lon=0&callback=test&id=2172797&lang=null&units=metric"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "abe8dfb9b6msh537e0118d274247p1b7b4cjsne9ffd12b58eb")
	req.Header.Add("x-rapidapi-host", "community-open-weather-map.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)
	/*
		if res.StatusCode != 200 {
			fmt.Println("TODO OK")

		}
	*/

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)

}
