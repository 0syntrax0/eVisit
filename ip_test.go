package main

import (
	"log"
	"testing"
	"time"
)

func TestIpRandom(t *testing.T) {
	start := time.Now()

	// generate 100,000 IPs
	for i := 0; i < 100000; i++ {
		if i%10000 == 0 {
			log.Printf("processed %d", i)
		}

		visits := make(chan uint8)
		go func() {
			v, _ := GenerateRandomNumbersRange(100, 255)
			visits <- v
		}()

		ip := randIpAddress()
		count, ok := iPList[ip]
		if ok {
			iPList[ip] = count + 1
		} else {
			iPList[ip] = int(<-visits)
		}
	}

	log.Printf("--- ended: %v ---", time.Since(start))
}
