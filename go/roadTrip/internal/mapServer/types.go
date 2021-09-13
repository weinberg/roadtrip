package mapServer

/**
 * Structs
 */

type Town struct {
  Id          string `bson:"id"`
  DisplayName string `bson:"display_name"`
  Description string `bson:"description"`
  StateId     string `bson:"state_id"`
}

type Road struct {
  Id          string `bson:"id"`
  DisplayName string `bson:"display_name"`
  TownAId     string `bson:"town_a_id"`
  TownBId     string `bson:"town_b_id"`
  LengthMiles int32  `bson:"length_miles"`
}

type State struct {
  Id   string `bson:"id"`
  Name string `bson:"name"`
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
