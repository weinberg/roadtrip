package memoryData

import (
  "errors"
  "fmt"
  . "github.com/brickshot/roadtrip/internal/server"
  "github.com/google/uuid"
  gonanoid "github.com/matoous/go-nanoid"
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

// StoreCar stores a car in the datastore. The UUID and plate will be assigned.
func (d MemoryProvider) NewCar(c Car, owner Character) (Car, error) {
  carsMutex.Lock()
  defer carsMutex.Unlock()

  var plate string
  for i := 0; i < 20; i++ {
    p1, _ := gonanoid.Generate("abcdefghijklmnopqrstuvwxyz", 3)
    p2, _ := gonanoid.Generate("0123456789", 3)
    plate = p1 + "-" + p2
    _, err := d.GetCarByPlate(plate)
    if err == nil {
      continue
    }
    plate = ""
  }

  if plate == "" {
    return Car{}, errors.New("Could not find un-used plate in 20 attempts.")
  }

  var o Character = characters[owner.UUID]

  newCar := Car{
    UUID: uuid.NewString(),
    Name: c.Name,
    Plate: plate,
    Owner: &o,
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

// GetCarByPlate returns the car referenced by Plate
func (d MemoryProvider) GetCarByPlate(plate string) (Car, error) {
  carsMutex.Lock()
  defer carsMutex.Unlock()

  for _,car := range cars {
    if car.Plate == plate {
      return car, nil
    }
  }
  return Car{}, errors.New("Car not found")
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

