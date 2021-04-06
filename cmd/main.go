package main

import (
	"fmt"
	"log"
	"os"
	"parsestats"
)

func main() {
	file, err := os.Open("/proc/stat")
	if err != nil {
		log.Fatalf("We were not able to open '/proc/stats': %v", err)
	}
	defer file.Close()

	stats, err := parsestats.Scan(file)
	if err != nil {
		log.Fatalf("We were not able to scan: %v", err)
	}

	for _, stats := range stats {
		fmt.Printf("%4s: user: %06d, system: %06d, idle: %06d\n",
			stats.ID, stats.User, stats.System, stats.Idle)
	}

}
