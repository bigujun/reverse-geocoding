package main

import "math"

type placemark struct {
	ID    int64   `json:"id"`
	Ptype string  `json:"type"`
	City  string  `json:"city"`
	State string  `json:"state"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

type places struct {
	places []placemark
}

func (p *placemark) distanceSquared(lat float64, long float64) float64 {
	deltaLat := lat - p.Lat
	deltaLong := long - p.Long
	return deltaLat*deltaLat + deltaLong*deltaLong
}
func (p *placemark) distance(lat float64, long float64) float64 {
	return math.Sqrt(p.distanceSquared(lat, long))
}

func (p *places) nearestPlace(lat float64, long float64) placemark {
	smaller := p.places[0]
	smallerDist := smaller.distanceSquared(lat, long)
	for _, p := range p.places {
		dist := p.distanceSquared(lat, long)
		if dist < smallerDist {
			smaller = p
			smallerDist = dist
		}
	}
	return smaller
}
