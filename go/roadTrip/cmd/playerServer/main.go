package main

import (
	"context"
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	mgrpc "github.com/brickshot/roadtrip/internal/mapServer/grpc"
	. "github.com/brickshot/roadtrip/internal/playerServer"
	pgrpc "github.com/brickshot/roadtrip/internal/playerServer/grpc"
	"github.com/brickshot/roadtrip/internal/playerServer/mongoData"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
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
	mapServerHost = "mapServer"
	mapServerPort = "9067"
)

type playerServer struct {
	pgrpc.UnimplementedRoadTripPlayerServer
}

/******************************
 *
 * Main
 *
 ******************************/

var tlsEnabled bool

func init() {
	flag.BoolVar(&tlsEnabled, "tls", false, "If true, use TLS. Defaults to false.")
}

func main() {
	fmt.Printf("PlayerServer started...\n")
	address := "0.0.0.0" + ":" + port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v...\n", address)

	fmt.Printf("Connecting to data provider...\n")

	// MongoData
	dp = mongoData.MongoProvider{}.Init(mongoData.Config{URI: mongoURI})
	defer dp.Shutdown()

	var s *grpc.Server

	if tlsEnabled {
		fmt.Printf("TLS enabled\n")
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}
		fmt.Printf("Found credentials\n")
		s = grpc.NewServer(
			grpc.Creds(tlsCredentials),
		)
	} else {
		fmt.Printf("TLS disabled\n")
		s = grpc.NewServer()
	}

	pgrpc.RegisterRoadTripPlayerServer(s, &playerServer{})

	// MapClient
	setupMapClient()

	fmt.Println("Ready\n")

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

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("/certs/server-cert.pem", "/certs/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

/******************************
 *
 * Player Service API methods
 *
 ******************************/

// CreateCharacter creates a new character record in the datastore and returns it.
// Car is a singleton associated with character. The new character will get a new car assigned to it.
func (*playerServer) CreateCharacter(ctx context.Context, request *pgrpc.CreateCharacterRequest) (*pgrpc.Character, error) {
	if request.CharacterName == "" {
		return &pgrpc.Character{}, status.Errorf(codes.Internal, "Name required")
	}

	nc, err := dp.CreateCharacter(request.CharacterName)
	if err != nil {
		return &pgrpc.Character{}, status.Errorf(codes.Internal, "Could not create character: "+err.Error())
	}

	// create car for new character
	// we always start in seattle for now
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	start, err := mapClient.GetTown(ctx, &mgrpc.GetTownRequest{Id: "states/washington/towns/seattle"})

	if err != nil {
		dp.DeleteCharacter(nc.Id)
		return &pgrpc.Character{}, status.Errorf(codes.Internal, "Could find starting town for new character: "+err.Error())
	}
	// create car
	car, err := dp.CreateCar(Car{}, nc, Location{TownId: start.Id})
	if err != nil {
		dp.DeleteCharacter(nc.Id)
		return &pgrpc.Character{}, status.Errorf(codes.Internal, "Could not create car for new character: "+err.Error())
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
	car.VelocityMph = 60
	car.LastLocationUpdateTimeUnix = time.Now().Unix()
	err = dp.UpdateCar(car)
	if err != nil {
		return &pgrpc.Character{}, status.Errorf(codes.Internal, "Could assign trip to new car: "+err.Error())
	}

	r := &pgrpc.Character{
		Id:            nc.Id,
		CharacterName: nc.Name,
		Car: &pgrpc.Car{
			Id:      car.Id,
			Plate:   car.Plate,
			CarName: car.Name,
		},
	}

	return r, nil
}

// GetCharacter returns the character with the given id. Errors if id not found.
func (*playerServer) GetCharacter(ctx context.Context, request *pgrpc.GetCharacterRequest) (*pgrpc.Character, error) {
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

	result := &pgrpc.Character{
		Id:            ch.Id,
		CharacterName: ch.Name,
		Car: &pgrpc.Car{
			Id:            ch.Car.Id,
			Plate:         ch.Car.Plate,
			CarName:       ch.Car.Name,
			OdometerMiles: ch.Car.OdometerMiles,
			VelocityMph:   ch.Car.VelocityMph,
			Location: &pgrpc.Location{
				TownId:        ch.Car.Location.TownId,
				RoadId:        ch.Car.Location.RoadId,
				PositionMiles: ch.Car.Location.PositionMiles,
			},
		},
	}

	return result, nil
}

// GetCar returns the car if the player has access to it.
func (*playerServer) GetCar(ctx context.Context, request *pgrpc.GetCarRequest) (*pgrpc.Car, error) {
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

	result := &pgrpc.Car{
		Id:            ch.Car.Id,
		Plate:         ch.Car.Plate,
		CarName:       ch.Car.Name,
		OdometerMiles: ch.Car.OdometerMiles,
		Location: &pgrpc.Location{
			TownId:        ch.Car.Location.TownId,
			RoadId:        ch.Car.Location.RoadId,
			PositionMiles: ch.Car.Location.PositionMiles,
		},
	}

	return result, nil
}

// GetCarTrip returns the car's trip if the player has access to the car.
func (*playerServer) GetCarTrip(ctx context.Context, request *pgrpc.GetCarTripRequest) (*pgrpc.Trip, error) {
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
		return &pgrpc.Trip{}, nil
	}

	var entries []*pgrpc.TripEntry = []*pgrpc.TripEntry{}
	ctx, _ = context.WithTimeout(context.Background(), time.Second*20)
	for _, e := range ch.Car.Trip.Entries {
		tripEntry := pgrpc.TripEntry{Type: e.Type}
		if e.Type == "road" {
			road, err := mapClient.GetRoad(ctx, &mgrpc.GetRoadRequest{Id: e.Id})
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Error loading a trip entry (road: %v) from db: ", e.Id, err)
			}
			tripEntry.Road = &pgrpc.Road{
				Id:          road.Id,
				DisplayName: road.DisplayName,
				LengthMiles: road.LengthMiles,
			}
		} else if e.Type == "town" {
			town, err := mapClient.GetTown(ctx, &mgrpc.GetTownRequest{Id: e.Id})
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Error loading a trip entry (town: %v) from db: ", e.Id, err)
			}
			tripEntry.Town = &pgrpc.Town{
				Id:          town.Id,
				StateId:     town.StateId,
				DisplayName: town.DisplayName,
				Description: town.Description,
			}
		}

		// entries = append(entries, &pgrpc.TripEntry{ Id:   e.Id, Type: e.Type, DisplayName: displayName})
		entries = append(entries, &tripEntry)
	}

	for _, e := range entries {
		fmt.Printf("tripEntry: %+v\n", e)
	}

	return &pgrpc.Trip{
		Entries: entries,
	}, nil
}

// GetTown returns the requested town. Proxies request to the map service.
func (*playerServer) GetTown(ctx context.Context, request *pgrpc.GetTownRequest) (*pgrpc.Town, error) {
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)
	town, err := mapClient.GetTown(ctx, &mgrpc.GetTownRequest{Id: request.Id})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot find town")
	}
	result := &pgrpc.Town{
		Id:          town.Id,
		StateId:     town.StateId,
		DisplayName: town.DisplayName,
		Description: town.Description,
	}
	return result, nil
}

// GetRoad returns the requested town. Proxies request to the map service.
func (*playerServer) GetRoad(ctx context.Context, request *pgrpc.GetRoadRequest) (*pgrpc.Road, error) {
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)
	road, err := mapClient.GetRoad(ctx, &mgrpc.GetRoadRequest{Id: request.Id})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot find road")
	}
	result := &pgrpc.Road{
		Id:          road.Id,
		DisplayName: road.DisplayName,
		LengthMiles: road.LengthMiles,
	}
	return result, nil
}

// GetUpdate returns an update packet
func (*playerServer) GetUpdate(ctx context.Context, request *pgrpc.GetUpdateRequest) (*pgrpc.Update, error) {
	contextUUID, err := getUUID(ctx)
	if err != nil {
		return nil, err
	}

	ch, err := dp.GetCharacter(contextUUID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot find character with that UUID")
	}
	if ch.Car == nil {
		return nil, status.Errorf(codes.NotFound, "cannot find Car for character")
	}

	var townName string
	var roadName string

	if ch.Car.Location.TownId != "" {
		ctx, _ = context.WithTimeout(ctx, 10*time.Second)
		town, err := mapClient.GetTown(ctx, &mgrpc.GetTownRequest{Id: ch.Car.Location.TownId})
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "cannot find town")
		}
		townName = town.DisplayName
	}

	if ch.Car.Location.RoadId != "" {
		ctx, _ = context.WithTimeout(ctx, 10*time.Second)
		road, err := mapClient.GetRoad(ctx, &mgrpc.GetRoadRequest{Id: ch.Car.Location.RoadId})
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "cannot find road")
		}
		roadName = road.DisplayName
	}

	result := pgrpc.Update{
		RoadName: roadName,
		TownName: townName,
		Character: &pgrpc.Character{
			Id:            ch.Id,
			CharacterName: ch.Name,
			Car: &pgrpc.Car{
				Id:            ch.Car.Id,
				Plate:         ch.Car.Plate,
				CarName:       ch.Car.Name,
				VelocityMph:   ch.Car.VelocityMph,
				OdometerMiles: ch.Car.OdometerMiles,
				Location: &pgrpc.Location{
					TownId:        ch.Car.Location.TownId,
					RoadId:        ch.Car.Location.RoadId,
					PositionMiles: ch.Car.Location.PositionMiles,
				},
			},
		},
	}

	return &result, nil
}
