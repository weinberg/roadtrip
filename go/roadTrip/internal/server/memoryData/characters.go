package memoryData

import (
  "errors"
  . "github.com/brickshot/roadtrip/internal/server"
  "github.com/google/uuid"
  "sync"
)

var charactersMutex = &sync.Mutex{}
var characters map[string]Character

// init
func InitCharacters() {
  characters = map[string]Character{}
}

func resetCharacters() {
  characters = map[string]Character{}
}

func ShutdownCharacters() {
  characters = nil
}

func (d MemoryProvider) GetCharacter(UUID string) (Character, error) {
  charactersMutex.Lock()
  defer charactersMutex.Unlock()

  if val, ok := characters[UUID]; ok {
    return Character{
      UUID: val.UUID,
      Name: val.Name,
      Car:  val.Car,
    }, nil
  }

  return Character{}, errors.New("Not found")
}

func (d MemoryProvider) NewCharacter() (Character, error) {
  charactersMutex.Lock()
  defer charactersMutex.Unlock()

  uuid := uuid.NewString();
  nc := Character{
    UUID: uuid,
    Name: "",
    Car:  nil,
  }
  characters[uuid] = nc

  return nc, nil
}

func (d MemoryProvider) DeleteCharacter(c Character) error {
  charactersMutex.Lock()
  defer charactersMutex.Unlock()

  _, err := uuid.Parse(c.UUID)
  if err != nil {
    return errors.New("Invalid UUID")
  }

  delete(characters, c.UUID)

  return nil
}

// SetCar sets the characters car. Accepts "" as no car.
func (d MemoryProvider) SetCar(charUUID string, carUUID string) (Character, error) {
  char, err := d.GetCharacter(charUUID)
  if err != nil {
    return Character{}, err
  }

  if carUUID == "" {
    char.Car = nil
    return char, nil
  }

  car, err := d.GetCar(carUUID)
  if err != nil {
    return Character{}, err
  }

  char.Car = &car

  return char, nil
}