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

	//Output: Avdeevka 48° 8′ 43″ N, 37° 44′ 42″ E
	// Kireyevka village 47° 46′ 26″ N, 40° 23′ 38″ E
	// 201.734 km
	fmt.Println("Avdeevka", Avdeevka)
	fmt.Println("Kireyevka village", KireyevkaVillage)
	fmt.Printf("%0.3f km", geo.DistanceHaversine(Avdeevka.Point(), KireyevkaVillage.Point())/1000)
}

func ExampleDecimal() {
	Melitopol := libgeo.Point{
		libgeo.NewDecimalLatitude(46.848889),
		libgeo.NewDecimalLongitude(35.3675),
	}

	Kerch := libgeo.Point{
		libgeo.NewDecimalLatitude(45.361944),
		libgeo.NewDecimalLongitude(36.471111),
	}

	// Output: Melitopol 46.848889, 35.3675
	// Kerch 45.361944, 36.471111
	// 186.151 km
	fmt.Println("Melitopol", Melitopol)
	fmt.Println("Kerch", Kerch)
	fmt.Printf("%0.3f km", geo.DistanceHaversine(Melitopol.Point(), Kerch.Point())/1000)
}
