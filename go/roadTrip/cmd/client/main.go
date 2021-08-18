package main

import (
  "bufio"
  "context"
  "fmt"
  "github.com/brickshot/roadtrip/internal/client/config"
  pb "github.com/brickshot/roadtrip/internal/grpc"
  "google.golang.org/grpc"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/metadata"
  "google.golang.org/grpc/status"
  "log"
  "os"
  "strings"
)

var cc *grpc.ClientConn
var id string
var reader *bufio.Reader
var client pb.RoadTripPlayerClient
var ctx context.Context
var character *pb.Character

func setup() {
  /*
     Config file has list of characterInfo which have id
     if characterInfo list is empty, add one
     setup grpc with first characterInfo id
     take first entry in characterInfo list and get character from server
     now we have a id for the grpc header and a pb.character
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

  // currently only use first character
  id = conf.Characters[0].Id

  // add id to grpc headers. for now we only allow one character
  md := metadata.New(map[string]string{"character_uuid": id})
  ctx = metadata.NewOutgoingContext(context.Background(), md)

  // get character from server
  character, err = client.GetCharacter(ctx, &pb.GetCharacterRequest{Id: id})
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
    } else {
      log.Fatalln("Error loading character")
    }
  }
}

func setupGrpc(conf config.ClientConfig) {
  opts := grpc.WithInsecure()
  serverAddress := conf.Server + ":" + conf.Port
  cc, err := grpc.Dial(serverAddress, opts)
  if err != nil {
    log.Fatalln(err)
  }

  client = pb.NewRoadTripPlayerClient(cc)
  md := metadata.New(map[string]string{"character_uuid": id})
  ctx = metadata.NewOutgoingContext(context.Background(), md)
}

func createNewCharacter() *pb.Character {
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
  char, err := client.CreateCharacter(ctx, &pb.CreateCharacterRequest{
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

// main
func main() {
  fmt.Println("Welcome to RoadTrip")
  setup()

  fmt.Printf("Character: \"%v\"\n", character)
  fmt.Printf("Car: \"%v\"\n", character.Car)
}
