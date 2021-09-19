package main

import (
  "bufio"
  "context"
  "crypto/tls"
  "crypto/x509"
  "flag"
  "fmt"
  "github.com/brickshot/roadtrip/internal/client/config"
  "github.com/brickshot/roadtrip/internal/client/ui"
  psgrpc "github.com/brickshot/roadtrip/internal/playerServer/grpc"
  "google.golang.org/grpc"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/credentials"
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
var tlsDisabled bool

func init() {
  flag.BoolVar(&tlsDisabled, "tls", false, "If false, use TLS. Defaults to false.")
}

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

func GetCACertPEM() []byte {
  return []byte(`
-----BEGIN CERTIFICATE-----
MIIFSTCCAzGgAwIBAgIUHJgFIt/AsBVn0fNpu8wb6HBKWvkwDQYJKoZIhvcNAQEL
BQAwNDERMA8GA1UECgwIUm9hZFRyaXAxHzAdBgNVBAMMFioucm9hZHRyaXAuaW5z
b2Zhci5jb20wHhcNMjEwOTE5MTcyNjU5WhcNMjIwOTE5MTcyNjU5WjA0MREwDwYD
VQQKDAhSb2FkVHJpcDEfMB0GA1UEAwwWKi5yb2FkdHJpcC5pbnNvZmFyLmNvbTCC
AiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAOCO/K31SLgVefkvbuTEoA6C
Iy8kqMV1RqydItl0axlyaLpdpJDgm1DaIuvVoM4jhcqebt69dtogQYpnP4yIiRaF
keG3xSlbpALHHbyqQDyaHAFoMbjrsHSUxsxyq71GkMIov2Ks+DcYwT+3sB5OvfZL
gGAn3weweifqC0HSntpGfKGmF0af93U8OAyn21Mc66hrOwpebNyd4DGqm08RcVtj
ptdBnRh6jAAfH7YuMyZqRSmGAMGjhC/xkzQWUiTqyQwoaTI0P+HH9gN3pUrjAn9F
6uedyYFA2JoWoudasZ5aZTUabyurkHBp6F+EKFdlr+VMgu5pxrHUNmxEywptWPm5
NjI+J0iddhSAmLwJQ+IUVxGrBFMiS2aZZ8561i1CZ3AYCLg/TE4b3dB2BgJGQE5Y
5Q5eBoGxMW97IAlSvlG/RXmED0ptr85T2tiqy1SEZ/Gj5Xzz9QIiIifyhpFc/DUr
54pJg+KnrMC1rat5Wezg5BeMGdV3LHf2TA3d0M36jKjOSMkCbLcfuOGpek9XkjNE
/hJBeedK/qfl/P4xnCoFCARvhSz91wrQOrbQVo+oHlPf2+dZr3RVwjZzwY8sdy4q
CaRwkAzzw7U0USWEN7kPNc5NudpraRBuXVfvlKP47U4Ob3Tg7wDxt2vIZTUvn+IL
q4ELmyyuwnsaetYrc+adAgMBAAGjUzBRMB0GA1UdDgQWBBTeWcqVlYbCMB5d27xc
YIxTLyIccTAfBgNVHSMEGDAWgBTeWcqVlYbCMB5d27xcYIxTLyIccTAPBgNVHRMB
Af8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4ICAQC5b9MmONA5CWKOpJm/Y0zZLylB
aQrd0yoOXTv81GI8JPtuPiq0q3iMqpCeRVAvJ4LhSu9+ugu5RYBB/DcjFIxfnYjk
m994EjwHpqQYLhgpr28/COb+kfN2DafimB4WZfEzJvIPYhWe4qy1Kp2MY+vlfffd
bsGn7f6qn9FlOn+v8Gx8M+SJjRIaHEzvpl/VV6tzD7ZrM7zMcAzQhJXZD5pudL8M
PzbTpZt0Rf7IIoKvHQufXBmFsS3WCXv0xO/kMbqOkvPvOLR3FYGTItv0HaiPofMi
ImRNP5NlnwMLkaXWm5bK9D5eCnGgdDVmN2ZUI3d+cn9zBu4EGkLJCGqx3Ac94jg9
81E2Ppj1fDmptNa8P2NSfAZkz6Ya5nxfjPE/3R6COk0+fIH2gbysTTAkRvl5sQVb
ey3kMwqAfowRoyUpd62p4iDF0urUufsKPCcPzO7M9NyEPu3qWhpQ+XrJ0Z3N7obP
n/vub4inipW7aCPhWxUv6nwEQp9G+jehkMRtS3D09tOhybqC4y3iAe8r5x9sIii0
AE+NZWOZRoxNIHbPutI1weCijJ/TMH6octnCfeaFOYSaT35yK7TrJ035ExZQkB3R
4mUBZzOrkSwaDphlh8PGLmnVj8VZes7JDojeKX+BfiKa2/lC7rJF2csR36RcukFy
H4XD+yHlKm6cgJ/lvQ==
-----END CERTIFICATE-----`)
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
  caCert := GetCACertPEM()
  certPool := x509.NewCertPool()
  if !certPool.AppendCertsFromPEM(caCert) {
    return nil, fmt.Errorf("failed to add server CA's certificate")
  }

  // Create the credentials and return it
  config := &tls.Config{
    RootCAs: certPool,
  }

  return credentials.NewTLS(config), nil
}

func setupGrpc(conf config.ClientConfig) {
  tlsCredentials, err := loadTLSCredentials()
  if err != nil {
  	log.Fatalln("Cannot load TLS credentials to make connection: ", err)
	}
  opts := grpc.WithTransportCredentials(tlsCredentials)

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
    log.Fatalf("Failed to create character: %v.\n", err)
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
