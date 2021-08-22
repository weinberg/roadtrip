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

var dp mongoData.MongoProvider

const (
  port = ":9067"
)

type mapServer struct {
  rpc.UnimplementedRoadTripMapServer
}

// GetTown gets a town by name
func (*mapServer) GetTown(ctx context.Context, request *rpc.GetTownRequest) (*rpc.Town, error) {
  t, err := dp.GetTownByName(request.Id)
  if err != nil {
    return nil, err
  }
  return &rpc.Town{
    Id:          t.Id,
    Description: t.Description,
    State:       t.State,
    DisplayName: t.Name,
  }, nil
}

/******************************
 *
 * Main
 *
 ******************************/

func main() {
  fmt.Println("Server started")
  address := "0.0.0.0" + port
  lis, err := net.Listen("tcp", address)
  if err != nil {
    log.Fatalf("Error %v", err)
  }
  fmt.Printf("Server is listening on %v...", address)

  fmt.Printf("Connecting to data provider...")

  // MongoData
  dp = mongoData.MongoProvider{}.Init(mongoData.Config{URI: "mongodb://root:example@localhost:27017"})
  defer dp.Shutdown()

  s := grpc.NewServer()
  rpc.RegisterRoadTripMapServer(s, &mapServer{})

  s.Serve(lis)
}
