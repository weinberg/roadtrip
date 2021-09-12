package main

import (
  "context"
  "fmt"
  mgrpc "github.com/brickshot/roadtrip/internal/mapServer/grpc"
  "github.com/brickshot/roadtrip/internal/playerServer/mongoData"
  "google.golang.org/grpc"
  "log"
  "time"
)

var dp mongoData.MongoProvider
var mapClient mgrpc.RoadTripMapClient

// var dp memoryData.MemoryProvider

const (
  port          = "9066"
  mongoURI      = "mongodb://root:example@mongo-service:27017"
  mapServerHost = "docker-mapserver"
  mapServerPort = "9067"
)

const updateInterval = time.Second * 10

/******************************
 *
 * Main
 *
 ******************************/

func main() {
  fmt.Println("UpdateService started")
  fmt.Printf("Connecting to data provider...")

  // MongoData
  dp = mongoData.MongoProvider{}.Init(mongoData.Config{URI: mongoURI})
  defer dp.Shutdown()

  // MapClient
  setupMapClient()

  c := make(chan int)
  go mainLoop(c)
  <-c
}

/******************************
 *
 * Helper methods
 *
 ******************************/

func setupMapClient() {
  opts := grpc.WithInsecure()
  serverAddress := mapServerHost + ":" + mapServerPort
  cc, err := grpc.Dial(serverAddress, opts)
  if err != nil {
    log.Fatalln(err)
  }

  mapClient = mgrpc.NewRoadTripMapClient(cc)
}

/******************************
 *
 * Update
 *
 ******************************/

// mainLoop updates the world every 10 seconds
func mainLoop(ch chan int) {
  fmt.Println("Looping")

  for {
    fmt.Println("Updating world...")
    update()
    time.Sleep(updateInterval)
  }
  ch <- 0
}

// update makes all changes to world state
// Currently that means moving cars around.
func update() {
  moveCars()
}

func moveCars() {
  // get all cars and their trips
  // if car has velocity then compute new position
  // if new position is past the end of the current road then set car to town
  // if in a town put car on next road
  // if no next road then set velocity to 0
  cars, err := dp.GetCars()
  if err != nil {
    log.Fatalln(err)
  }
  fmt.Printf("UpdateService: %v cars\n", len(cars))
  for _, car := range cars {
    fmt.Printf("Car:\n")
    fmt.Printf("\tPlate:        %v\n", car.Plate)
    fmt.Printf("\tName:         %v\n", car.Name)
    fmt.Printf("\tMPH:          %v\n", car.VelocityMph)
    fmt.Printf("\tDirection:    %v\n", car.Direction)
    fmt.Printf("\tLast update:  %v\n", car.LastLocationUpdateTimeUnix)
    fmt.Printf("\tRoadID:        %v\n", car.Location.RoadId)
    fmt.Printf("\tTownID:        %v\n", car.Location.TownId)

    var road *mgrpc.Road
    var town *mgrpc.Town

    if car.Location.RoadId != "" { // on a road
      ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
      road, err = mapClient.GetRoad(ctx, &mgrpc.GetRoadRequest{
        Id: car.Location.RoadId,
      })
      if err != nil { continue }
    } else if car.Location.TownId != "" { // in a town
      ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
      town, err = mapClient.GetTown(ctx, &mgrpc.GetTownRequest{
        Id: car.Location.TownId,
      })
      if err != nil {
        fmt.Println("Error: ", err)
        continue
      }
    }

    fmt.Printf("Current Location:\n")
    if road != nil {
      fmt.Printf("\tOn Road:       %v\n", road.DisplayName)
      fmt.Printf("\tLength:        %v miles\n", car.Location.RoadId)
      fmt.Printf("\tPosition:      %v miles\n", car.Location.PositionMiles)
    } else if town != nil {
      fmt.Printf("\tIn Town:\n")
      fmt.Printf("\tName:          %v\n", town.DisplayName)
      fmt.Printf("\tDescription:    %v\n", town.Description)
    }
    if car.VelocityMph == 0 {
      fmt.Printf("Car velocity is 0. Not updating position.\n")
      continue
    }

    lastUpdate := time.Unix(car.LastLocationUpdateTimeUnix, 0)
    fmt.Printf("Last location update: %v\n", lastUpdate.Format(time.RFC850))

    now := time.Now()
    diff := now.Sub(lastUpdate)
    fmt.Printf("Duration since last update: %v\n", diff.String())

    vmps := float64(car.VelocityMph / 3600)
    newMiles := car.Location.PositionMiles + float32(vmps * diff.Seconds())
    fmt.Printf("Updating position to: %v...\n", newMiles)
    car.Location.PositionMiles = newMiles
    car.LastLocationUpdateTimeUnix = now.Unix()

    err := dp.UpdateCar(car)
    if err != nil {
      fmt.Printf("Error: %v\n", err)
    }
  }
}
