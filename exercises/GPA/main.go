package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	//logger init

	func() {
		log.Println("Initializing Logger")
		log.SetFlags((log.Ldate | log.Ltime | log.Lshortfile))
		fileOut, err1 := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0666)
		if err1 != nil {
			log.Fatal("Logger Initialization Failed. ", err1)
		}
		log.SetOutput(fileOut)
	}()

	log.Println("Awaiting input...")
	i := 0
	in_score := 0
	in_credits := 0
	running_products := 0.0
	credits := 0

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a set of scores, -1 when done, you can only stop when entering scores, not credits.")
READER:
	for {

		in_score = func() int {
			log.Println(fmt.Sprintf("Reading input %d", i))
			i += 1
			fmt.Println("Insert a score")
			in, is_prefix, err := reader.ReadLine()
			if is_prefix {
				fmt.Println("Input too long. Please try again.")
				return -2
			}
			if err != nil {
				log.Fatal("Failed to read input. ", err)
			}

			ret, err1 := strconv.Atoi(string(in))

			if err1 != nil {
				fmt.Println("Invalid input. Please try again.")
				return -2
			}

			if ret == 31 {
				ret = 30
			} else if ret > 31 {
				fmt.Println("Input must be between 1 and 31, please try again.")
				return -2
			}

			return ret

		}()
		switch in_score {
		case -1:
			log.Println("Stopped Accepting User Input.")
			break READER
		case -2:
			continue
		}

		in_credits = func() int {
			log.Println(fmt.Sprintf("Reading input %d", i))
			i += 1
			fmt.Println("Insert a credit")
			in, is_prefix, err := reader.ReadLine()
			if is_prefix {
				fmt.Println("Input too long. Please try again.")
				return -2
			}
			if err != nil {
				log.Fatal("Failed to read input. ", err)
			}

			ret, err1 := strconv.Atoi(string(in))

			if err1 != nil {
				fmt.Println("Invalid input. Please try again.")
				return -2
			}

			return ret

		}()
		switch in_credits {
		case -1:
			log.Println("Must insert a corresponding number of credits.")
			continue
		case -2:
			continue

		}

		credits += in_credits
		running_products += float64(in_score) * float64(in_credits)

	}

	log.Println("Calculating GPA...")

	GPA := func() float64 {
		return running_products / float64(credits)
	}()

	log.Println("GPA calculated.")
	fmt.Println("GPA: ", GPA)

	log.Println("Exiting...")
}
