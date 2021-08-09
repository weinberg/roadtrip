package mongoData

import (
  . "github.com/brickshot/roadtrip/internal/server/types"
  "go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

// init
func InitCars() {
  collection = client.Database("roadtrip").Collection("cars")
}

func ShutdownCars() {
  collection = nil
}

// StoreCar stores a car in the datastore. The UUID will be assigned.
func (d MongoProvider) NewCar(c Car) (Car, error) {
  return Car{}, nil
}

// GetCar returns the car referenced by UUID
func (d MongoProvider) GetCar(UUID string) (Car, error) {
  return Car{}, nil
}

// GetCharacters returns the characters in a car.
func (d MongoProvider) GetCharacters(UUID string) ([]Character, error) {
  return []Character{}, nil
}

