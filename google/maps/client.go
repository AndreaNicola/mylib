package maps

import (
	"fmt"
	"googlemaps.github.io/maps"
)

type MapClient struct {
	client *maps.Client
}

func NewMapClient(apiKey string) (*MapClient, error) {
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("error while creating google maps client: %v", err)
	}
	return &MapClient{client: client}, nil
}
