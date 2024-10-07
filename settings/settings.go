package settings

import "fmt"

type ErrorSetting struct {
	RC         int
	Pct        int
	RandValues []int
}

type ErrorSettings []ErrorSetting

// String returns the string representation of the ErrorSettings type
func (e *ErrorSettings) String() string {
	return fmt.Sprintf("%v", *e)
}

// Set parses the input string and sets the value of the ErrorSettings type
func (e *ErrorSettings) Set(value string) error {
	var pct, rc int

	_, err := fmt.Sscanf(value, "%d,%d", &rc, &pct)
	if err != nil {
		return fmt.Errorf("invalid format: %v", err)
	}

	if !e.exists(rc) {
		if e.TotalPct() < 100 {
			*e = append(*e, ErrorSetting{RC: rc, Pct: pct})
		} else {
			fmt.Printf("Total error percentage is at or over 100. Return code (%v) will not be added.\n", rc)
		}
	} else {
		fmt.Printf("Return code (%d) already exists and will not be added again.\n", rc)
	}

	return nil
}

// Type returns the type of the ErrorSettings type
func (e *ErrorSettings) Type() string {
	return "ErrorSettings"
}

func (e *ErrorSettings) Validate() error {
	pctTotal := 0

	for _, es := range *e {
		if es.RC < 0 || es.RC > 255 {
			return fmt.Errorf("return code (%d) must be between 0 and 255", es.RC)
		}

		if es.Pct < 0 || es.Pct > 100 {
			return fmt.Errorf("percentage (%d) must be between 0 and 100", es.Pct)
		}

		pctTotal += es.Pct
	}

	// if pctTotal > 100 {
	// 	return fmt.Errorf("total percentage (%d) must be less than or equal to 100", pctTotal)
	// }

	return nil
}

func (e *ErrorSettings) TotalPct() int {
	pctTotal := 0

	for _, es := range *e {
		pctTotal += es.Pct
	}

	return pctTotal
}

func (e *ErrorSettings) exists(rc int) bool {
	for _, es := range *e {
		if es.RC == rc {
			return true
		}
	}

	return false
}
