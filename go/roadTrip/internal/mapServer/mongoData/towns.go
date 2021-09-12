package mongoData

import (
  context "context"
  "errors"
  . "github.com/brickshot/roadtrip/internal/mapServer"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "log"
  "time"
)

var townsColl *mongo.Collection

// init
func InitTowns() {
  townsColl = database.Collection("towns")
}

func ShutdownTowns() {
  townsColl = nil
}

// GetTownByName returns the town by name
func (d MongoProvider) GetTown(id string) (Town, error) {
  filter := bson.D{{"id", id}}
  result := Town{}
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  err := townsColl.FindOne(ctx, filter).Decode(&result)
  if err == mongo.ErrNoDocuments {
    return Town{}, errors.New("Not found")
  } else if err != nil {
    log.Fatal(err)
  }

  return result, nil
}

