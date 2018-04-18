package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// DivideCommand is an instruction processed by DivideWorker
type DivideCommand struct {
	dividend int
	divisor  int
}

// DivideWorker transforms DivideCommand's into integer outputs
func DivideWorker(inbox chan DivideCommand) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Caught err: {%v}\n", r)
			go DivideWorker(inbox) // spawn new worker (instead of recursing)
		}
	}()

	log.Printf("New worker spawned\n")

	for cmd := range inbox {
		log.Printf("Dividing %d by %d yields %d\n", cmd.dividend, cmd.divisor, cmd.dividend/cmd.divisor)
	}
}

func main() {
	workBox := make(chan DivideCommand)

	for idx := 0; idx < 3; idx++ {
		go DivideWorker(workBox)
	}

	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Welcome to the DIVIDER 9000!\n")

	// Input processing loop
	for reader.Scan() {
		inStr := reader.Text()
		if strings.Contains(inStr, "/") {
			divIdx := strings.Index(inStr, "/")
			dividend, err := strconv.Atoi(inStr[:divIdx])
			if err != nil {
				fmt.Print("Error: Malformed dividend.\n")
				continue
			}

			divisor, err := strconv.Atoi(inStr[divIdx+1:])
			if err != nil {
				fmt.Print("Error: Malformed divisor.\n")
				continue
			}

			workBox <- DivideCommand{
				dividend: dividend,
				divisor:  divisor,
			}
			continue
		}
		fmt.Print("Error: Malformed input.\n")
	}
}
