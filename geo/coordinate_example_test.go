package geo_test

import (
	"fmt"

	"github.com/paulmach/orb/geo"

	libgeo "github.com/libmonsoon-dev/go-lib/geo"
)

func ExampleDMS() {
	Avdeevka := libgeo.Point{
		libgeo.DMS{48, 8, 43, libgeo.N},
		libgeo.DMS{37, 44, 42, libgeo.E},
	}

	KireyevkaVillage := libgeo.Point{
		libgeo.DMS{47, 46, 26, libgeo.N},
		libgeo.DMS{40, 23, 38, libgeo.E},
	}

	//Output: 201.734 km
	fmt.Printf("%0.3f km", geo.DistanceHaversine(Avdeevka.Point(), KireyevkaVillage.Point())/1000)
}

func ExampleDecimal() {

}
