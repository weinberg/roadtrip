package main

import (
  "bufio"
  "context"
  "errors"
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
var UUID string
var reader *bufio.Reader
var client pb.RoadTripPlayerClient
var ctx context.Context

func setup() {
  reader = bufio.NewReader(os.Stdin)
  conf, err := config.LoadConfig()
  if err != nil {
    log.Fatalln("Config file error: " + err.Error())
  }

  UUID = conf.Characters[0].UUID

  opts := grpc.WithInsecure()
  serverAddress := conf.Server + ":" + conf.Port
  cc, err = grpc.Dial(serverAddress, opts)
  if err != nil {
    log.Fatalln(err)
  }

  client = pb.NewRoadTripPlayerClient(cc)

  md := metadata.New(map[string]string{"character_uuid": UUID})
  ctx = metadata.NewOutgoingContext(context.Background(), md)
}

func shutdown() {
  err := cc.Close()
  if err != nil {
    log.Fatalln(err)
  }
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

  char, err := client.NewCharacter(ctx, &pb.NewCharacterRequest{
    Uuid: UUID,
    Name: name,
  })
  st := status.Convert(err)
  if st != nil {
    log.Fatalf("Failed to create character: %v\n", st.Err())
  }

  return char
}

// getCharacter gets the character from the server or creates a new one and stores it
func getCharacter() (char *pb.Character) {
  if UUID == "" {
    char = createNewCharacter()
  } else {
    var err error
    for char == nil {
      char, err = client.GetCharacter(ctx, &pb.Empty{})
      if err != nil {
        st := status.Convert(err)
        if st.Code() == codes.NotFound {
          log.Println("Sorry the character referenced by your UUID is not found on the server.")
          fmt.Printf("Would you like to start over? (Enter yes to start over with a new character): ")
          r, _ := reader.ReadString('\n')
          if r != "yes\n" {
            fmt.Println("OK, exiting.")
            os.Exit(0)
          }
          char = createNewCharacter()
        } else {
          log.Fatalln("Error loading character")
        }
      }
      fmt.Printf("Welcome back, %v\n", char.Name)
    }
  }

  return
}

// getCharacter gets the character from the server or creates a new one and stores it
func getCar() (*pb.Car, error) {
  var car *pb.Car
  if UUID == "" {
    return nil, errors.New("invalid UUID")
  } else {
    var err error
    for car == nil {
      car, err = client.GetCar(ctx, &pb.Empty{})
      if err != nil {
        st := status.Convert(err)
        if st.Code() == codes.NotFound {
          fmt.Printf("You don't have a car. Let's make one...\n")
          var name string
          for name == "" {
            fmt.Printf("What do you want to name your car? ")
            name, _ = reader.ReadString('\n')
            name = strings.TrimRight(name, "\r\n")
            if name == "" {
              fmt.Println("That name is too short.")
            }
          }
          car, err = client.NewCar(ctx, &pb.NewCarRequest{Name: name})
          if err != nil {
            return nil, err
          }
        } else {
          return nil, err
        }
      }
      fmt.Printf("Found your car \"%v\".\n", car.Name)
    }
  }

  return car, nil
}

func main() {
  fmt.Println("Welcome to RoadTrip")
  setup()
  defer shutdown()

  char := getCharacter()
  car, err := getCar()

  if err != nil {
    log.Fatalf("Could not load car.")
  }

  fmt.Printf("Character: \"%v\"\n", char.Name)
  fmt.Printf("Car: \"%v\"\n", car.Name)

  /*
    var car *pb.Car
    car, err = client.GetCar(ctx, &pb.Empty{})
    st = status.Convert(err)
    if st != nil {
      if st.Code() == codes.NotFound {
        fmt.Println("Car does not exist, creating a car")
        car, err = client.NewCar(ctx, &pb.NewCarRequest{Name: "My New Car"})
        st = status.Convert(err)
        if st != nil {
          fmt.Printf("Could not create car\n")
        } else {
          fmt.Printf("Car created: #{car}\n")
        }
      }
    }
    fmt.Printf("Car response: %v\n", car)
  */
}
