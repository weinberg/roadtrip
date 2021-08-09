package mongoData

import (
  "context"
  "fmt"
  "github.com/brickshot/roadtrip/internal/server/types"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "log"
  "time"
)

type MongoProvider struct {
  types.DataProvider
}

var client *mongo.Client
var database *mongo.Database

type Config struct {
  URI string
  types.InitConfig
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

  database = client.Database("roadtrip")

  /*
     List databases
  */
  databases, err := client.ListDatabaseNames(ctx, bson.M{})
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(databases)


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
