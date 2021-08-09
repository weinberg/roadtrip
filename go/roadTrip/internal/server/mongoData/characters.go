package mongoData

import (
  "context"
  "errors"
  . "github.com/brickshot/roadtrip/internal/server/types"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "log"
  "time"
)

var coll *mongo.Collection

// init
func InitCharacters() {
  coll = database.Collection("characters")
}

func ShutdownCharacters() {
  collection = nil
}

func (d MongoProvider) GetCharacter(UUID string) (Character, error) {
  filter := bson.D{{ "uuid", UUID }}
  result := Character{}
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  err := coll.FindOne(ctx, filter).Decode(&result)
  if err == mongo.ErrNoDocuments {
    return Character{}, errors.New("Not found")
  } else if err != nil {
    log.Fatal(err)
  }
  return result, nil
}

func (d MongoProvider) StoreCharacter(c Character) error {
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  _, err := coll.InsertOne(ctx, bson.D{{"uuid", c.UUID}, {"name", c.Name}})
  return err
}

// SetCar sets the characters car. Accepts "" as no car.
func (d MongoProvider) SetCar(charUUID string, carUUID string) (Character, error) {
  return Character{}, nil
}