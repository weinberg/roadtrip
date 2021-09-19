package mongoData

import (
  "context"
  "github.com/brickshot/roadtrip/internal/playerServer"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "log"
  "time"
)

const dbName = "rtPlayerDB"

type MongoProvider struct {
  playerServer.DataProvider
}

var client *mongo.Client
var database *mongo.Database

type Config struct {
  URI string
  playerServer.InitConfig
}

func (p MongoProvider) Init (c Config) MongoProvider {
  var err error
  client, err = mongo.NewClient(options.Client().ApplyURI(c.URI))
  if err != nil {
    log.Fatal(err)
  }
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  err = client.Connect(ctx)
  if err != nil {
    log.Fatal("mongoDataProvider:Init:", err)
  }

  database = client.Database(dbName)

  // Init collections
  InitCharacters()
  InitCars()
  return p
}

func (p MongoProvider) Shutdown() {
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  if err := client.Disconnect(ctx); err != nil {
    panic(err)
  }
}
