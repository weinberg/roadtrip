package mongoData

import (
  . "github.com/brickshot/roadtrip/internal/server/types"
)

/**
 * Methods
 */

// StoreCar stores a car in the datastore. The UUID will be assigned.
func (d Provider) NewCar(c Car) (Car, error) {
  return Car{}, nil
}

// GetCar returns the car referenced by UUID
func (d Provider) GetCar(UUID string) (Car, error) {
  return Car{}, nil
}

// GetCharacters returns the characters in a car.
func (d Provider) GetCharacters(UUID string) ([]Character, error) {
  return []Character{}, nil
}

