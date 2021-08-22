package main

import (
  "context"
  "flag"
  "fmt"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "log"
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
  err = client.Connect(context.Background())
  if err != nil {
    log.Fatal(err)
  }
  defer client.Disconnect(context.Background())

  database = client.Database(dbName)

  if !skipCleanFlag {
    fmt.Println("Dropping map database...")
    database.Drop(context.Background())
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
    bson.D{{"Id", "/states/Colorado"}, {"Name", "Colorado"}},
    bson.D{{"Id", "/states/NewMexico"}, {"Name", "New Mexico"}},
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
      "Id":          "/states/Colorado/towns/Denver",
      "Name":        "Denver",
      "Description": "The capitol of Colorado",
      "State":       "/states/Colorado",
    },
    bson.M{
      "Id":          "/states/Colorado/towns/GrandJunction",
      "Name":        "Grand Junction",
      "Description": "A mid-sized Western Slope town.",
      "State":       "/states/Colorado",
    },
    bson.M{
      "Id":          "/states/Colorado/towns/Breckenridge",
      "Name":        "Breckenridge",
      "Description": "A cozy ski town.",
      "State":       "/states/Colorado",
    },
  })
}

func SeedRoads() {
  var coll *mongo.Collection
  coll = database.Collection("roads")
  coll.InsertMany(context.Background(), []interface{}{
    bson.M{
      "Id":    "/states/Colorado/roads/I-25",
      "Name":  "I-70",
      "TownA": "/states/Colorado/towns/Denver",
      "TownB": "/states/Colorado/towns/Breckenridge",
    },
  })
}