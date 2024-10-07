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
	AppDescription = "A simple utility for randomly generating error, warning, and success return codes"
)

var (
	allPctMap = make(map[int]int)

	// Flags
	verbose     bool
	errSettings settings.ErrorSettings
	errPct      int
	errRc       int
	sleepTime   int
	sleepDur    time.Duration
)

func validateFlags() {
	validRC := 33

	if err := errSettings.Validate(); err != nil {
		fmt.Println(err)
		os.Exit(validRC)
	}

	if errPct < 0 || errPct > 100 {
		fmt.Printf("error percentage (%d) must be between 0 and 100\n", errPct)
		os.Exit(validRC)
	}

	if errRc < 0 || errRc > 255 {
		fmt.Printf("error return code (%d) must be between 0 and 255\n", errRc)
		os.Exit(validRC)
	}

	if sleepTime < 0 {
		fmt.Printf("sleep time (%d) must be greater than or equal to 0\n", sleepTime)
		os.Exit(validRC)
	}
}

// genRandomNumbers generate a slice of random numbers based on the percentage
func genRandomNumbers(pct int) []int {
	var randNbrs []int

	i := 0

	for i < pct {
		randNbr := rand.Int() % 100

		if !contains(randNbrs, randNbr) {
			randNbrs = append(randNbrs, randNbr)
			i++
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

func doSleep() {
	if verbose {
		fmt.Println("")
		fmt.Printf("Sleeping for %d second%s...\n", sleepTime, pluralize(sleepTime))
	}

	time.Sleep(sleepDur)

	if verbose {
		fmt.Println("")
	}
}

func init() {
	// Setup flags
	pflag.Var(&errSettings, "set-err", "set error return code and percentage in the format code,percentage. If no --set-err flag is provided, the default is 2,30")
	pflag.IntVarP(&errPct, "err-pct", "p", 2, "set the error percentage (between 0-100)")
	pflag.IntVarP(&errRc, "err-rc", "r", 99, "set the error return code (between 0-255)")
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

	if !pflag.Lookup("set-err").Changed {
		errSettings = append(errSettings, settings.ErrorSetting{RC: 30, Pct: 2})
	}

	validateFlags()

	sleepDur = time.Duration(sleepTime) * time.Second

	for i := 1; i <= 100; i++ {
		allPctMap[i] = i
	}
}

func main() {
	if verbose {
		fmt.Printf("Welcome to %s, v%s\n", AppName, AppVersion)
		fmt.Println("")
		fmt.Println("Using the following settings:")
		fmt.Println("    - Sleep time:", sleepDur.String())
		fmt.Println("    - Error Settings:")

		for i := 0; i < len(errSettings); i++ {
			fmt.Printf("        - RC: %d (%d%%)\n", errSettings[i].RC, errSettings[i].Pct)
		}
	}

	randErrNbrs := genRandomNumbers(errPct)

	if verbose {
		fmt.Println("")

		for i := 0; i < len(randErrNbrs); i++ {
			fmt.Println("Randomly Generated Error number:", randErrNbrs[i])
		}
	}

	randVal := rand.Int() % 100

	doSleep()

	if contains(randErrNbrs, randVal) {
		if verbose {
			fmt.Println("ERROR with return code of", errRc)
		}

		os.Exit(errRc)
	}

	if verbose {
		fmt.Println("SUCCESS with return code of 0")
	}
}
