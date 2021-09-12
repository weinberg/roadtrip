package main

import (
  "context"
  "fmt"
  rpc "github.com/brickshot/roadtrip/internal/mapServer/grpc"
  "github.com/brickshot/roadtrip/internal/mapServer/mongoData"
  "google.golang.org/grpc"
  "log"
  "net"
)

const mongoURI = "mongodb://root:example@mongo-service:27017"
var dp mongoData.MongoProvider

const (
  port = "9067"
)

type mapServer struct {
  rpc.UnimplementedRoadTripMapServer
}

// GetTown gets a town by id
func (*mapServer) GetTown(ctx context.Context, request *rpc.GetTownRequest) (*rpc.Town, error) {
  t, err := dp.GetTown(request.Id)
  if err != nil {
    return nil, err
  }
  return &rpc.Town{
    Id:          t.Id,
    Description: t.Description,
    State:       t.State,
    DisplayName: t.DisplayName,
  }, nil
}

// GetRoad gets a road by id
func (*mapServer) GetRoad(ctx context.Context, request *rpc.GetRoadRequest) (*rpc.Road, error) {
  r, err := dp.GetRoad(request.Id)
  if err != nil {
    return nil, err
  }
  return &rpc.Road{
    Id:          r.Id,
    DisplayName: r.DisplayName,
    LengthMiles: r.LengthMiles,
    TownA:       r.TownAId,
    TownB:       r.TownBId,
  }, nil
}

/******************************
 *
 * Main
 *
 ******************************/

func main() {
  fmt.Println("MapServer started")
  address := "0.0.0.0:" + port
  lis, err := net.Listen("tcp", address)
  if err != nil {
    log.Fatalf("Error %v", err)
  }
  fmt.Println("Server is listening on %v...", address)
  fmt.Println("Connecting to data provider...")

  // MongoData
  dp = mongoData.MongoProvider{}.Init(mongoData.Config{URI: mongoURI})
  defer dp.Shutdown()

  s := grpc.NewServer()
  rpc.RegisterRoadTripMapServer(s, &mapServer{})

  s.Serve(lis)
}
