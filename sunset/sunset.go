package sunset

import (
	"log"
	"time"

	"github.com/cloudsftp/Sunangel/horizon"
	"github.com/cloudsftp/Sunangel/sunangel"
)

// EstimateSunsetOf returns an estimate of the time of sunset
// at a given time and place.
func EstimateSunsetOf(date time.Time, horizon *horizon.Horizon) time.Time {
	year := date.Year()
	month := date.Month()
	day := date.Day()
	loc := date.Location()

	lowerBound := time.Date(year, month, day, 12, 0, 0, 0, loc)
	upperBound := time.Date(year, month, day, 23, 59, 59, 1e9-1, loc)

	return binarySunsetSearch(lowerBound, upperBound, horizon)
}

func binarySunsetSearch(lowerBound, upperBound time.Time, horizon *horizon.Horizon) time.Time {
	limitSearchDuration := time.Duration(1e9)

	for {
		currentSearchDuration := upperBound.Sub(lowerBound)
		newBound := lowerBound.Add(currentSearchDuration / 2)

		azimutAngle := sunangel.AzimutSunAngleAt(newBound, horizon.Place)
		horizonAngle := sunangel.AltitudeSunAngleAt(newBound, horizon.Place)

		horizonAngleGoal := horizon.GetAltitude(azimutAngle)

		if horizonAngle < horizonAngleGoal {
			upperBound = newBound
		} else {
			lowerBound = newBound
		}

		log.Printf(
			"Searching sunset, current range %s - %s, precision %v",
			formatBound(lowerBound), formatBound(upperBound), currentSearchDuration,
		)

		if currentSearchDuration < limitSearchDuration {
			break
		}
	}

	return lowerBound.Round(limitSearchDuration)
}

func formatBound(date time.Time) string {
	return date.Format("15:04:05.000 MST")
}
