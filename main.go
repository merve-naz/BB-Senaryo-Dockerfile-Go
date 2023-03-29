package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var count int
var mutex = &sync.Mutex{}

func main() {
	r := gin.Default()
	r.Static("/public", "./public")
	r.GET("/counter", func(c *gin.Context) {

		html := fmt.Sprintf(`<html><body>
			<h1>Counter: %d</h1>
			<form method="post" action="/increment">
				<button type="submit">Increment</button>
			</form>
		</body></html>`, count)

		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, html)
	})

	r.POST("/increment", func(c *gin.Context) {
		mutex.Lock()
		count++
		mutex.Unlock()

		c.Redirect(http.StatusSeeOther, "/counter")
	})

	r.Run(":8080")
}
