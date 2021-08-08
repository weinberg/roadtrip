playerStates = {}

const DMV = (token, state, input = "") => {
  playerState = playerStates[token]
  if (state === 0) {
    console.log("Hello... you are at the DMV!\n");
    if (playerState.hasCar) {
      console.log("You have a car");
    }
    console.log("Please pick an option:");
    console.log("A: Talk to Donna");
    console.log("B: Talk to Security Guard");
    console.log("C: Leave");
    return {next: 1, inputRequired: true};
  }

  if (state === 1) {
    if (input === "a") {
      console.log("You go over to talk to donna\n")
      console.log("Please pick an option:");
      if (playerState.hasCar) {
        console.log("You already have a car so you can't register another one")
      } else {
        console.log("A: Register a car");
      }
      console.log("B: Chit chat");
      console.log("C: Go back to front door");
      return {next: 2, inputRequired: true};
    }
  }

  if (state === 2) {
    if (input === "a") {
      console.log("You register a new car and go back to the front");
      playerState.hasCar = true
      return DMV(token, 0)
    } else if (input === "b") {
      console.log("You chit chat with donna. She hates you!");
      return DMV(token, 1, "a")
    } else if (input === "c") {
      console.log("You return to the front");
      return DMV(token, 0)
    }
  }

  return 1
};

token = "xyz"
playerStates[token] = { hasCar: false }

let next = 0;
let input;
({next, inputRequired} = DMV(token, next));

input = "a";
({next, inputRequired} = DMV(token, next, input));

input = "a";
({next, inputRequired} = DMV(token, next, input));

input = "a";
({next, inputRequired} = DMV(token, next, input));
