package space

type Planet string

var planetEarthYearToSecondsInPlanetMap = map[string]float64{
	"Mercury": 7600543.81992,
	"Venus":   19414149.052176,
	"Earth":   31557600,
	"Mars":    59354032.69008,
	"Jupiter": 374355659.124,
	"Saturn":  929292362.8848,
	"Uranus":  2651370019.3296,
	"Neptune": 5200418560.032,
}

func Age(seconds float64, planet Planet) float64 {
	yearInPlanetSeconds := planetEarthYearToSecondsInPlanetMap[string(planet)]
	if yearInPlanetSeconds == 0 {
		return -1
	}
	return seconds / yearInPlanetSeconds
}
