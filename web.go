package main

import (
	"fmt"
	"net/http"
)

type WebCommand string

const (
	Get WebCommand = "get"
)

func handleWeb(command []string) {
	switch WebCommand(command[0]) {
	case Get:
		handleGet(command[1])
	}
}

func handleGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

func loadWeb() {

}

//https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&models=dwd_icon_seamless&current=temperature_2m,is_day,precipitation,wind_direction_10m,wind_speed_10m,apparent_temperature&daily=temperature_2m_max&timezone=Europe%2FBerlin&forecast_days=1
