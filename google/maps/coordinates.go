package maps

import (
	"context"
	"fmt"
	"googlemaps.github.io/maps"
)

func (c *MapClient) FindCoordinates(context context.Context, address string) (lat float64, lng float64, err error) {
	geocode, err := c.client.Geocode(context, &maps.GeocodingRequest{
		Address: address,
	})
	if err != nil {
		return
	}

	if len(geocode) == 0 {
		return -1, -1, fmt.Errorf("maps: no results found for address: %s\n", address)
	}

	lat = geocode[0].Geometry.Location.Lat
	lng = geocode[0].Geometry.Location.Lng
	return

}
