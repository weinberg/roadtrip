package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/brickshot/roadtrip/internal/client/config"
	"github.com/brickshot/roadtrip/internal/client/ui"
	psgrpc "github.com/brickshot/roadtrip/internal/playerServer/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"strings"
	"time"
)

var id string
var reader *bufio.Reader
var client psgrpc.RoadTripPlayerClient
var character *psgrpc.Character
var update *psgrpc.Update
var trip *psgrpc.Trip
var screen ui.Screen

func getCtx() context.Context {
	// add id to grpc headers. for now, we only allow one character
	md := metadata.New(map[string]string{"character_uuid": id})
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)

	return ctx
}

func setup() {
	/*
	   Config file has list of characterInfo which have id
	   if characterInfo list is empty, add one
	   setup grpc with first characterInfo id
	   take first entry in characterInfo list and get character from server
	   now we have an id for the grpc header and a pb.character
	*/
	// read config file. LoadConfig will create config if it doesn't exist.
	// the new config file will have no characterInfos
	reader = bufio.NewReader(os.Stdin)
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Config file error: ", err)
	}
	setupGrpc(conf)

	// if character list is empty, add one
	if conf.Characters == nil || len(conf.Characters) == 0 {
		createNewCharacter()
		conf, err = config.LoadConfig()
		if err != nil {
			log.Fatalln("Config file error: ", err)
		}
	}

	// currently, only use first character
	id = conf.Characters[0].Id

	// get character from server
	character, err = client.GetCharacter(getCtx(), &psgrpc.GetCharacterRequest{Id: id})
	if err != nil {
		st := status.Convert(err)
		if st.Code() == codes.NotFound {
			log.Println("Sorry your character was not found on the server.")
			fmt.Printf("Would you like to delete this character? (yes to delete): ")
			r, _ := reader.ReadString('\n')
			if r != "yes\n" {
				fmt.Println("OK, exiting.")
				os.Exit(0)
			}
			_, err = config.DeleteCharacterInfo(id)
			if err != nil {
				log.Fatalln("Config file error ", err)
			}
			log.Println("Deleted. Try starting over.")
			os.Exit(0)
		} else if st.Code() == codes.Unavailable {
			log.Fatalf("Cannot contact the server at %v:%v (%v)\n", conf.Server, conf.Port, err)
		} else {
			log.Fatalf("An error occured when starting up: %v\n", err)
		}
	}

	trip, err = client.GetCarTrip(getCtx(), &psgrpc.GetCarTripRequest{
		CarId: character.Car.Id,
	})
	if err != nil {
		log.Fatalln("Cannot get trip: ", err)
	}
}

func setupGrpc(conf config.ClientConfig) {
	opts := grpc.WithInsecure()
	serverAddress := conf.Server + ":" + conf.Port
	fmt.Printf("Connecting to %v\n", serverAddress)
	cc, err := grpc.Dial(serverAddress, opts)
	if err != nil {
		log.Fatalln(err)
	}

	client = psgrpc.NewRoadTripPlayerClient(cc)
}

func createNewCharacter() *psgrpc.Character {
	fmt.Println("Creating a new character...")
	var name string
	for name == "" {
		fmt.Printf("What would you like your name to be?  ")
		name, _ = reader.ReadString('\n')
		name = strings.TrimRight(name, "\r\n")
		if name == "" {
			fmt.Println("That name is too short.")
		}
	}

	// create in server
	char, err := client.CreateCharacter(getCtx(), &psgrpc.CreateCharacterRequest{
		CaptchaId:     "",
		CharacterName: name,
	})
	st := status.Convert(err)
	if st != nil {
		log.Fatalf("Failed to create character: %v\n", st.Err())
	}

	// store in config
	_, _, err = config.NewCharacterInfo(char.Id)
	if err != nil {
		log.Fatalln("Cannot create new character: ", err)
	}

	return char
}

func getUpdate() *psgrpc.Update {
	var err error
	update, err = client.GetUpdate(getCtx(), &psgrpc.GetUpdateRequest{})
	if err != nil {
		log.Fatalln("Cannot get update: ", err)
	}
	return update
}

func roadTripTitle() {
	text := `
    __ __              ______     
   '  )  )           /   /        
     /--' __ __.  __/ --/__  o _  
    /  \_(_)(_/|_(_/_(_// (_<_/_)_
                             /    
                          __/     
`
	fmt.Println(text)
}

// main
func main() {
	setup()

	trip, err := client.GetCarTrip(getCtx(), &psgrpc.GetCarTripRequest{
		CarId: character.Car.Id,
	})
	if err != nil {
		log.Fatalln("Cannot get trip: ", err)
	}

	screen = ui.Screen{Width: 80, Height: 25}
	updates := make(chan *psgrpc.Update)
	go screen.RenderUpdate(updates, trip)
	for {
		updates <- getUpdate()
		time.Sleep(10 * time.Second)
	}
}
