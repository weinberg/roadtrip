package mapServer

/**
 * Structs
 */

type Town struct {
  Id          string
  Name        string
  Description string
  State       string
}

type Road struct {
  Id    string
  Name  string
  TownA string
  TownB string
}

type State struct {
  Id   string
  Name string
}

/**
 * Interfaces
 */

type InitConfig struct {
}

type DataProvider interface {
  Init(c InitConfig) DataProvider
  Shutdown()

  GetTown(id string) (Town, error)
  GetRoad(id string) (Road, error)
  GetRoads(id string) ([]Road, error)
  GetStates() ([]State, error)
}
