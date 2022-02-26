package main

import (
	"fmt"
	"time"

	"github.com/cloudsftp/Sunangel/angle"
	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/sunangel"
	"github.com/cloudsftp/Sunangel/sunset"
)

func main() {
	date := time.Now()
	// loc := *location.NewLocation(48.8187132, 9.5878127) //Gaensberg
	// loc := *location.NewLocation(48.8230357, 9.5823731) // Burg
	loc := *location.NewLocation(48.8292463, 9.5773359) // Hinter Freibad

	estimatedSunsetTime := sunset.EstimateSunsetOf(date, loc)
	azimutAngle := sunangel.AzimutSunAngleAt(estimatedSunsetTime, loc)

	fmt.Println(estimatedSunsetTime)
	fmt.Printf(
		"azimut %f altitute %f\n",
		angle.DegreesFromRadians(azimutAngle),
		angle.DegreesFromRadians(loc.GetHorizonAngleAt(azimutAngle)),
	)
	/*
		fmt.Print(loc.Horizon)
		for i := 0; i < len(loc.Horizon); i++ {
			fmt.Printf("index %4d, angle %f\n", i, angle.DegreesFromRadians(loc.Horizon[i]))
		} */
}
