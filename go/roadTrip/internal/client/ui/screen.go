package ui

import (
	"fmt"
	player_grpc "github.com/brickshot/roadtrip/internal/playerServer/grpc"
	"math"
	"time"
)

var esc = "\033"

type Screen struct {
	Width  int
	Height int
}

func cls() {
	fmt.Print("\033[2J")
}

func home() {
	fmt.Printf("%s[H", esc)
}

func moveToColumn(c int) {
	fmt.Printf("%s[%dG", esc, c)
}

func moveTo(x int, y int) {
	fmt.Printf("%s[%d;%dH", esc, y, x)
}

func color(fg int, bg int) {
	fmt.Printf("%s[%d;%dm", esc, fg, bg)
}

func (s *Screen) statusLine(u *player_grpc.Update) {
	odo := fmt.Sprintf("Odometer: %08.1f m", u.Character.Car.OdometerMiles)
	home()
	fmt.Printf("%s", u.RoadName)
	moveToColumn(s.Width - len(odo) + 1)
	fmt.Printf("%s", odo)
}

func (s *Screen) canvas(u *player_grpc.Update, trip *player_grpc.Trip) {
	for i := 2; i < 24; i++ {
		moveTo(1, i)
		fmt.Printf("                                                                                ")
	}
	var x = 5
	var y = 5
	var fc rune
	for i, e := range trip.Entries {
		moveTo(x, y)
		if e.Type == "town" {
			if u.Character.Car.Location.TownId == e.Town.Id {
				fc = '*'
			} else {
				if i == 0 {
					fc = '╔'
				} else if i == len(trip.Entries)-1 {
					fc = '╚'
				} else {
					fc = '╠'
				}
			}
			fmt.Printf("%c═ ", fc)
			fmt.Printf("%v", e.Town.DisplayName)
			y++
		} else if e.Type == "road" {
			moveTo(x, y)
			if u.Character.Car.Location.RoadId == e.Road.Id {
				fc = '*'
			} else {
				fc = '║'
			}
			fmt.Printf("%c   %s", fc, e.Road.DisplayName)
			y++
		}
	}
}

func (s *Screen) showTime() {
	date := fmt.Sprintf("%s", time.Now().Format(time.RFC850))
	moveTo(s.Width-len(date)+1, 25)
	fmt.Printf(date)
}

func (s *Screen) footer(u *player_grpc.Update) {
	moveTo(1, 24)
	nameStr := fmt.Sprintf("[%-18s]                                         Plate:%s %dmph",
		u.Character.CharacterName, u.Character.Car.Plate, int(math.Floor(float64(u.Character.Car.VelocityMph))))
	fmt.Printf("%s", nameStr)
	s.showTime()
}

func (s *Screen) time(done chan bool) {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-done:
			return
		case _ = <-ticker.C:
			s.showTime()
		}
	}
}

func (s *Screen) RenderUpdate(updates <-chan *player_grpc.Update, trip *player_grpc.Trip) {
	done := make(chan bool)
	go s.time(done)
	for {
		update, more := <-updates
		if !more {
			break
		}
		cls()
		s.statusLine(update)
		s.canvas(update, trip)
		moveTo(4, 2)
		s.footer(update)
	}
	done <- true
	/*
	   fmt.Printf("Character: %v\n", character.CharacterName)
	   fmt.Printf("Car Plate: %v\n", character.Car.Plate)
	   fmt.Printf("Location:\n")
	   fmt.Printf("  Town : %v\n", character.Car.Location.TownId)
	   fmt.Printf("  Road : %v\n", character.Car.Location.RoadId)
	   fmt.Printf("  Position : %v\n", character.Car.Location.PositionMiles)
	   fmt.Println()
	*/

	/*
	   town, err := client.GetTown(context.Background(), &psgrpc.GetTownRequest{Id: character.Car.Location.TownId})
	   if err != nil {
	     log.Fatalf("Cannot find town details: %v\n", err)
	   }
	   fmt.Printf("Town Details:\n")
	   fmt.Printf("  State: %v\n", town.StateId)
	   fmt.Printf("  Town : %v\n", town.DisplayName)
	   fmt.Printf("  Info : %v\n", town.Description)
	*/
}
