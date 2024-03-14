package main

import (
	"errors"
	"fmt"
)

func commandMap(config *configuration) error {
	locationsResponse, err := config.pokeapiClient.ListLocations(config.nextLocationsURL)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationsResponse.Next
	config.prevLocationsURL = locationsResponse.Previous

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(config *configuration) error {
	if config.prevLocationsURL == nil {
		return errors.New("you are on the first page")
	}

	locationsResponse, err := config.pokeapiClient.ListLocations(config.prevLocationsURL)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationsResponse.Next
	config.prevLocationsURL = locationsResponse.Previous

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}
