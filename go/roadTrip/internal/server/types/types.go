package types

/**
 * Structs
 */

type Car struct {
  UUID       string
  Name       string
  Plate      string
}

type Character struct {
  UUID string
  Name string
  Car  *Car
}

/**
 * Interfaces
 */

type InitConfig struct {
}

type DataProvider interface {
  Init(c InitConfig) (DataProvider)
  Shutdown()
  NewCar(c Car) (Car, error)
  GetCar(UUID string) (Car, error)
  GetCharacters(UUID string) ([]Character, error)
  GetCharacter(UUID string) (Character, error)
  StoreCharacter(c Character) error
  SetCar(charUUID string, carUUID string) (Character, error)
}
