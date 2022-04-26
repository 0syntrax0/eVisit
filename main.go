package main

import "github.com/gin-gonic/gin"

var (
	iPList = make(map[string]int) // [ip address] visit count
	// topIpVisitorsM = make(map[int]string)              // [visit count] ip address
	topIpVisitors = make([][]string, topVisitorCount) // [key: visit counts] ip address
)

const (
	maxRows         = 20000000
	topVisitorCount = 200 // for performance, this should be a hardcoded max, if a bigger number is needed/requested, then it should be done at the DB or cache layer
)

func main() {
	// start monitoring goRoutine
	go monitoringVisitors()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.PUT("/seed", seedProject)
	router.POST("/registerIP", request_handled)
	router.GET("/top100/", top)
	router.POST("/clear", clear)

	router.Run()
}
