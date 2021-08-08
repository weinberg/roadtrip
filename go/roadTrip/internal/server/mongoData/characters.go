package mongoData

import (
  . "github.com/brickshot/roadtrip/internal/server/types"
)

func (d Provider) GetCharacter(UUID string) (Character, error) {
  return Character{}, nil
}

func (d Provider) StoreCharacter(c Character) error {
  return nil
}

// SetCar sets the characters car. Accepts "" as no car.
func (d Provider) SetCar(charUUID string, carUUID string) (Character, error) {
  return Character{}, nil
}