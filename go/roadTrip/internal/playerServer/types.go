package playerServer

/**
 * Structs
 */

type Car struct {
  Id         string
  Name       string
  Plate      string
  Location   Location
}

// A location is either in a town or on a road. If on a road the position is from (0..1) from road.TownA -> road.TownB
type Location struct {
  RoadId string
  Position float32
  TownId string
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

  GetLocation(id string) (Location, error)
}
