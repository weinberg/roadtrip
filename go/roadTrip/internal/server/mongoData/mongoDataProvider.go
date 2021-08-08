package mongoData

import (
  "context"
  "fmt"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "log"
  "time"
)

type Provider struct {
}

type Config struct {
  URI string
}

func Init(c Config) (p Provider) {
  p = Provider{}
  client, err := mongo.NewClient(options.Client().ApplyURI(c.URI))
  if err != nil {
    log.Fatal(err)
  }
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  err = client.Connect(ctx)
  if err != nil {
    log.Fatal(err)
  }
  defer client.Disconnect(ctx)

  /*
     List databases
  */
  databases, err := client.ListDatabaseNames(ctx, bson.M{})
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(databases)

  return p
}
