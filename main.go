package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Health struct {
	Name      string        `json:"name"`
	NRP       string        `json:"nrp"`
	Status    bool          `json:"status"`
	Timestamp time.Time     `json:"timestamp"`
	Uptime    time.Duration `json:"uptime"`
}

var startTime time.Time

var healthSeed = Health{
	"Jalu Cahyo Senodiputro",
	"5025241155",
	true,
	time.Now(),
	time.Since(startTime),
}

func formatUptime(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes())
	seconds := int(d.Seconds())

    fmt.Println(hours, minutes, seconds)
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
		"nama":      healthSeed.Name,
		"nrp":       healthSeed.NRP,
		"status":    formatStatus(healthSeed.Status),
		"timestamp": formatTimestamp(healthSeed.Timestamp),
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
