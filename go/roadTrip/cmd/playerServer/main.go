package main

import (
  "context"
  "errors"
  "fmt"
  mgrpc "github.com/brickshot/roadtrip/internal/mapServer/grpc"
  . "github.com/brickshot/roadtrip/internal/playerServer"
  grpc2 "github.com/brickshot/roadtrip/internal/playerServer/grpc"
  "github.com/brickshot/roadtrip/internal/playerServer/mongoData"
  "google.golang.org/grpc"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/metadata"
  "google.golang.org/grpc/status"
  "log"
  "net"
  "time"
)

var dp mongoData.MongoProvider
var mapClient mgrpc.RoadTripMapClient

const (
  port          = "9066"
  mongoURI      = "mongodb://root:example@mongo-service:27017"
  mapServerHost = "docker-mapserver"
  mapServerPort = "9067"
)

type playerServer struct {
  grpc2.UnimplementedRoadTripPlayerServer
}

/******************************
 *
 * Main
 *
 ******************************/

func main() {
  fmt.Println("PlayerServer started")
  address := "0.0.0.0" + ":" + port
  lis, err := net.Listen("tcp", address)
  if err != nil {
    log.Fatalf("Error %v", err)
  }
  fmt.Printf("Server is listening on %v...", address)

  fmt.Printf("Connecting to data provider...")

  // MongoData
  dp = mongoData.MongoProvider{}.Init(mongoData.Config{URI: mongoURI})
  defer dp.Shutdown()

  // MapClient
  setupMapClient()

  s := grpc.NewServer()
  grpc2.RegisterRoadTripPlayerServer(s, &playerServer{})

  fmt.Println("Ready")

  s.Serve(lis)
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

/******************************
 *
 * Player Service API methods
 *
 ******************************/

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
  // we always start in seattle for now
  ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
  start, err := mapClient.GetTown(ctx, &mgrpc.GetTownRequest{Id: "states/washington/towns/seattle"})
  if err != nil {
    dp.DeleteCharacter(nc.Id)
    return &grpc2.Character{}, status.Errorf(codes.Internal, "Could find starting town for new character: "+err.Error())
  }
  // create car
  car, err := dp.CreateCar(Car{}, nc, Location{TownId: start.Id})
  if err != nil {
    dp.DeleteCharacter(nc.Id)
    return &grpc2.Character{}, status.Errorf(codes.Internal, "Could not create car for new character: "+err.Error())
  }
  // update car with trip
  // hardcoded for now
  car.Trip = &Trip{
    Entries: []TripEntry{
      TripEntry{Id: "states/washington/towns/seattle", Type: "town"},
      TripEntry{Id: "roads/i-90", Type: "road"},
      TripEntry{Id: "states/washington/towns/ellensburg", Type: "town"},
      TripEntry{Id: "roads/i-82", Type: "road"},
      TripEntry{Id: "states/oregon/towns/hermiston", Type: "town"},
      TripEntry{Id: "roads/i-84-a", Type: "road"},
      TripEntry{Id: "states/idaho/towns/boise", Type: "town"},
      TripEntry{Id: "roads/i-84-b", Type: "road"},
      TripEntry{Id: "states/utah/towns/ogden", Type: "town"},
      TripEntry{Id: "roads/i-80-a", Type: "road"},
      TripEntry{Id: "states/wyoming/towns/evanston", Type: "town"},
      TripEntry{Id: "roads/i-80-b", Type: "road"},
      TripEntry{Id: "states/wyoming/towns/cheyenne", Type: "town"},
    },
  }
  err = dp.UpdateCar(car)
  if err != nil {
    return &grpc2.Character{}, status.Errorf(codes.Internal, "Could assign trip to new car: "+err.Error())
  }

  r := &grpc2.Character{
    Id:            nc.Id,
    CharacterName: nc.Name,
    Car: &grpc2.Car{
      Id:      car.Id,
      Plate:   car.Plate,
      CarName: car.Name,
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

  if contextUUID != request.Id {
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
    Id:            ch.Id,
    CharacterName: ch.Name,
    Car: &grpc2.Car{
      Id:      ch.Car.Id,
      Plate:   ch.Car.Plate,
      CarName: ch.Car.Name,
      Location: &grpc2.Location{
        TownId:        ch.Car.Location.TownId,
        RoadId:        ch.Car.Location.RoadId,
        PositionMiles: ch.Car.Location.PositionMiles,
      },
    },
  }

  return result, nil
}

// GetCar returns the car if the player has access to it.
func (*playerServer) GetCar(ctx context.Context, request *grpc2.GetCarRequest) (*grpc2.Car, error) {
  // currently we only allow you to get info on your own car so we look up the requested car on your own char
  contextUUID, err := getUUID(ctx)
  if err != nil {
    return nil, err
  }
  ch, err := dp.GetCharacter(contextUUID)
  if err != nil {
    return nil, status.Errorf(codes.NotFound, "cannot find your character")
  }

  if ch.Car == nil {
    return nil, status.Errorf(codes.NotFound, "character is missing car")
  }

  if ch.Car.Id != request.Id {
    return nil, status.Errorf(codes.NotFound, "can only get your own car currently")
  }

  result := &grpc2.Car{
    Id:      ch.Car.Id,
    Plate:   ch.Car.Plate,
    CarName: ch.Car.Name,
    Location: &grpc2.Location{
      TownId:        ch.Car.Location.TownId,
      RoadId:        ch.Car.Location.RoadId,
      PositionMiles: ch.Car.Location.PositionMiles,
    },
  }

  return result, nil
}

// GetCarTrip returns the car's trip if the player has access to the car.
func (*playerServer) GetCarTrip(ctx context.Context, request *grpc2.GetCarTripRequest) (*grpc2.Trip, error) {
  // currently we only allow you to get info on your own car so we look up the requested car on your own char
  contextUUID, err := getUUID(ctx)
  if err != nil {
    return nil, err
  }
  // player has access to car if they own it or (eventually) if they are riding in it
  ch, err := dp.GetCharacter(contextUUID)
  if err != nil {
    return nil, status.Errorf(codes.NotFound, "cannot find your character")
  }

  if ch.Car == nil {
    return nil, status.Errorf(codes.NotFound, "character is missing car")
  }

  if ch.Car.Id != request.CarId {
    return nil, status.Errorf(codes.NotFound, "can only get trip for your own car currently")
  }

  if ch.Car.Trip == nil {
    return &grpc2.Trip{}, nil
  }

  var entries []*grpc2.TripEntry = []*grpc2.TripEntry{}
  for _,e := range ch.Car.Trip.Entries {
    entries = append(entries, &grpc2.TripEntry{ Id:   e.Id, Type: e.Type })
  }

  return &grpc2.Trip{
    Entries: entries,
  }, nil
}

// GetTown returns the requested town. Proxies request to the map service.
func (*playerServer) GetTown(ctx context.Context, request *grpc2.GetTownRequest) (*grpc2.Town, error) {
  ctx, _ = context.WithTimeout(ctx, 10*time.Second)
  town, err := mapClient.GetTown(ctx, &mgrpc.GetTownRequest{Id: request.Id})
  if err != nil {
    return nil, status.Errorf(codes.NotFound, "cannot find town")
  }
  result := &grpc2.Town{
    Id:          town.Id,
    StateId:     town.State,
    TownName:    town.DisplayName,
    Description: town.Description,
  }
  return result, nil
}
