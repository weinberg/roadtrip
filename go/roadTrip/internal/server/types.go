package server

/**
 * Structs
 */

type Car struct {
  Id         string
  Name       string
  Plate      string
}

type Character struct {
  Id string
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

  CreateCharacter(name string) (Character, error)
  GetCharacter(UUID string) (Character, error)
  DeleteCharacter(UUID string) error

  CreateCar() (Car, error)
  SetCar(charUUID string, carUUID string) (Character, error)
  GetCar(UUID string) (Car, error)
  GetCarByPlate(plate string) (Car, error)
  GetCarByCharacter(UUID string) (Car, error)
}
