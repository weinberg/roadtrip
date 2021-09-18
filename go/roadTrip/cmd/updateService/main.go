package main

import (
	"context"
	"errors"
	"fmt"
	mgrpc "github.com/brickshot/roadtrip/internal/mapServer/grpc"
	. "github.com/brickshot/roadtrip/internal/playerServer"
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
	mapServerHost = "mapServer"
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
		fmt.Printf("\tSpeed:         %v MPH\n", car.VelocityMph)
		fmt.Printf("\tOdometer:      %v miles\n", car.OdometerMiles)
		fmt.Printf("\tLast update:   %v\n", car.LastLocationUpdateTimeUnix)
		fmt.Printf("\tRoadID:        %v\n", car.Location.RoadId)
		fmt.Printf("\tTownID:        %v\n", car.Location.TownId)
		fmt.Printf("\tMiles on road: %v\n", car.Location.PositionMiles)

		if car.Location.RoadId != "" { // on a road
			handleRoad(car)
		} else if car.Location.TownId != "" { // in a town
			handleTown(car)
		}

	}
}

// nextTripEntry returns the next trip entry for the car. If the car is at the end of the trip then nil.
func nextTripEntry(car Car) (*TripEntry, error) {
	var found bool = false
	var i int
	for _, l := range car.Trip.Entries {
		if l.Id == car.Location.RoadId || l.Id == car.Location.TownId {
			found = true
			break
		}
		i++
	}

	if !found {
		return nil, errors.New("Car is at a location which is not found in it's own trip entries.")
	}

	i++
	if i >= len(car.Trip.Entries) {
		return nil, nil
	}
	return &car.Trip.Entries[i], nil
}

// handleTown when a car is in a town
func handleTown(car Car) error {
	next, err := nextTripEntry(car)
	if err != nil {
		return err
	}
	if next == nil {
		fmt.Printf("HandleTown: No next trip entry. End of trip\n")
		car.VelocityMph = 0
	} else {
		car.Location = &Location{
			RoadId:        next.Id, // next location from a town is always a road
			PositionMiles: 0,
		}
	}

	car.LastLocationUpdateTimeUnix = time.Now().Unix()
	err = dp.UpdateCar(car)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return nil
}

// handleRoad when a car is on a road
func handleRoad(car Car) error {
	if car.VelocityMph == 0 {
		fmt.Printf("Car velocity is 0. Not updating position.\n")
		return nil
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// need road length
	road, err := mapClient.GetRoad(ctx, &mgrpc.GetRoadRequest{
		Id: car.Location.RoadId,
	})
	if err != nil {
		return err
	}

	lastUpdate := time.Unix(car.LastLocationUpdateTimeUnix, 0)
	now := time.Now()
	diff := now.Sub(lastUpdate)
	vmps := float64(car.VelocityMph / 3600)
	addMiles := float32(vmps * diff.Seconds())
	newMiles := car.Location.PositionMiles + addMiles
	fmt.Printf("HandleRoad: Updating position to: %v...\n", newMiles)

	// Handle the car going past the end of the road
	if newMiles >= float32(road.LengthMiles) {
		addMiles = float32(road.LengthMiles) - car.Location.PositionMiles
		fmt.Printf("HandleRoad: New position is past end of road\n")
		next, err := nextTripEntry(car)
		if err != nil {
			return err
		}
		if next == nil {
			// dunno.. trip should end with a town not a road but whatevs
			fmt.Printf("HandleRoad: No next trip entry. End of trip\n")
			car.VelocityMph = 0
		} else {
			fmt.Printf("HandleRoad: Next tripEntry is %+v\n", next)
			car.Location = &Location{
				TownId:        next.Id, // next from a road is always a town
				PositionMiles: 0,
			}
		}
	} else {
		car.Location.PositionMiles = newMiles
	}
	car.OdometerMiles += addMiles
	car.LastLocationUpdateTimeUnix = now.Unix()
	err = dp.UpdateCar(car)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return nil
}
