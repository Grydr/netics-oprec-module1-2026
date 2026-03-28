package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime time.Time

func formatUptime(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func formatTimestamp(t time.Time) string {
    return t.Format(time.RFC1123)
}

func formatStatus(status bool) string {
	if status {
		return "UP"
	} else {
		return "DOWN"
	}
}

func getHealth(c *gin.Context) {
	response := gin.H{
		"nama":      "Jalu Cahyo Senodiputro",
		"nrp":       "5025241155",
		"status":    formatStatus(true),
		"timestamp": formatTimestamp(time.Now()),
		"uptime":    formatUptime(time.Since(startTime)),
	}

	c.IndentedJSON(http.StatusOK, response)
}

func init() {
	startTime = time.Now()
}

func main() {
	r := gin.Default()

	r.GET("/health", getHealth)

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
