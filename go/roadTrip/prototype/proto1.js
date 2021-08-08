// this proto is like how you would write it local not client/server

const p = console.log

function DMV() {
  var hasCar = false
  p("Welcome to DMV")
  p("Please pick an option:");
  p("A: Talk to Donna");
  p("B: Talk to Security Guard");
  p("C: Leave");

  const input = "a"

  if (input === "a") {
    var talkingToDonna = true
    while (talkingToDonna) {
      p("You are talking to Donna");
      p("Please pick an option:");
      if (hasCar) {
        p("You already have a car so you can't register another one")
      } else {
        p("A: Register a car");
      }
      p("B: Chit chat");
      p("C: Go back to front door");
    }
  }
}

DMV()