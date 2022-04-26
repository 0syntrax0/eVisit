package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	rnd "math/rand"
	"sync"

	"github.com/gin-gonic/gin"
)

func seedProject(g *gin.Context) {
	// create a huge row map
	for i := 0; i < maxRows; i++ {
		// calculate random visits
		visits := make(chan uint8)
		go func() {
			v, err := GenerateRandomNumbersRange(10, 255)
			if err == nil {
				visits <- v
			}
		}()

		// generate a randomized IP addresses list
		ip := randIpAddress()
		count, ok := iPList[ip]
		if ok {
			iPList[ip] = count + 1
		} else {
			iPList[ip] = 1
		}
	}
}

func randIpAddress() string {
	var ip1, ip2, ip3, ip4 uint8
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		var err error
		ip1, err = GenerateRandomNumbersRange(0, 255)
		if err != nil {
			log.Printf("[ip1] failed to generate random IP: %v", err)
		}
	}()
	go func() {
		defer wg.Done()
		var err error
		ip2, err = GenerateRandomNumbersRange(0, 255)
		if err != nil {
			log.Printf("[ip2] failed to generate random IP: %v", err)
		}
	}()
	go func() {
		defer wg.Done()
		var err error
		ip3, err = GenerateRandomNumbersRange(0, 255)
		if err != nil {
			log.Printf("[ip3] failed to generate random IP: %v", err)
		}
	}()
	go func() {
		defer wg.Done()
		var err error
		ip4, err = GenerateRandomNumbersRange(0, 255)
		if err != nil {
			log.Printf("[ip4] failed to generate random IP: %v", err)
		}
	}()

	wg.Wait()
	return fmt.Sprintf("%d.%d.%d.%d", ip1, ip2, ip3, ip4)
}

// GenerateRandomNumbersRange returns random numbers in a given range
// avoiding clock-base seeding as the system's clock is represented in nanoseconds, but the precision isn't nanoseconds
// this is faster and more secure than using `time.Now().UnixNano()`
func GenerateRandomNumbersRange(min, max int) (uint8, error) {
	var b [8]byte
	_, err := rand.Read(b[:])
	if err != nil {
		log.Printf("cannot seed math/rand with cryptographically secure random number generator: min: %d | max: %d", min, max)
		return 0, err
	}
	rnd.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	return uint8(rnd.Intn(max-min+1) + min), nil
}
