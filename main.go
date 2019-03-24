package main

import "fmt"
import "math/rand"
import "time"

//CAR ... no use using real enums here - a string will do
const CAR = "car"

func main(){

    switchChoice := 0
    keepChoice := 0
    totalRuns := 0
    for {
        //setup: seed the rand, set up the choice and car position
        rand.Seed(time.Now().UnixNano())
        fmt.Println("Starting Monty Hall Problem Simulation")
        doors := []string{"goat", "goat", "goat"}
        carPosition := rand.Intn(3)
        doors[carPosition] = CAR
        choice := rand.Intn(3)
        var openDoor int
        var finalChoice string
        for {
            openDoor = rand.Intn(3)
            if (openDoor) != choice {
                break
            }
        }
        //keep choice: very simple, just check if we picked the car the first time
        if doors[choice] == CAR {
            keepChoice++
        }
        fmt.Printf("keep: %+v\n", doors[choice])
        //switch choice - could use a rework, but was useful while testing
        doors[choice] = ""
        doors[openDoor] = ""
        for _, v := range(doors) {
            if v != "" {
                finalChoice = v
            }
        }
        fmt.Printf("switch: %+v\n", finalChoice)
        if finalChoice == CAR {
            switchChoice++
        }
        //print values
        totalRuns++
        fmt.Printf("Success percentage while switching: %d / %d :  %f\n", switchChoice, totalRuns, float32(switchChoice)/float32(totalRuns))
        fmt.Printf("Success percentage while keeping: %d / %d :  %f\n", keepChoice, totalRuns, float32(keepChoice)/float32(totalRuns))
        time.Sleep(100 * time.Millisecond)
    }
}
