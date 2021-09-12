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

var roadsColl *mongo.Collection

// init
func InitRoads() {
  roadsColl = database.Collection("roads")
}

func ShutdownRoads() {
  roadsColl = nil
}

// GetRoads returns the roads connected to a road id
func (d MongoProvider) GetRoads(id string) ([]Road, error) {
  filter := bson.M{"$or": []bson.M{{"a": id}, {"b": id}}}
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  cur, err := roadsColl.Find(ctx, filter)
  if err != nil {
    return []Road{}, err
  }

  var roads []Road
  if err = cur.All(ctx, &roads); err != nil {
    return []Road{}, err
  }

  return roads, nil
}

// GetRoad returns the road
func (d MongoProvider) GetRoad(id string) (Road, error) {
  filter := bson.D{{"id", id}}
  result := Road{}
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  err := roadsColl.FindOne(ctx, filter).Decode(&result)
  if err == mongo.ErrNoDocuments {
    return Road{}, errors.New("Not found")
  } else if err != nil {
    log.Fatal(err)
  }

  return result, nil
}
