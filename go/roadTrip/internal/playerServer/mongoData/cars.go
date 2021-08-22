package mongoData

import (
  context "context"
  "errors"
  . "github.com/brickshot/roadtrip/internal/playerServer"
  "github.com/google/uuid"
  gonanoid "github.com/matoous/go-nanoid"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "time"
)

var carsColl *mongo.Collection

// init
func InitCars() {
  carsColl = database.Collection("cars")
}

func ShutdownCars() {
  carsColl = nil
}

// CreateCar stores a car in the datastore. The id and plate will be assigned.
// Currently all cars start in Seattle.
func (d MongoProvider) CreateCar(c Car, owner Character) (Car, error) {
  // find an unused plate number
  var plate string
  for i := 0; i < 10; i++ {
    p1, _ := gonanoid.Generate("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 3)
    p2, _ := gonanoid.Generate("0123456789", 3)
    plate = p1 + "-" + p2
    _, err := d.GetCarByPlate(plate)
    if err == nil {
      continue
    }
    break
  }
  if plate == "" {
    return Car{}, errors.New("Could not find un-used plate in 10 attempts.")
  }

  var ctx context.Context
  ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
  result, err := carsColl.InsertOne(ctx, bson.D{
    {"id", uuid.NewString()},
    {"plate", plate},
    {"name", c.Name},
    {"owner_id", owner.Id},
    {"location", Location{
      RoadId:   "",
      Position: 0,
      TownId:   "states/Washington/towns/Seattle",
    }},
  })
  if err != nil {
    return Car{}, err
  }

  newCar := Car{}
  ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
  filter := bson.D{{"_id", result.InsertedID}}
  err = carsColl.FindOne(ctx, filter).Decode(&newCar)
  if err != nil {
    return Car{}, err
  }
  return newCar, nil
}

// GetCar returns the car referenced by id
func (d MongoProvider) GetCar(id string) (Car, error) {
  return Car{}, nil
}

// GetCarByPlate returns the car referenced by Plate
func (d MongoProvider) GetCarByPlate(plate string) (Car, error) {
  return Car{}, nil
}

// GetCharacters returns the characters in a car.
func (d MongoProvider) GetCharacters(id string) ([]Character, error) {
  return []Character{}, nil
}
