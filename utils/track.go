package utils

import (
	"fmt"
	"time"
)

//Track measures execution time on current scope. Source: https://medium.com/@2xb/execution-time-tracking-in-golang-9379aebfe20e#.js5x4bxf1
func Track(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

//Add "defer Track(time.Now(), name)" to start tracking
