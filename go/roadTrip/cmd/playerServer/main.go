package main

import (
  "context"
  "errors"
  "fmt"
  . "github.com/brickshot/roadtrip/internal/playerServer"
  grpc2 "github.com/brickshot/roadtrip/internal/playerServer/grpc"
  "github.com/brickshot/roadtrip/internal/playerServer/mongoData"
  "google.golang.org/grpc"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/metadata"
  "google.golang.org/grpc/status"
  "log"
  "net"
)

var dp mongoData.MongoProvider

// var dp memoryData.MemoryProvider

const (
  port = ":9066"
)

type playerServer struct {
  grpc2.UnimplementedRoadTripPlayerServer
}

// getUUID returns the UUID from the metadata
func getUUID(ctx context.Context) (string, error) {
  md, ok := metadata.FromIncomingContext(ctx)
  if !ok {
    return "", status.Errorf(codes.PermissionDenied, "failed to get UUID from metadata")
  }
  if val, ok := md["character_uuid"]; ok {
    return val[0], nil
  }
  return "", errors.New("UUID not in metadata")
}

// CreateCharacter creates a new character record in the datastore and returns it.
// Car is a singleton associated with character. The new character will get a new car assigned to it.
func (*playerServer) CreateCharacter(ctx context.Context, request *grpc2.CreateCharacterRequest) (*grpc2.Character, error) {
  if request.CharacterName == "" {
    return &grpc2.Character{}, status.Errorf(codes.Internal, "Name required")
  }

  nc, err := dp.CreateCharacter(request.CharacterName)
  if err != nil {
    return &grpc2.Character{}, status.Errorf(codes.Internal, "Could not create character: "+err.Error())
  }

  // create car for new character
  car, err := dp.CreateCar(Car{}, nc)
  if err != nil {
    dp.DeleteCharacter(nc.Id)
    return &grpc2.Character{}, status.Errorf(codes.Internal, "Could not create car for new character: "+err.Error())
  }

  r := &grpc2.Character{
    Id: nc.Id,
    CharacterName: nc.Name,
    Car: &grpc2.Car{
      Id:  car.Id,
      Plate: car.Plate,
      CarName:  car.Name,
    },
  }

  return r, nil
}

// GetCharacter returns the character with the given id. Errors if id not found.
func (*playerServer) GetCharacter(ctx context.Context, request *grpc2.GetCharacterRequest) (*grpc2.Character, error) {
  contextUUID, err := getUUID(ctx)
  if err != nil {
    return nil, err
  }

  if contextUUID != request.Id  {
    return nil, status.Errorf(codes.PermissionDenied, "Permission denied for that character ID")
  }

  ch, err := dp.GetCharacter(contextUUID)
  if err != nil {
    return nil, status.Errorf(codes.NotFound, "cannot find character with that UUID")
  }

  if ch.Car == nil {
    return nil, status.Errorf(codes.NotFound, "cannot find Car for character")
  }

  result := &grpc2.Character{
    Id: ch.Id,
    CharacterName: ch.Name,
    Car: &grpc2.Car{
      Id:  ch.Car.Id,
      Plate: ch.Car.Plate,
      CarName:  ch.Car.Name,
      Location: &grpc2.Location{
        TownId:   ch.Car.Location.TownId,
        RoadId:   ch.Car.Location.RoadId,
        Position: ch.Car.Location.Position,
      },
    },
  }

  return result, nil
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
  grpc2.RegisterRoadTripPlayerServer(s, &playerServer{})

  s.Serve(lis)
}
