package mongoData

import (
  "context"
  "github.com/brickshot/roadtrip/internal/mapServer"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "log"
  "time"
)

const dbName = "rtMapDB"

type MongoProvider struct {
  mapServer.DataProvider
}

var client *mongo.Client
var database *mongo.Database

type Config struct {
  URI string
  mapServer.InitConfig
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
    log.Fatal(err)
  }

  database = client.Database(dbName)

  InitStates()
  InitTowns()
  InitRoads()

  return p
}

func (p MongoProvider) Shutdown() {
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  if err := client.Disconnect(ctx); err != nil {
    panic(err)
  }
}
