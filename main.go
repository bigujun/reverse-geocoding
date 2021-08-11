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

	places := &places{
		places: document.getPlaces(),
	}
	fmt.Printf("Places: %d \n", len(places.places))

	CreateServer(":8080", places)
}
