package config

import (
	"strconv"
	"time"

	"github.com/alecthomas/units"
)

// Duration is a time.Duration
type Duration time.Duration

// Size is an int64
type Size int64

// UnmarshalTOML parses the duration from the TOML config file
func (d *Duration) UnmarshalText(b []byte) error {
	// convert to string
	durStr := string(b)

	// Value is a TOML number (e.g. 3, 10, 3.5)
	// First try parsing as integer seconds
	sI, err := strconv.ParseInt(durStr, 10, 64)
	if err == nil {
		dur := time.Second * time.Duration(sI)
		*d = Duration(dur)
		return nil
	}
	// Second try parsing as float seconds
	sF, err := strconv.ParseFloat(durStr, 64)
	if err == nil {
		dur := float64(time.Second) * sF
		*d = Duration(dur)
		return nil
	}

	// Finally, try value is a TOML string (e.g. "3s", 3s) or literal (e.g. '3s')
	if durStr == "" {
		durStr = "0s"
	}
	// special case: logging interval had a default of 0d, which silently
	// failed, but in order to prevent issues with default configs that had
	// uncommented the option, change it from zero days to zero hours.
	if durStr == "0d" {
		durStr = "0h"
	}

	dur, err := time.ParseDuration(durStr)
	if err != nil {
		return err
	}

	*d = Duration(dur)
	return nil
}

func (s *Size) UnmarshalText(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	str := string(b)
	val, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		*s = Size(val)
		return nil
	}
	val, err = units.ParseStrictBytes(str)
	if err != nil {
		return err
	}
	*s = Size(val)
	return nil
}
