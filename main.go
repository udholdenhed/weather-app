package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	owm "github.com/briandowns/openweathermap"
	"github.com/fatih/color"
)

const (
	Unit   = "C"
	Lang   = "EN"
	ApiKey = "dee433f83a76207bc072eab389c056bf"
)

const FgColor = color.FgCyan

func main() {
	w, err := owm.NewCurrent(Unit, Lang, ApiKey)
	if err != nil {
		fmt.Println(err.Error())
	}

	for true {
		err = w.CurrentByName(GetLocation())
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		PrintCurrentWeatherData(w)
	}
}

func GetLocation() string {
	_, err := color.New(FgColor).Print("Enter location: ")
	if err != nil {
		fmt.Println(err.Error())
	}

	var location string
	stdin := bufio.NewReader(os.Stdin)
	_, err = fmt.Fscanf(stdin, "%s", &location)
	if err != nil {
		fmt.Println(err.Error())
	}

	return location
}

func PrintCurrentWeatherData(data *owm.CurrentWeatherData) {
	formattedWeatherData := fmt.Sprintf(
		"Main :\n"+
			"\tCoordinates :\t %f, %f.\n"+
			"\tDescription :\t %s.\n"+
			"\tTemp/feels like: %.2f/%.2f.\n"+
			"\tTemp max/min:\t %.2f/%.2f.\n"+
			"\tPressure:\t %d.\n"+
			"\tHumidity:\t %d%%.\n",
		data.GeoPos.Latitude, data.GeoPos.Longitude,
		data.Weather[0].Description,
		data.Main.Temp, data.Main.FeelsLike,
		data.Main.TempMax, data.Main.TempMin,
		int(data.Main.Pressure),
		data.Main.Humidity,
	)

	formattedWeatherData += fmt.Sprintf(
		"Wind :\n"+
			"\tDeg:\t\t %.2f°.\n"+
			"\tSpeed:\t\t %.2fm/s.\n",
		data.Wind.Deg,
		data.Wind.Speed,
	)

	_, err := color.New(FgColor).Println(formattedWeatherData)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
