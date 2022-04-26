package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// This function accepts a string containing an IP address like “145.87.2.109”.
// This function will be called by the web service every time it handles a request.
// The calling code is outside the scope of this project. Since it is being called very often, this function needs to have a fast runtime.
type RequestHandled struct {
	IP string `json:"ip"`
}

func (rh *RequestHandled) Verify() error {
	if rh.IP == "" {
		return errors.New("missing ip address")
	}

	// check that the given IP is valid
	// by checking formatting and/or pinging it via goRoutine

	return nil
}

func request_handled(c *gin.Context) {
	// accept payload
	var req RequestHandled
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// verify all fields
	if err := req.Verify(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// try to append given IP
	count := iPList[req.IP]
	iPList[req.IP] = count + 1

	// respond with IP accepted
	c.Status(http.StatusAccepted)

	// update top visitors count
	updateTopVisitors(req.IP, count)
}

// This function should return the top 100 IP addresses by request count, with the highest traffic IP address first.
// This function also needs to be fast. Imagine it needs to provide a quick response (< 300ms) to display on a dashboard, even with 20 millions IP addresses.
// This is a very important requirement.
// Don’t forget to satisfy this requirement.
func top(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"list": topIpVisitors})
}

// Called at the start of each day to forget about all IP addresses and tallies.
func clear(c *gin.Context) {
	iPList = make(map[string]int)
	topIpVisitors = [][]string{}
	c.Status(http.StatusResetContent)
}
