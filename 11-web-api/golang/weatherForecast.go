package main

type WeatherForecast struct {
	Date         string
	TemperatureC int
	TemperatureF int
	Summary      string
}

func (w *WeatherForecast) setTemperatureF() {
	w.TemperatureF = 32 + int((float32(w.TemperatureC) / 0.5556))
}

func NewWeatherForecast(date string, temperatureC int, summary string) *WeatherForecast {
	w := new(WeatherForecast)
	w.Date = date
	w.TemperatureC = temperatureC
	w.Summary = summary
	w.setTemperatureF()
	return w
}
