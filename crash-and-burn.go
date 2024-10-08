package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/pflag"

	"crash-and-burn/settings"
)

const (
	AppName        = "crash-and-burn"
	AppVersion     = "1.0.0"
	AppDescription = "A simple utility for randomly generating success and failure return codes"
)

var (
	allPctMap = make(map[int]int)

	// Flags
	verbose   bool
	failRCs   settings.FailureRCs
	sleepTime int
	sleepDur  time.Duration
)

// genRandomNumbers generate a slice of random numbers based on the percentage
func genRandomNumbers(pct int) []int {
	var randNbrs []int

	i := 0

	for i < pct {
		randNbr := rand.Int() % 101

		if !contains(randNbrs, randNbr) {
			if _, exists := allPctMap[randNbr]; exists {
				delete(allPctMap, randNbr)
				randNbrs = append(randNbrs, randNbr)
				i++
			}
		}
	}

	return randNbrs
}

// contains check if a slice contains a value
func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// pluralize returns "s" if the count is not 1, otherwise returns an empty string.
func pluralize(count int) string {
	if count == 1 || count == -1 {
		return ""
	}
	return "s"
}

func init() {
	es := settings.FailureRCs{}
	es.Set("2,2")

	// Setup flags
	pflag.VarP(&failRCs, "set-fail", "f", "set the percentage of a specified failure return code, The format is rc,percentage. Can be set multiple times. Return codes must be between 1 and 255 and percentages must be between 1 and 100.")
	pflag.IntVarP(&sleepTime, "sleep", "s", 0, "set the sleep time in seconds (must be greater or equal to 0) (default: random value between 0-10 seconds)")
	pflag.BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	pflag.BoolP("help", "h", false, fmt.Sprintf("help for %s", AppName))
	pflag.Bool("version", false, fmt.Sprintf("version of %s", AppName))
	pflag.CommandLine.SortFlags = false

	pflag.Parse()

	// Handle flags
	if pflag.Lookup("help").Changed {
		fmt.Println(AppDescription)
		fmt.Println("")
		fmt.Printf("Usage: %s [flags]\n", os.Args[0])
		fmt.Println("")
		fmt.Println("Flags:")
		pflag.PrintDefaults()
		os.Exit(0)
	}

	if pflag.Lookup("version").Changed {
		fmt.Printf("%s version %s\n", AppName, AppVersion)
		os.Exit(0)
	}

	// Generate a random sleep time if not set
	if !pflag.Lookup("sleep").Changed {
		sleepTime = rand.Int() % 11
	}

	if sleepTime < 0 {
		fmt.Printf("sleep time (%d) must be greater than or equal to 0\n", sleepTime)
		fmt.Printf("Usage: %s [flags]\n", os.Args[0])
		fmt.Println("")
		fmt.Println("Flags:")
		pflag.PrintDefaults()
		fmt.Printf("sleep time (%d) must be greater than or equal to 0\n", sleepTime)
		os.Exit(2)
	}

	sleepDur = time.Duration(sleepTime) * time.Second

	for i := 1; i <= 100; i++ {
		allPctMap[i] = i
	}
}

func main() {
	if verbose {
		fmt.Printf("%s, v%s\n", AppName, AppVersion)
		fmt.Println("")
		fmt.Println("Sleep time will be:", sleepDur.String())
		fmt.Println("Return Code Settings:")

		if failRCs.TotalPct() < 100 {
			succPct := 100 - failRCs.TotalPct()
			fmt.Printf("    - RC: 0 (%d%%) [SUCCESS]\n", succPct)
		}

		for i := 0; i < len(failRCs); i++ {
			fmt.Printf("    - RC: %d (%d%%) [FAILURE]\n", failRCs[i].RC, failRCs[i].Pct)
		}
	}

	// Assign random values to the failure RCs based on percentage
	for i := 0; i < len(failRCs); i++ {
		failRCs[i].RandValues = genRandomNumbers(failRCs[i].Pct)
	}

	if verbose {
		fmt.Println("")
		fmt.Printf("Sleeping for %d second%s...\n", sleepTime, pluralize(sleepTime))
	}

	time.Sleep(sleepDur)

	if verbose {
		fmt.Println("")
	}

	randVal := rand.Int() % 100

	// Check if the random value is in the failure RCs
	for i := 0; i < len(failRCs); i++ {
		if contains(failRCs[i].RandValues, randVal) {
			if verbose {
				fmt.Println("FAILURE with return code of", failRCs[i].RC)
			}

			os.Exit(failRCs[i].RC)
		}
	}

	if verbose {
		fmt.Println("SUCCESS with return code of 0")
	}
}
