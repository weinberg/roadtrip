package main

import (
  "bufio"
  "context"
  "fmt"
  "github.com/brickshot/roadtrip/internal/client/config"
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

func showTrip() {
  trip, err := client.GetCarTrip(getCtx(), &psgrpc.GetCarTripRequest{
    CarId: character.Car.Id,
  })
  if err != nil {
    log.Fatalln("Cannot get trip: ", err)
  }

  fmt.Printf("Current trip plan:\n")
  for i, e := range trip.Entries {
    if e.Type == "town" {
      if i == 0 {
        fmt.Printf("╔═ ")
      } else if i == len(trip.Entries)-1 {
        fmt.Printf("╚═ ")
      } else {
        fmt.Printf("╠═ ")
      }
      fmt.Printf("%v\n", e.Town.DisplayName)
    } else if e.Type == "road" {
      fmt.Printf("║\n║  %v\n║\n", e.Road.DisplayName)
    }
  }
}

func welcome() {
  title :=
      `
Welcome to...

    __ __              ______     
   '  )  )           /   /        
     /--' __ __.  __/ --/__  o _  
    /  \_(_)(_/|_(_/_(_// (_<_/_)_
                             /    
                          __/     
`
  fmt.Println(title)

}

func printStatus() {
  fmt.Printf("Character: %v\n", character.CharacterName)
  fmt.Printf("Car Plate: %v\n", character.Car.Plate)
  fmt.Printf("Location:\n")
  fmt.Printf("  Town : %v\n", character.Car.Location.TownId)
  fmt.Printf("  Road : %v\n", character.Car.Location.RoadId)
  fmt.Printf("  Position : %v\n", character.Car.Location.PositionMiles)

  town, err := client.GetTown(context.Background(), &psgrpc.GetTownRequest{Id: character.Car.Location.TownId})
  if err != nil {
    log.Fatalf("Cannot find town details: %v\n", err)
  }
  fmt.Printf("Town Details:\n")
  fmt.Printf("  State: %v\n", town.StateId)
  fmt.Printf("  Town : %v\n", town.DisplayName)
  fmt.Printf("  Info : %v\n", town.Description)
  showTrip()
}

// main
func main() {
  welcome()
  setup()
  printStatus()
}
