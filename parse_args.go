package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type ParseError struct {
	Msg string
}

func (e ParseError) Error() string {
	return fmt.Sprintf("Error: %s", e.Msg)
}


func parseArgs() (string, time.Duration, error) {
	timeout := 60 * 60 * time.Second

	if len(os.Args) < 2 {
		return "", 0, ParseError{"not enough arguments"}
	}

	if len(os.Args) > 3 {
		return "", 0, ParseError{"too many arguments"}
	}

	if len(os.Args) == 3 {
		n, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return "", 0, ParseError{err.Error()}
		}
		timeout = time.Duration(n * 1_000_000_000)
	}

	if timeout < 0 {
		return "", 0, ParseError{"timeout must be positive"}
	}

	return os.Args[1], timeout, nil
}