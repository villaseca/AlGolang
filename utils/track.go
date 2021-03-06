package utils

import (
	"fmt"
	"time"
)

//Source: https://medium.com/@2xb/execution-time-tracking-in-golang-9379aebfe20e#.js5x4bxf1

//Track measures execution time on current scope.
func Track(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

//Add "defer utils.Track(time.Now(), name)" to start tracking
