package main

import "math"

type placemark struct {
	id    int64
	ptype string
	city  string
	state string
	lat   float64
	long  float64
}

func (p *placemark) distanceSquared(lat float64, long float64) float64 {
	deltaLat := lat - p.lat
	deltaLong := long - p.long
	return deltaLat*deltaLat + deltaLong*deltaLong
}
func (p *placemark) distance(lat float64, long float64) float64 {
	return math.Sqrt(p.distanceSquared(lat, long))
}

func nearestPlace(places []placemark, lat float64, long float64) placemark {
	smaller := places[0]
	smallerDist := smaller.distanceSquared(lat, long)
	for _, p := range places {
		dist := p.distanceSquared(lat, long)
		if dist < smallerDist {
			smaller = p
			smallerDist = dist
		}
	}
	return smaller
}
