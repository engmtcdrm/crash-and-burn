package settings

import (
	"fmt"
	"regexp"
)

type FailureRC struct {
	RC         int
	Pct        int
	RandValues []int
}

type FailureRCs []FailureRC

// String returns the string representation of the FailureRC type
func (e *FailureRCs) String() string {
	// return fmt.Sprintf("%v", *e)
	return ""
}

// Set parses the input string and sets the value of the FailureRC type
func (e *FailureRCs) Set(value string) error {
	var pct, rc int

	validInput := regexp.MustCompile(`^\d+,\d+$`)

	if !validInput.MatchString(value) {
		return fmt.Errorf("invalid format: input must be two integers separated by a comma")
	}

	_, err := fmt.Sscanf(value, "%d,%d", &rc, &pct)
	if err != nil {
		return fmt.Errorf("invalid format: %v", err)
	}

	if rc < 1 || rc > 255 {
		return fmt.Errorf("return code (%d) must be between 1 and 255", rc)
	}

	if pct < 1 || pct > 100 {
		return fmt.Errorf("percentage (%d) must be between 1 and 100", pct)
	}

	if !e.exists(rc) {
		if e.TotalPct()+pct <= 100 {
			*e = append(*e, FailureRC{RC: rc, Pct: pct})
		} else {
			fmt.Printf("Total failure percentage is at or over 100. Return code (%d) and percentage (%d) will not be added.\n", rc, pct)
		}
	} else {
		fmt.Printf("Return code (%d) already exists and will not be added again.\n", rc)
	}

	return nil
}

// Type returns the type of the FailureRC type
func (e *FailureRCs) Type() string {
	return "int,int"
}

// TotalPct returns the total percentage of all the FailureRC types
func (e *FailureRCs) TotalPct() int {
	pctTotal := 0

	for _, es := range *e {
		pctTotal += es.Pct
	}

	return pctTotal
}

// exists checks if a return code already exists in the FailureRC type
func (e *FailureRCs) exists(rc int) bool {
	for _, es := range *e {
		if es.RC == rc {
			return true
		}
	}

	return false
}
