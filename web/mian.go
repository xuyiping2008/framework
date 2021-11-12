package main

import (
	"net/http"
	"web/ping"
)

func main() {

	r := ping.Default()
	r.GET("/panic", func(c *ping.Context) {
		names := []string{"ping"}
		c.String(http.StatusOK, names[100])
	})
	r.GET("/ok", func(c *ping.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.Run(":9999")
}
