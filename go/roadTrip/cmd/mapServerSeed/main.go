package main

import (
  "context"
  "flag"
  "fmt"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "log"
  "time"
)

var database *mongo.Database

const dbName = "rtMapDB"

/******************************
 *
 * Init
 *
 ******************************/

var skipCleanFlag bool
func init() {
  flag.BoolVar(&skipCleanFlag, "skipclean", false, "If true, do not delete the database prior to seeding")
}

func main() {
  fmt.Println("Connecting to data provider...")

  client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
  if err != nil {
    log.Fatal(err)
  }
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  err = client.Connect(ctx)
  if err != nil {
    log.Fatal(err)
  }
  defer client.Disconnect(context.Background())

  database = client.Database(dbName)

  if !skipCleanFlag {
    fmt.Println("Dropping map database...")
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    database.Drop(ctx)
    if ctx.Err() == context.DeadlineExceeded {
      log.Fatalf("Deadline exceeded talking to db: %v\n", err)
    }
  }

  fmt.Println("Seeding map data...")
  SeedStates()
  SeedTowns()
  SeedRoads()
  fmt.Println("Done")
}

func SeedStates() {
  var coll *mongo.Collection
  coll = database.Collection("states")
  _, err := coll.InsertMany(context.Background(), []interface{}{
    bson.D{{"id", "states/colorado"}, {"name", "Colorado"}},
    bson.D{{"id", "states/newmexico"}, {"name", "New Mexico"}},
  })
  if err != nil {
    log.Fatal(err)
  }
}

func SeedTowns() {
  var coll *mongo.Collection
  coll = database.Collection("towns")
  coll.InsertMany(context.Background(), []interface{}{
    bson.M{
      "id":          "states/washington/towns/seattle",
      "name":        "Seattle",
      "description": "A seaport city on the West Coast of the United States.",
      "state":       "/states/washington",
    },
    bson.M{
      "id":          "states/colorado/towns/denver",
      "name":        "Denver",
      "description": "The capitol of Colorado",
      "state":       "/states/colorado",
    },
    bson.M{
      "id":          "states/colorado/towns/grandjunction",
      "name":        "Grand Junction",
      "description": "A mid-sized Western Slope town.",
      "state":       "/states/colorado",
    },
    bson.M{
      "id":          "states/colorado/towns/breckenridge",
      "name":        "Breckenridge",
      "description": "A ski town.",
      "state":       "/states/colorado",
    },
  })
}

func SeedRoads() {
  var coll *mongo.Collection
  coll = database.Collection("roads")
  coll.InsertMany(context.Background(), []interface{}{
    bson.M{
      "id":    "states/colorado/roads/i-25",
      "name":  "I-70",
      "town_a_id": "states/colorado/towns/denver",
      "town_b_id": "states/colorado/towns/breckenridge",
    },
  })
}