package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  var isHeistOn = true
  var eludedGuards = rand.Intn(100)

  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
  }
  var openedVault = rand.Intn(100)
  if openedVault >= 70 {
    fmt.Println("Grab and GO!")
  } else if isHeistOn {
    fmt.Println("VAULT COULDN'T BE OPENED!")
  }
  var leftSafely = rand.Intn(5)
  if isHeistOn {
      switch leftSafely {
        case 0:
          fmt.Println("FAILED!")
          isHeistOn = false
        case 1:
          fmt.Println("FAILED!")
          isHeistOn = false
        case 2:
          fmt.Println("FAILED!")
          isHeistOn = false
        case 3:
          fmt.Println("FAILED!")
          isHeistOn = false
        case 4:
          fmt.Println("SUCCESFUL!")
        default:
          fmt.Println("SUCCESFUL!")
      } 
  }
  if amtstolen := 10000 + rand.Intn(1000000); isHeistOn {
    fmt.Println("You stole:", amtstolen)
  }
}
