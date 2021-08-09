package memoryData

import (
  "errors"
  "fmt"
  . "github.com/brickshot/roadtrip/internal/server/types"
  "github.com/google/uuid"
  "sync"
)

var carsMutex = &sync.Mutex{}
var cars map[string]Car

/**
 * Methods
 */

// init
func InitCars() {
  cars = map[string]Car{}
}

func ShutdownCars() {
  cars = nil
}

func resetCars() {
  cars = map[string]Car{}
}

// StoreCar stores a car in the datastore. The UUID will be assigned.
func (d MemoryProvider) NewCar(c Car) (Car, error) {
  carsMutex.Lock()
  defer carsMutex.Unlock()

  newCar := Car{
    UUID: uuid.NewString(),
    Name: c.Name,
    Plate: "",
  }

  cars[newCar.UUID] = newCar

  return newCar, nil
}

// GetCar returns the car referenced by UUID
func (d MemoryProvider) GetCar(UUID string) (Car, error) {
  carsMutex.Lock()
  defer carsMutex.Unlock()

  fmt.Printf("GetCar: cars: %v\n", cars)

  if _,ok := cars[UUID]; !ok {
    return Car{}, errors.New("Car not found")
  }

  return cars[UUID], nil
}

// GetCharacters returns the characters in a car.
func (d MemoryProvider) GetCharacters(UUID string) ([]Character, error) {
  _, err := d.GetCar(UUID)
  if err != nil {
    return []Character{}, err
  }

  result := []Character{}

  for _,c := range characters {
    if c.UUID == UUID {
      result = append(result, c)
    }
  }
  return result, nil
}

