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
    bson.D{{"id", "states/washington"}, {"name", "Washington"}},
    bson.D{{"id", "states/oregon"}, {"name", "Oregon"}},
    bson.D{{"id", "states/idaho"}, {"name", "Idaho"}},
    bson.D{{"id", "states/utah"}, {"name", "Utah"}},
    bson.D{{"id", "states/wyoming"}, {"name", "Wyoming"}},
    bson.D{{"id", "states/colorado"}, {"name", "Colorado"}},
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
      "id":           "states/washington/towns/seattle",
      "display_name": "Seattle",
      "description":  "A seaport city on the West Coast of the United States.",
      "state":        "/states/washington",
    },
    bson.M{
      "id":           "states/washington/towns/ellensburg",
      "display_name": "Ellensburg",
      "description":  "Ellensburg, originally named Ellensburgh for the wife of town founder John Alden Shoudy, was founded in 1871.",
      "state":        "/states/washington",
    },
    bson.M{
      "id":           "states/oregon/towns/hermiston",
      "display_name": "Hermison",
      "description":  "Hermiston is a city in Umatilla County, Oregon, United States. Its population of 19,354 makes it the largest city in Eastern Oregon. (Wikipedia)",
      "state":        "/states/oregon",
    },
    bson.M{
      "id":           "states/idaho/towns/boise",
      "display_name": "Boise",
      "description":  "Boise is the capital city of Idaho.",
      "state":        "/states/oregon",
    },
    bson.M{
      "id":           "states/utah/towns/ogden",
      "display_name": "Ogden",
      "description":  "Ogden, Utah",
      "state":        "/states/utah",
    },
    bson.M{
      "id":           "states/wyoming/towns/evanston",
      "display_name": "Evanston",
      "description":  "The population was 12,359 at the 2010 census. It is located near the border with Utah.",
      "state":        "/states/wyoming",
    },
    bson.M{
      "id":           "states/wyoming/towns/cheyenne",
      "display_name": "Cheyenne",
      "description":  "Cheyenne is the capital city of Wyoming. It’s home to the Cheyenne Frontier Days Old West Museum.",
      "state":        "/states/wyoming",
    },
    bson.M{
      "id":           "states/colorado/towns/fortcollins",
      "display_name": "Fort Collins",
      "description":  "Fort Fun",
      "state":        "/states/colorado",
    },
    bson.M{
      "id":           "states/colorado/towns/denver",
      "display_name": "Denver",
      "description":  "The capitol of Colorado",
      "state":        "/states/colorado",
    },
  })
}

func SeedRoads() {
  var coll *mongo.Collection
  coll = database.Collection("roads")
  coll.InsertMany(context.Background(), []interface{}{
    bson.M{
      "id":        "roads/i-90",
      "name":      "I-90",
      "town_a_id": "states/washington/towns/seattle",
      "town_b_id": "states/washington/towns/ellensburg",
      "miles":     110,
    },
    bson.M{
      "id":        "roads/i-82",
      "name":      "I-82",
      "town_a_id": "states/washington/towns/ellensburg",
      "town_b_id": "states/oregon/towns/hermiston",
      "miles":     136,
    },
    bson.M{
      "id":        "roads/i-84-a",
      "name":      "I-84",
      "town_a_id": "states/oregon/towns/hermiston",
      "town_b_id": "states/idaho/towns/boise",
      "miles":     256,
    },
    bson.M{
      "id":        "roads/i-84-b",
      "name":      "I-84",
      "town_a_id": "states/idaho/towns/boise",
      "town_b_id": "states/utah/towns/ogden",
      "miles":     305,
    },
    bson.M{
      "id":        "roads/i-80-a",
      "name":      "I-80",
      "town_a_id": "states/utah/towns/ogden",
      "town_b_id": "states/wyoming/towns/evanston",
      "miles":     77,
    },
    bson.M{
      "id":        "roads/i-80-b",
      "name":      "I-80",
      "town_a_id": "states/wyoming/towns/evanston",
      "town_b_id": "states/wyoming/towns/cheyenne",
      "miles":     358,
    },
    bson.M{
      "id":        "roads/i-25-a",
      "name":      "I-25",
      "town_a_id": "states/wyoming/towns/cheyenne",
      "town_b_id": "states/colorado/towns/fortcollins",
      "miles":     47,
    },
    bson.M{
      "id":        "roads/i-25-b",
      "name":      "I-25",
      "town_a_id": "states/colorado/towns/fortcollins",
      "town_b_id": "states/colorado/towns/denver",
      "miles":     63,
    },
  })
}
