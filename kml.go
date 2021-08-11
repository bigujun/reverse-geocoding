package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"strconv"
)

type kmlFile struct {
	XMLName  xml.Name    `xml:"kml"`
	Document kmlDocument `xml:"Document"`
	Schema   kmlSchema   `xml:"Schema"`
}

type kmlSchema struct {
	XMLName  xml.Name `xml:"Schema"`
	InnerXml string   `xml:",innerxml"`
}

type kmlDocument struct {
	XMLName xml.Name    `xml:"Document"`
	Name    string      `xml:"name"`
	Folders []kmlFolder `xml:"Folder"`
}

type kmlFolder struct {
	XMLName    xml.Name       `xml:"Folder"`
	Name       string         `xml:"name"`
	Placemarks []kmlPlacemark `xml:"Placemark"`
}

type kmlPlacemark struct {
	XMLName      xml.Name        `xml:"Placemark"`
	Name         string          `xml:"name"`
	ExtendedData kmlExtendedData `xml:"ExtendedData"`
}

type kmlExtendedData struct {
	XMLName    xml.Name      `xml:"ExtendedData"`
	SchemaData kmlSchemaData `xml:"SchemaData"`
	Point      kmlPoint      `xml:"Point"`
}

type kmlSchemaData struct {
	XMLName    xml.Name        `xml:"SchemaData"`
	SimpleData []kmlSimpleData `xml:"SimpleData"`
}

type kmlSimpleData struct {
	XMLName xml.Name `xml:"SimpleData"`
	Key     string   `xml:"name,attr"`
	Value   string   `xml:",chardata"`
}

type kmlPoint struct {
	XMLName xml.Name `xml:"Point"`
}

func openKml(fileName string) (*kmlDocument, error) {
	xmlFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var file kmlFile

	err = xml.Unmarshal(byteValue, &file)

	if err != nil {
		return nil, err
	}
	return &file.Document, nil
}

func (d *kmlExtendedData) get(key string) string {
	for _, sd := range d.SchemaData.SimpleData {
		if sd.Key == key {
			return sd.Value
		}
	}
	return ""
}

func (f *kmlFolder) getPlaces() []placemark {
	places := make([]placemark, len(f.Placemarks))
	for i, p := range f.Placemarks {
		places[i] = makePlace(&p)
	}
	return places
}

func (d *kmlDocument) getPlaces() []placemark {
	places := make([]placemark, 0)
	for _, folder := range d.Folders {
		places = append(places, folder.getPlaces()...)
	}
	return places
}

func makePlace(x *kmlPlacemark) placemark {
	id, _ := strconv.ParseInt(x.ExtendedData.get("ID"), 10, 32)
	lat, _ := strconv.ParseFloat(x.ExtendedData.get("LAT"), 64)
	long, _ := strconv.ParseFloat(x.ExtendedData.get("LONG"), 64)

	return placemark{
		id:    id,
		ptype: x.ExtendedData.get("TIPO"),
		city:  x.ExtendedData.get("NM_MUNICIPIO"),
		state: x.ExtendedData.get("NM_UF"),
		lat:   lat,
		long:  long,
	}
}
