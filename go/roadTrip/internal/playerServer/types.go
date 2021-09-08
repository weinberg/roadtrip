package playerServer

/**
 * Structs
 */

type Car struct {
  Id          string
  Name        string
  Plate       string
  VelocityMPH float32
  Direction   int32
  Location    *Location
  Trip        *Trip
}

type Character struct {
  Id   string
  Name string
  Car  *Car
}

// A location is either in a town or on a road. If on a road the position is miles from town a->b
type Location struct {
  RoadId        string
  PositionMiles float32
  TownId        string
}

type Trip struct {
  TownIds []string
}

/**
 * Interfaces
 */

type InitConfig struct {
}

type DataProvider interface {
  Init(c InitConfig) DataProvider
  Shutdown()

  CreateCharacter(name string) (Character, error)
  GetCharacter(UUID string) (Character, error)
  DeleteCharacter(UUID string) error

  CreateCar() (Car, error)
  SetCar(charUUID string, carUUID string) (Character, error)
  GetCar(UUID string) (Car, error)
  GetCarByPlate(plate string) (Car, error)
  GetCarByCharacter(UUID string) (Car, error)
  GetTripByCar(UUID string) (Trip, error)
}
