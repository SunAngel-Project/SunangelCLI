package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cloudsftp/Sunangel/args"
	"github.com/cloudsftp/Sunangel/horizon"
	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/persist"
	"github.com/cloudsftp/Sunangel/sunset"
)

const dateLayout = "2006-01-02"
const timeLayout = "15:04:05 MST"

func main() {
	arguments, err := args.ParseSunArguments(os.Args)
	if err != nil {
		args.PrintSunUsage(err)
	}

	var place *location.Location
	switch arguments.Mode {
	case args.Coordinates:
		place = location.NewLocation(arguments.Place.Latitude, arguments.Place.Longitude)
	case args.Name:
		place, err = persist.GetLocation(arguments.Name)
		if err != nil {
			fmt.Printf("\nLocation %s does not exist in the database\n\n", arguments.Name)
			os.Exit(1)
		}

		fmt.Printf("\nSuccessfully loaded location %s\n\n", arguments.Name)
	default:
		fmt.Printf("\nUnexpected internal state. Exiting\n")
		os.Exit(1)
	}

	horizon := horizon.NewHorizon(place, arguments.StartRadius)

	date := time.Now().Add(time.Duration(arguments.DayOffset * 24 * int(time.Hour)))

	sunsetTime := sunset.EstimateSunsetOf(date, horizon)

	fmt.Printf(
		"\nResult:\n  The sun sets at %s on %s\n\n",
		sunsetTime.Format(timeLayout),
		sunsetTime.Format(dateLayout),
	)
}
