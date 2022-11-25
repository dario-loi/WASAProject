package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// Struct for a train stop
type Stop struct {
	Station  string `json:"name"`
	Distance int    `json:"distance"`
	ETA      int    `json:"ETA-from-previous-station"`
	Platform string `json:"platform"`
}

// Struct for a schedule
type Schedule struct {
	Line      string    `json:"line"`
	Stops     []Stop    `json:"stations"`
	Departure time.Time `json:"departingTime"`
}

// Struct for the output of the calculations
type Calculations struct {
	TotalDistance int       `json:"total-distance"`
	TotalETA      int       `json:"total-ETA"`
	ArrivalTime   time.Time `json:"arrival-time"`
}

// Method to return the representation of a stop
func (s *Stop) String() string {
	return fmt.Sprintf("Station: %s\nDistance %d\nETA: %d\nPlatform: %s\n", s.Station, s.Distance, s.ETA, s.Platform)
}

func main() {

	// Initialize the logger
	fileOut, err1 := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err1 != nil {
		log.Fatal("Logger Initialization Failed. ", err1)
	}
	log.SetOutput(fileOut)

	// Log cwd
	log.Println("Working Directory: ", func() string {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		return dir
	}())

	log.Println("Reading from the schedule file...")
	rawJSON, err3 := os.ReadFile("exercises\\trains1\\in.json")

	if err3 != nil {
		log.Fatal("Failed to read the schedule file. ", err3)
	}

	// Create a new schedule
	var schedule Schedule
	err4 := json.Unmarshal(rawJSON, &schedule)
	if err4 != nil {
		log.Fatal("Error reading trains.txt.\n", err4)
	}

	log.Println("Schedule read successfully, Printing Schedule to stdout and calculating operations.")
	fmt.Println("Line: ", schedule.Line)

	totDist := 0
	totETA := 0

	for _, v := range schedule.Stops {

		fmt.Println(v.String())

		totDist += v.Distance
		totETA += v.ETA

	}

	arrivalTime := schedule.Departure.Add(time.Duration(totETA) * time.Minute)

	fmt.Println("Departure: ", schedule.Departure)
	log.Println("Done Printing.")

	log.Println("Printing Results")

	fmt.Println("Total Distance: ", totDist)
	fmt.Println("Total ETA: ", totETA)
	fmt.Println("Arrival: ", arrivalTime)

	log.Println("Done Printing Results")
	log.Println("Writing to out.json")

	func() {
		os.WriteFile("exercises\\trains1\\out.json", func() []byte {
			data, err := json.MarshalIndent(Calculations{totDist, totETA, arrivalTime}, "", "\t")
			if err != nil {
				log.Fatal("Error writing to out.json.\n", err)
			}
			return data
		}(), 0666)
	}()

	log.Println("Output complete.")
	log.Println("Done.")
}
