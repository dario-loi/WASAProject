package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	NumberStops   int       `json:"number-stops"`
}

// Method to return the representation of a stop
func (s *Stop) String() string {
	return fmt.Sprintf("Station: %s\nDistance %d\nETA: %d\nPlatform: %s\n", s.Station, s.Distance, s.ETA, s.Platform)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	log.Println("Handling http request from ", r.RemoteAddr)

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal("Failed to read the request body. ", err)
	}

	var schedule Schedule
	if err := json.Unmarshal(body, &schedule); err != nil {
		log.Fatal("Failed to unmarshal the request body. ", err)
	}

	log.Println("Schedule read successfully.")
	fmt.Println(fmt.Sprintf("Schedule %+v\n", schedule))

	totDist := 0
	totETA := 0

	log.Println("Performing calculations...")
	for _, v := range schedule.Stops {
		totDist += v.Distance
		totETA += v.ETA
	}

	arrivalTime := schedule.Departure.Add(time.Duration(totETA) * time.Minute)

	fmt.Println("Departure: ", schedule.Departure)

	log.Println("Printing Results...")

	fmt.Println("Total Distance: ", totDist)
	fmt.Println("Total ETA: ", totETA)
	fmt.Println("Arrival: ", arrivalTime)

	log.Println("Composing .json response")

	outData, err := json.MarshalIndent(Calculations{totDist, totETA, arrivalTime, len(schedule.Stops)}, "", "\t")
	if err != nil {
		log.Fatal("Failed to marshal the calculations into the response. ", err)
	}

	log.Println("Sending response to client")
	i, err := w.Write(outData)
	if err != nil {
		log.Fatal("Failed to write the response to the client. ", err)
	}

	log.Println("Response sent successfully. ", i, " bytes written.")
}

func main() {

	// Initialize the logger
	fileOut, err1 := os.OpenFile("trains.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err1 != nil {
		log.Fatal("Logger Initialization Failed. ", err1)
	}
	log.SetOutput(fileOut)
	defer func() {
		log.Println("Terminating.")
	}()

	// Log cwd
	log.Println("Working Directory: ", func() string {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		return dir
	}())

	log.Println("Starting server on port 8080")

	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)

}
