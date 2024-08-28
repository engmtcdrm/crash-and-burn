package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Generate a slice of random numbers
//
// randomRange: the range of random numbers to generate
//
// pct: the percentage of random numbers to generate
func genRandomNumbers(randomRange int, pct int) []int {
	var randNbrs []int

	i := 0

	for i < pct {
		randNbr := rand.Int() % randomRange

		if !contains(randNbrs, randNbr) {
			randNbrs = append(randNbrs, randNbr)
			i++
		}
	}

	return randNbrs
}

// Check if a slice contains a value
//
// slice: the slice to check
//
// value: the value to check for
func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	randomRange := 100
	errPct := 2
	errRc := 99
	warnPct := 30
	warnRc := 1
	sleepTime := rand.Int() % 10
	sleepDur := time.Duration(sleepTime) * time.Second
	debug := false
	var err error

	args := os.Args

	// Parse command line arguments
	for i := 0; i < len(args); i++ {
		if args[i] == "-d" || args[i] == "--debug" {
			debug = true
		} else if args[i] == "-r" || args[i] == "--random-range" {
			if i+1 >= len(args) {
				println("Error: missing argument for random range")
				os.Exit(1)
			}

			randomRange, err = strconv.Atoi(args[i+1])

			if err != nil {
				println("Error: invalid argument for random range. Value must be an integer")
				os.Exit(1)
			}
		} else if args[i] == "-ep" || args[i] == "--error-pct" {
			if i+1 >= len(args) {
				println("Error: missing argument for error percentage")
				os.Exit(1)
			}

			errPct, err = strconv.Atoi(args[i+1])

			if err != nil {
				println("Error: invalid argument for error percentage. Value must be an integer")
				os.Exit(1)
			}
		} else if args[i] == "-er" || args[i] == "--error-rc" {
			if i+1 >= len(args) {
				println("Error: missing argument for error return code")
				os.Exit(1)
			}

			errRc, err = strconv.Atoi(args[i+1])

			if err != nil {
				println("Error: invalid argument for error return code. Value must be an integer")
				os.Exit(1)
			}
		} else if args[i] == "-wp" || args[i] == "--warn-pct" {
			if i+1 >= len(args) {
				println("Error: missing argument for warning percentage")
				os.Exit(1)
			}

			warnPct, err = strconv.Atoi(args[i+1])

			if err != nil {
				println("Error: invalid argument for warning percentage. Value must be an integer")
				os.Exit(1)
			}
		} else if args[i] == "-wr" || args[i] == "--warn-rc" {
			if i+1 >= len(args) {
				println("Error: missing argument for warning return code")
				os.Exit(1)
			}

			warnRc, err = strconv.Atoi(args[i+1])

			if err != nil {
				println("Error: invalid argument for warning return code. Value must be an integer")
				os.Exit(1)
			}
		} else if args[i] == "-s" || args[i] == "--sleep" {
			if i+1 >= len(args) {
				println("Error: missing argument for sleep time in seconds")
				os.Exit(1)
			}

			sleepTime, err = strconv.Atoi(args[i+1])

			if err != nil {
				println("Error: invalid argument for sleep time in seconds. Value must be an integer")
				os.Exit(1)
			}

			sleepDur = time.Duration(sleepTime) * time.Second
		}
	}

	if debug {
		println("Debug mode enabled")
		println("")
		println("Random range:", randomRange)
		println("Error percentage:", errPct)
		println("Error return code:", errRc)
		println("Warning percentage:", warnPct)
		println("Warning return code:", warnRc)
		println("Sleep time:", sleepTime)
		println("Sleep duration:", sleepDur.String())
	}

	randErrNbrs := genRandomNumbers(randomRange, errPct)
	randWarnNbrs := genRandomNumbers(randomRange, warnPct)

	if debug {
		for i := 0; i < len(randErrNbrs); i++ {
			println("Error number:", randErrNbrs[i])
		}

		for i := 0; i < len(randWarnNbrs); i++ {
			println("Warning number:", randWarnNbrs[i])
		}
	}

	randVal := rand.Int() % randomRange

	println("Sleeping for", sleepTime, "seconds")
	time.Sleep(sleepDur)

	if contains(randErrNbrs, randVal) {
		println("ERROR with return code of", errRc)

		os.Exit(errRc)
	}

	if contains(randWarnNbrs, randVal) {
		println("WARNING with return code of", warnRc)

		os.Exit(warnRc)
	}

	println("SUCCESS with return code of 0")
}
