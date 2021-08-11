package main

import (
	"fmt"
	"log"
)

func main() {
	document, err := openKml("data/BR_Localidades_2010_v1.kml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(document.Name)

	for _, folder := range document.Folders {
		fmt.Println(" |__" + folder.Name)
	}
	fmt.Println()

	places := document.getPlaces()
	fmt.Printf("Places: %d \n", len(places))

	nearest := nearestPlace(places, -3.1066748, -60.00793)
	fmt.Println("City: " + nearest.city)
	fmt.Println("State: " + nearest.state)
	fmt.Println("Type: " + nearest.ptype)
	//fmt.Printf("Dist %f", nearest.distance(-31.3190535, -52.4629328))
}
