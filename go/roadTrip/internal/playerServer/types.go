package playerServer

/**
 * Structs
 */

type Car struct {
	Id                         string    `bson:"id"`
	Name                       string    `bson:"name"`
	Plate                      string    `bson:"plate"`
	VelocityMph                float32   `bson:"velocity_mph"`
	OdometerMiles              float32   `bson:"odometer_miles"`
	Location                   *Location `bson:"location"`
	LastLocationUpdateTimeUnix int64     `bson:"last_location_update_time_unix"`
	Trip                       *Trip     `bson:"trip"`
	OwnerId                    string    `bson:"owner_id"`
}

type Character struct {
	Id   string `bson:"id"`
	Name string `bson:"name"`
	Car  *Car   `bson:"car"`
}

// A location is either in a town or on a road. If on a road the position is miles from town a->b
type Location struct {
	RoadId        string  `bson:"road_id"`
	PositionMiles float32 `bson:"position_miles"`
	TownId        string  `bson:"town_id"`
}

type TripEntry struct {
	Id          string `bson:"id"`
	Type        string `bson:"type"`
	DisplayName string `bson:"display_name"`
}

type Trip struct {
	Entries []TripEntry `bson:"entries"`
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
}
