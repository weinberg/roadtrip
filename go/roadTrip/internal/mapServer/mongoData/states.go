package mongoData

import (
  context "context"
  . "github.com/brickshot/roadtrip/internal/mapServer"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "time"
)

var statesColl *mongo.Collection

// init
func InitStates() {
  statesColl = database.Collection("states")
}

func ShutdownStates() {
  statesColl = nil
}

// GetStates returns all the states
func (d MongoProvider) GetStates(id string) ([]State, error) {
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  cur, err := statesColl.Find(ctx, bson.D{})

  if err != nil {
    return []State{}, err
  }

  var states []State
  if err = cur.All(ctx, &states); err != nil {
    return []State{}, err
  }

  return states, nil
}
