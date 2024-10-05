package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/pflag"
)

const (
	AppName        = "crash-and-burn"
	AppVersion     = "1.0.0"
	AppDescription = "A simple utility for randomly generating error, warning, and success return codes"
)

var (
	debug     bool
	errPct    int
	errRc     int
	warnPct   int
	warnRc    int
	sleepTime = rand.Int() % 11
	sleepDur  time.Duration
)

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
	if count == 1 {
		return ""
	}
	return "s"
}

// validatePercentage validates the percentage value
func validatePercentage(pct int) {
	if pct < 0 || pct > 100 {
		fmt.Println("Error: Percentage must be between 0 and 100")
		os.Exit(errRc)
	}
}

// validateReturnCode validates the return code value
func validateReturnCode(rc int) {
	if rc < 0 || rc > 255 {
		fmt.Println("Error: Return code must be between 0 and 255")
		os.Exit(99)
	}
}

// validateSleep validates the sleep time value
func validateSleep(sleep int) {
	if sleep < 0 {
		fmt.Println("Error: Sleep time must be greater than or equal to 0")
		os.Exit(errRc)
	}
}

func init() {
	// Setup flags
	pflag.IntVar(&errPct, "err-pct", 2, "set the error percentage (between 0-100)")
	pflag.IntVar(&errRc, "err-rc", 99, "set the error return code (between 0-255)")
	pflag.IntVar(&warnPct, "warn-pct", 30, "set the warning percentage (between 0-100)")
	pflag.IntVar(&warnRc, "warn-rc", 1, "set the warning return code (between 0-255)")
	pflag.IntVarP(&sleepTime, "sleep", "s", 0, "set the sleep time in seconds (must be greater or equal to 0) (default: random value between 0-10 seconds)")
	pflag.BoolVarP(&debug, "debug", "d", false, "enable debug mode")
	pflag.BoolP("help", "h", false, fmt.Sprintf("help for %s", AppName))
	pflag.BoolP("version", "v", false, fmt.Sprintf("version of %s", AppName))
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

	// Validate flags
	validatePercentage(errPct)
	validateReturnCode(errRc)
	validatePercentage(warnPct)
	validateReturnCode(warnRc)
	validateSleep(sleepTime)

	sleepDur = time.Duration(sleepTime) * time.Second
}

func main() {
	fmt.Println("Welcome to", AppName)
	fmt.Println("Version:", AppVersion)
	fmt.Println("")
	fmt.Println("Using the following settings:")
	fmt.Println("    - Error percentage:", errPct)
	fmt.Println("    - Error return code:", errRc)
	fmt.Println("    - Warning percentage:", warnPct)
	fmt.Println("    - Warning return code:", warnRc)
	fmt.Println("    - Sleep time:", sleepDur.String())

	randErrNbrs := genRandomNumbers(errPct)
	randWarnNbrs := genRandomNumbers(warnPct)

	if debug {
		fmt.Println("")

		for i := 0; i < len(randErrNbrs); i++ {
			fmt.Println("DEBUG: Error number:", randErrNbrs[i])
		}

		fmt.Println("")

		for i := 0; i < len(randWarnNbrs); i++ {
			fmt.Println("DEBUG: Warning number:", randWarnNbrs[i])
		}
	}

	randVal := rand.Int() % 100

	fmt.Println("")
	fmt.Printf("Sleeping for %d second%s...\n", sleepTime, pluralize(sleepTime))
	time.Sleep(sleepDur)
	fmt.Println("")

	if contains(randErrNbrs, randVal) {
		fmt.Println("ERROR with return code of", errRc)

		os.Exit(errRc)
	}

	if contains(randWarnNbrs, randVal) {
		fmt.Println("WARNING with return code of", warnRc)

		os.Exit(warnRc)
	}

	fmt.Println("SUCCESS with return code of 0")
}
