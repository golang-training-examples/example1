package server_gin

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Hostname string `json:"hostname"`
}

// ServerGin starts a web Gin (framework) server
func ServerGin(port int) {
	HOSTNAME, _ := os.Hostname()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hostname": HOSTNAME,
		})
	})

	fmt.Printf("Listening on 0.0.0.0:%d, see http://127.0.0.1:%d\n", port, port)
	fmt.Println("Using Gin implementation")
	r.Run(":" + strconv.Itoa(port))
}
