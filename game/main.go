package game

import (
	"math/rand"
	"strconv"
	"strings"
)

var randomFunc func() int

func Start(outputHandler func(interface{}), inputChannel chan string, exitGame chan struct{}) {

	
	randomFunc = func() int {
		return rand.Intn(10)
	}

	A := randomFunc()
	T := 0

	outputHandler(strings.Repeat(" ", 24) + "COUNT DOWN\n")
	outputHandler(strings.Repeat(" ", 20) + "CREATIVE COMPUTING\n")
	outputHandler(strings.Repeat(" ", 18) + "MORRISTOWN, NEW JERSEY\n\n\n\n")
	outputHandler("YOU HAVE ACTIVATED THE SELF-DESTRUCT MECHANISM IN THIS SCHOOL.\n")
	outputHandler("IF YOU WISH, YOU MAY STOP THE MECHANISM.\n")
	outputHandler("TO DO SO, JUST TYPE IN THE CORRECT NUMBER,\n")
	outputHandler("WHICH WILL STOP THE COUNT-DOWN.\n")
	outputHandler("PLEASE HURRY!! THERE IS NO TIME TO WASTE!!!!!!!\n")

	for {
		outputHandler("WHAT'LL IT BE? ")
		X, err := strconv.Atoi(<-inputChannel)
		if err != nil {
			outputHandler("!NUMBER EXPECTED - RETRY INPUT LINE\n? ")
			continue
		}
		if T == 4 {
			outputHandler("\n\n\n\n")
			outputHandler(strings.Repeat(" ", 32) + "TOO LATE\n\n\n\n\n")
			outputHandler(strings.Repeat(" ", 32) + "\\ **** /\n")
			outputHandler(strings.Repeat(" ", 31) + "-- BOOM --\n")
			outputHandler(strings.Repeat(" ", 32) + "/ **** \\\n\n\n\n")
			break
		}
		outputHandler("\n")
		if X < A {
			outputHandler("TOO SMALL!!!!!\n")
		} else if X > A {
			outputHandler("TOO BIG!!!!!\n")
		} else if X == A {
			outputHandler("CORRECT!!!!\n")
			outputHandler("THE COUNTDOUN HAS STOPPED.\n")
			outputHandler("YOU HAVED SAVED THE SCHOOL!\n")
			outputHandler("(HAVE YOU SEEN YOUR SHRINK LATELY ?)\n")
			break
		}
		outputHandler("YOUR NUMBER DOES NOT COMPUTE!!\n")
		outputHandler("PLEASE TRY AGAIN!!!!\n")
		T += 1
		if T == 2 {
			outputHandler("TIME GROWS SHORT, PLEASE HURRY!!!!!!!!\n")
		} else if T == 3 {
			outputHandler("HURRY, THE COUNT-DOWN IS APPROACHING ZERO!!!!!!!!!\n")
		}
	}
	close(exitGame)
}
