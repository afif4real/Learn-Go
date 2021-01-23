package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
  fmt.Println("you have", fuel, "gallons of fuel left")
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
  var fuel int
  switch planet {
  case "Venus":
    fuel = 300000
  case "Mercury":
    fuel = 5000000
  case "Mars":
    fuel = 7000000
  default:
    fuel = 0
  }

  return fuel
}

// Create the function greetPlanet() here
func greetPlanet (planet string) {
  fmt.Println("welcome to planet", planet)
}

// Create the function cantFly() here
func cantFly () {
  fmt.Println("we do not have the availble fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet (planet string, fuel int) int {
  var fuelRemaining, fuelCost int
  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)

  if fuelRemaining >= fuelCost {
    greetPlanet(planet)
    fuelRemaining -= fuelCost
  } else {
    cantFly()
  }

  return fuelRemaining
}
 
func main() {
  // Test your functions!

  // Create `planetChoice` and `fuel`
  var fuel int
  fuel = 1000000

  var planetChoice string
  planetChoice = "Venus"
  // And then liftoff!
  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)
}