package main

import (
  "context"
  "errors"
  "fmt"
  pb "github.com/brickshot/roadtrip/internal/grpc"
  "github.com/brickshot/roadtrip/internal/server/mongoData"
  . "github.com/brickshot/roadtrip/internal/server/types"
  "google.golang.org/grpc"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/metadata"
  "google.golang.org/grpc/status"
  "log"
  "net"
)

var dp DataProvider

const (
  port = ":9066"
)

type playerServer struct {
  pb.UnimplementedRoadTripPlayerServer
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

// NewCharacter creates a new character record in the datastore and returns it. The client specifies
// the UUID (not in the header but in the request). Errors if UUID already exists.
func (*playerServer) NewCharacter(ctx context.Context, request *pb.NewCharacterRequest) (*pb.Character, error) {
  _, error := dp.GetCharacter(request.Uuid)
  if error == nil {
    return nil, status.Errorf(codes.AlreadyExists, "NewCharacter: Already exists")
  }

  dp.StoreCharacter(Character{UUID: request.Uuid, Name: request.Name})

  return &pb.Character{
    Uuid: request.Uuid,
    Name: request.Name,
  }, nil
}

// GetCharacter returns the character with the given UUID. Errors if UUID not found.
func (*playerServer) GetCharacter(ctx context.Context, request *pb.Empty) (*pb.Character, error) {
  UUID, err := getUUID(ctx)
  if err != nil {
    return nil, err
  }

  c, error := dp.GetCharacter(UUID)

  if error != nil {
    return nil, status.Errorf(codes.NotFound, "GetCharacter: cannot find character with that UUID")
  }

  return &pb.Character{
    Uuid: c.UUID,
    Name: c.Name,
  }, nil

}

func (*playerServer) GetCar(ctx context.Context, request *pb.Empty) (*pb.Car, error) {
  UUID, err := getUUID(ctx)
  if err != nil {
    return nil, err
  }

  c, err := dp.GetCharacter(UUID)
  if err != nil {
    return nil, err
  }

  if c.Car == nil {
    return nil, status.Errorf(codes.NotFound, "GetCar: cannot find car with that UUID")
  }

  return &pb.Car{
    Name:  c.Car.Name,
    Plate: c.Car.Plate,
    Uuid:  c.Car.UUID,
  }, nil
}

func (*playerServer) NewCar(ctx context.Context, request *pb.NewCarRequest) (*pb.Car, error) {
  UUID, err := getUUID(ctx)
  if err != nil {
    return nil, err
  }

  car, err := dp.NewCar(Car{Name: request.Name})
  if err != nil {
    return nil, status.Errorf(codes.Internal, "Could not create car in data store")
  }

  _, err2 := dp.SetCar(UUID, car.UUID)
  if err2 != nil {
    return nil, status.Errorf(codes.Internal, "Could not assign car to character in data store")
  }

  return &pb.Car{
    Name:  car.Name,
    Plate: car.Plate,
    Uuid:  car.UUID,
  }, nil
}

/*
func (*playerServer) NameCar(ctx context.Context, request *pb.NameCarRequest) (*pb.Car, error) {
  carName = request.Name
  return &pb.Car{Plate: "1", Name: carName}, nil
}
*/

func main() {
  fmt.Println("Server started")
  address := "0.0.0.0" + port
  lis, err := net.Listen("tcp", address)
  if err != nil {
    log.Fatalf("Error %v", err)
  }
  fmt.Printf("Server is listening on %v...", address)

  fmt.Printf("Connecting to data provider...")
  dp = mongoData.Init(mongoData.Config{ URI: "mongodb://root:example@localhost:27017" })

  s := grpc.NewServer()
  pb.RegisterRoadTripPlayerServer(s, &playerServer{})

  s.Serve(lis)
}
